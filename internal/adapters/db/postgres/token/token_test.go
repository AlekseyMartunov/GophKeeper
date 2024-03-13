package tokenrepo

import (
	"GophKeeper/internal/adapters/db/postgres/migration"
	"GophKeeper/internal/entity/token"
	"GophKeeper/internal/entity/users"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
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
	repo TokenStorage
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

	query := `TRUNCATE tokens RESTART IDENTITY CASCADE;`
	_, err = pool.Exec(context.Background(), query)
	assert.NoError(suite.T(), err, "truncate db error")

	suite.repo = *NewTokenStorage(pool)
	suite.pool = pool
}

// TearDownTest this function is called every time the test completes
func (suite *dataBaseSuite) TearDownTest() {
	query := `TRUNCATE tokens RESTART IDENTITY CASCADE;`
	_, err := suite.pool.Exec(context.Background(), query)
	assert.NoError(suite.T(), err, "truncate db after tests error")

	err = migration.MigrationsDown(dns)
	assert.NoError(suite.T(), err, "remove tables error")
}

func (suite *dataBaseSuite) TestSaveToken() {
	err := suite.supportCreateUsers()
	assert.NoError(suite.T(), err)

	t1 := token.Token{
		Name:        "for my hom pc 1",
		Token:       "1234567890",
		CreatedTime: time.Now(),
		UserID:      1,
	}

	t2 := token.Token{
		Name:        "for my hom pc 2",
		Token:       "1234567890",
		CreatedTime: time.Now(),
		UserID:      1,
	}

	t3 := token.Token{
		Name:        "for my hom pc 1",
		Token:       "1234567890",
		CreatedTime: time.Now(),
		UserID:      1,
	}

	err = suite.repo.Save(context.Background(), t1)
	assert.NoError(suite.T(), err)

	err = suite.repo.Save(context.Background(), t2)
	assert.NoError(suite.T(), err)

	err = suite.repo.Save(context.Background(), t3)
	assert.ErrorIs(suite.T(), err, token.ErrTokenAlreadyExists)

	tokens := suite.supportGetTokens()
	expectedTokens := []token.Token{t1, t2}
	for i, t := range tokens {
		assert.Equal(suite.T(), t.Token, expectedTokens[i].Token)
		assert.Equal(suite.T(), t.Name, expectedTokens[i].Name)
		assert.Equal(suite.T(), t.Token, expectedTokens[i].Token)
		assert.Equal(suite.T(), t.UserID, expectedTokens[i].UserID)
		assert.Equal(suite.T(), t.IsBlocked, false)
	}
}

func (suite *dataBaseSuite) TestBlockToken() {
	err := suite.supportCreateTokens()
	assert.NoError(suite.T(), err)

	suite.repo.BlockToken(context.Background(), "1234567890")

	tokens := suite.supportGetTokens()

	assert.True(suite.T(), tokens[1].IsBlocked)
	assert.False(suite.T(), tokens[0].IsBlocked)
}

func (suite *dataBaseSuite) TestGetUserID() {
	err := suite.supportCreateTokens()
	assert.NoError(suite.T(), err)

	id, err := suite.repo.GetUserIDByExternalID(context.Background(), "1234567890")
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), id, 1)

}

func (suite *dataBaseSuite) supportGetTokens() []token.Token {
	query := `SELECT * FROM tokens`

	rows, err := suite.pool.Query(context.Background(), query)
	if err != nil {
		return nil
	}

	tokens := make([]token.Token, 0, 3)

	for rows.Next() {
		t := token.Token{}
		err = rows.Scan(&t.TokenID, &t.Name, &t.Token, &t.CreatedTime, &t.IsBlocked, &t.UserID)
		if err != nil {
			return nil
		}

		tokens = append(tokens, t)
	}

	if rows.Err() != nil {
		return nil
	}

	return tokens
}

func (suite *dataBaseSuite) supportCreateTokens() error {
	suite.supportCreateUsers()
	t1 := token.Token{
		Name:        "for my hom pc 1",
		Token:       "1234567890",
		CreatedTime: time.Now(),
		UserID:      1,
	}

	t2 := token.Token{
		Name:        "for my hom pc 2",
		Token:       "aaaaaaaaaa",
		CreatedTime: time.Now(),
		UserID:      1,
	}

	query := `INSERT INTO tokens (token_name, token, created_time, fk_user_id)
				VALUES ($1, $2, $3, $4)`

	_, err := suite.pool.Exec(context.Background(), query, t1.Name, t1.Token, t1.CreatedTime, t1.UserID)
	if err != nil {
		return err
	}

	_, err = suite.pool.Exec(context.Background(), query, t2.Name, t2.Token, t2.CreatedTime, t2.UserID)
	if err != nil {
		return err
	}
	return nil
}

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
