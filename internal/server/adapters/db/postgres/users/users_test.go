package usersrepo

import (
	"context"
	"testing"

	"GophKeeper/internal/server/adapters/db/postgres/migration"
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
	repo UserStorage
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
	assert.NoError(suite.T(), err, "truncate db error")

	suite.repo = *NewUserStorage(pool)
	suite.pool = pool
}

// TearDownTest this function is called every time the test completes
func (suite *dataBaseSuite) TearDownTest() {
	query := `TRUNCATE users RESTART IDENTITY CASCADE;`
	_, err := suite.pool.Exec(context.Background(), query)
	assert.NoError(suite.T(), err, "truncate db after tests error")

	err = migration.MigrationsDown(dns)
	assert.NoError(suite.T(), err, "remove tables error")
}

func (suite *dataBaseSuite) TestSaveUser() {
	testCase := []struct {
		user           users.User
		expectedResult users.User
	}{
		{
			user: users.User{
				Login:    "AAA",
				Password: "123",
			},
			expectedResult: users.User{
				Login:    "AAA",
				Password: "123",
				ID:       1,
			},
		},
		{
			user: users.User{
				Login:    "BBB",
				Password: "123",
			},
			expectedResult: users.User{
				Login:    "BBB",
				Password: "123",
				ID:       2,
			},
		},
		{
			user: users.User{
				Login:    "CCC",
				Password: "123",
			},
			expectedResult: users.User{
				Login:    "CCC",
				Password: "123",
				ID:       3,
			},
		},
	}

	for _, tc := range testCase {
		suite.repo.Save(context.Background(), tc.user)
	}

	usersArr, err := suite.supportGetUsersForTest()
	assert.NoError(suite.T(), err, "getting users error")

	for ind, tc := range testCase {
		assert.Equal(suite.T(), tc.expectedResult.ID, usersArr[ind].ID)
		assert.Equal(suite.T(), tc.expectedResult.Login, usersArr[ind].Login)
		assert.Equal(suite.T(), tc.expectedResult.Password, usersArr[ind].Password)

		assert.NotEmpty(suite.T(), usersArr[ind].ExternalID)
	}
}

func (suite *dataBaseSuite) TestUserAlreadyExists() {
	u := users.User{
		Login:    "AAA",
		Password: "BBB",
	}

	err := suite.repo.Save(context.Background(), u)
	assert.NoError(suite.T(), err, "save user error")

	err = suite.repo.Save(context.Background(), u)
	assert.ErrorIs(suite.T(), err, users.ErrUserAlreadyExists,
		"repo dose not return ErrUserAlreadyExists error")
}

func (suite *dataBaseSuite) TestUserDoseNotExist() {
	u := users.User{
		Login:    "AAA",
		Password: "BBB",
	}

	err := suite.repo.Save(context.Background(), u)
	assert.NoError(suite.T(), err, "save user error")

	u2 := users.User{
		Login:    "BBB",
		Password: "qwerty",
	}

	id, err := suite.repo.GetExternalID(context.Background(), u2)

	assert.Empty(suite.T(), id, "id is not empty")
	assert.ErrorIs(suite.T(), err, users.ErrUserDoseNotExist,
		"repo dose not return ErrUserDoseNotExist error")
}

func (suite *dataBaseSuite) TestGetExternalID() {
	u := users.User{
		Login:    "AAA",
		Password: "qwerty",
	}

	err := suite.repo.Save(context.Background(), u)
	assert.NoError(suite.T(), err, "save user error")

	result, err := suite.repo.GetExternalID(context.Background(), u)
	assert.NoError(suite.T(), err, "get id error")

	assert.NotEmpty(suite.T(), result, "external Id is empty")
}

// supportGetUsersForTest support function for get users from db
func (suite *dataBaseSuite) supportGetUsersForTest() ([]users.User, error) {
	query := `SELECT user_id, login, password, external_id FROM users`

	row, err := suite.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	usersArr := make([]users.User, 0, 1)

	for row.Next() {
		u := users.User{}

		err := row.Scan(&u.ID, &u.Login, &u.Password, &u.ExternalID)
		if err != nil {
			return nil, err
		}
		usersArr = append(usersArr, u)
	}

	err = row.Err()
	if err != nil {
		return nil, err
	}

	return usersArr, nil
}
