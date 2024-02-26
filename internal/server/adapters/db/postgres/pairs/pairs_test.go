package pairsrepo

import (
	"context"
	"testing"
	"time"

	"GophKeeper/internal/server/adapters/db/postgres/migration"
	"GophKeeper/internal/server/entity/pairs"
	"GophKeeper/internal/server/entity/users"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ==============================================================================
// ==================== BEFORE TESTING RUN docker compose up ====================
// ==============================================================================
//
// there is a docker-compose.yml file in the build folder, in main project folder

const (
	dns = "postgres://admin:1234@localhost:5432/test?sslmode=disable"
)

type dataBaseSuite struct {
	suite.Suite
	repo PairStorage
	pool *pgxpool.Pool
}

func TestBaseSuite(t *testing.T) {
	suite.Run(t, new(dataBaseSuite))
}

// SetupTest this function is called every time the test starts
func (suite *dataBaseSuite) SetupTest() {
	err := migration.MigrationsUp(dns)
	assert.NoError(suite.T(), err, "create migrations error")

	pool, err := pgxpool.New(context.Background(), dns)
	assert.NoError(suite.T(), err, "open pool error")

	query := `TRUNCATE users RESTART IDENTITY CASCADE;`
	_, err = pool.Exec(context.Background(), query)
	assert.NoError(suite.T(), err, "truncate users table error")

	query = `TRUNCATE pairs RESTART IDENTITY CASCADE;`
	_, err = pool.Exec(context.Background(), query)
	assert.NoError(suite.T(), err, "truncate pairs table error")

	suite.repo = *NewPairsStorage(pool)
	suite.pool = pool
}

// TearDownTest this function is called every time the test completes
func (suite *dataBaseSuite) TearDownTest() {
	query := `TRUNCATE users RESTART IDENTITY CASCADE;`
	_, err := suite.pool.Exec(context.Background(), query)
	assert.NoError(suite.T(), err, "truncate user tables after tests error")

	query = `TRUNCATE pairs RESTART IDENTITY CASCADE;`
	_, err = suite.pool.Exec(context.Background(), query)
	assert.NoError(suite.T(), err, "truncate pairs tables after tests error")

	err = migration.MigrationsDown(dns)
	assert.NoError(suite.T(), err, "remove tables error")

}

func (suite *dataBaseSuite) TestSavePais() {
	err := suite.supportCreateUsers()
	assert.NoError(suite.T(), err)
	ctx := context.Background()

	p := pairs.Pair{
		Name:        "login and password for example.com",
		Login:       "some login",
		Password:    "qwerty",
		CreatedTime: time.Now(),
	}

	// user with id = 1 save pair
	err = suite.repo.Save(ctx, p, users.User{ID: 1})
	assert.NoError(suite.T(), err)

	// user with id = 2 save pair
	err = suite.repo.Save(ctx, p, users.User{ID: 2})
	assert.NoError(suite.T(), err)

	// user with id = 2 save pair
	// error because user with id = 2 already has pair with same name
	err = suite.repo.Save(ctx, p, users.User{ID: 2})
	assert.ErrorIs(suite.T(), err, pairs.ErrPairAlreadyExists)
}

func (suite *dataBaseSuite) TestGetPair() {
	err := suite.supportCreatePairs()
	assert.NoError(suite.T(), err)
	ctx := context.Background()

	// get pair with name pair1 for user with id 1
	result, err := suite.repo.Get(ctx, users.User{ID: 1}, "pair1")
	assert.NoError(suite.T(), err)

	wanted := pairs.Pair{
		Name:     "pair1",
		Login:    "aaa",
		Password: "AAA",
		User:     users.User{ID: 1},
	}

	assert.Equal(suite.T(), wanted.Name, result.Name)
	assert.Equal(suite.T(), wanted.Login, result.Login)
	assert.Equal(suite.T(), wanted.Password, result.Password)
	assert.Equal(suite.T(), wanted.User, result.User)

	// should be ErrPairDoseNotExist
	result, err = suite.repo.Get(ctx, users.User{ID: 20}, "pair1")
	assert.ErrorIs(suite.T(), err, pairs.ErrPairDoseNotExist)
}

func (suite *dataBaseSuite) TestGetAllPairs() {
	err := suite.supportCreatePairs()
	assert.NoError(suite.T(), err)
	ctx := context.Background()

	wanted := []pairs.Pair{
		{
			Name:     "pair2",
			Login:    "bbb",
			Password: "BBB",
			User:     users.User{ID: 2},
		},
		{
			Name:     "pair3",
			Login:    "ccc",
			Password: "CCC",
			User:     users.User{ID: 2},
		},
	}

	// get all pairs for user with id = 2
	result, err := suite.repo.GetAll(ctx, users.User{ID: 2})
	assert.NoError(suite.T(), err)

	for i, w := range wanted {
		assert.Equal(suite.T(), w.Name, result[i].Name)
		assert.Equal(suite.T(), w.Login, result[i].Login)
		assert.Equal(suite.T(), w.Password, result[i].Password)
		assert.Equal(suite.T(), w.User, result[i].User)
	}

	result, err = suite.repo.GetAll(ctx, users.User{ID: 20})
	assert.ErrorIs(suite.T(), err, pairs.ErrPairDoseNotExist)

}

func (suite *dataBaseSuite) TestDelete() {
	err := suite.supportCreatePairs()
	assert.NoError(suite.T(), err)
	ctx := context.Background()

	err = suite.repo.Delete(ctx, users.User{ID: 2}, "pair3")
	assert.NoError(suite.T(), err)

	wanted := []pairs.Pair{
		{
			Name:     "pair1",
			Login:    "aaa",
			Password: "AAA",
			User:     users.User{ID: 1},
		},
		{
			Name:     "pair2",
			Login:    "bbb",
			Password: "BBB",
			User:     users.User{ID: 2},
		},
	}

	result, err := suite.supportGetAllPairsFromDB()

	for i, w := range wanted {
		assert.Equal(suite.T(), w.Name, result[i].Name)
		assert.Equal(suite.T(), w.Login, result[i].Login)
		assert.Equal(suite.T(), w.Password, result[i].Password)
		assert.Equal(suite.T(), w.User, result[i].User)
	}

	// nothing to delete
	err = suite.repo.Delete(ctx, users.User{ID: 20}, "pair3")
	assert.ErrorIs(suite.T(), err, pairs.ErrPairNothingToDelete)
}

// supportCreateUsers support function for TestSavePais
// this function save users in users table
func (suite *dataBaseSuite) supportCreateUsers() error {
	// user db after running supportCreateUsers
	//
	//  user_id | login | password |             external_id
	// ---------+-------+----------+--------------------------------------
	//        1 | a     | a        | 6e47ae9b-1137-45ac-90ac-1b985356ac3e
	//        2 | b     | b        | a329374c-9edd-412b-9e33-99da940e1d89
	//        3 | c     | c        | beb85e92-f1c6-40ad-a181-780ceadb5dd3

	ctx := context.Background()
	u1 := users.User{
		Login:    "a",
		Password: "a",
	}
	u2 := users.User{
		Login:    "b",
		Password: "b",
	}
	u3 := users.User{
		Login:    "c",
		Password: "c",
	}

	query := `INSERT INTO users (login, password) VALUES ($1, $2)`

	_, err := suite.pool.Exec(ctx, query, u1.Login, u1.Password)
	if err != nil {
		return err
	}

	_, err = suite.pool.Exec(ctx, query, u2.Login, u2.Password)
	if err != nil {
		return err
	}

	_, err = suite.pool.Exec(ctx, query, u3.Login, u3.Password)
	if err != nil {
		return err
	}

	return nil
}

// supportCreatePairs support function
// this function save pairs in DB
func (suite *dataBaseSuite) supportCreatePairs() error {
	// supportCreatePairs should be run after supportCreateUsers
	suite.supportCreateUsers()

	// pair db after run supportCreatePairs
	//
	//  pair_id | pair_name | password | login |        created_time        | fk_user_id
	// ---------+-----------+----------+-------+----------------------------+------------
	//        1 | pair1     | AAA      | aaa   | 2024-02-26 16:03:04.940471 |          1
	//        2 | pair2     | BBB      | bbb   | 2024-02-26 16:03:04.940471 |          2
	//        3 | pair3     | CCC      | ccc   | 2024-02-26 16:03:04.940471 |          2

	ctx := context.Background()
	p1 := pairs.Pair{
		Name:        "pair1",
		Login:       "aaa",
		Password:    "AAA",
		CreatedTime: time.Now(),
	}

	p2 := pairs.Pair{
		Name:        "pair2",
		Login:       "bbb",
		Password:    "BBB",
		CreatedTime: time.Now(),
	}

	p3 := pairs.Pair{
		Name:        "pair3",
		Login:       "ccc",
		Password:    "CCC",
		CreatedTime: time.Now(),
	}

	query := `INSERT INTO pairs (pair_name, password, login, created_time, fk_user_id) 
				values ($1, $2, $3, $4, $5);`

	_, err := suite.pool.Exec(ctx, query, p1.Name, p1.Password, p1.Login, p1.CreatedTime, 1)
	if err != nil {
		return err
	}
	_, err = suite.pool.Exec(ctx, query, p2.Name, p2.Password, p2.Login, p2.CreatedTime, 2)
	if err != nil {
		return err
	}
	_, err = suite.pool.Exec(ctx, query, p3.Name, p3.Password, p3.Login, p3.CreatedTime, 2)
	if err != nil {
		return err
	}
	return nil
}

// supportGetAllPairsFromDB returns all objects that remain in the database
func (suite *dataBaseSuite) supportGetAllPairsFromDB() ([]pairs.Pair, error) {
	query := `SELECT pair_id, pair_name, password, login, created_time, fk_user_id FROM pairs`

	rows, err := suite.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	pairsArr := make([]pairs.Pair, 0, 10)
	for rows.Next() {
		p := pairs.Pair{}
		u := users.User{}

		err := rows.Scan(&p.ID, &p.Name, &p.Password, &p.Login, &p.CreatedTime, &u.ID)
		if err != nil {
			return nil, err
		}
		p.User = u
		pairsArr = append(pairsArr, p)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return pairsArr, nil
}
