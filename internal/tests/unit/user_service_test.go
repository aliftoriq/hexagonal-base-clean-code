package unit

import (
	"github.com/hibbannn/hexagonal-boilerplate/internal/adapters/cache"
	"github.com/hibbannn/hexagonal-boilerplate/internal/adapters/repository"
	"github.com/hibbannn/hexagonal-boilerplate/internal/core/domain"
	"github.com/hibbannn/hexagonal-boilerplate/internal/core/usecase"
	"github.com/hibbannn/hexagonal-boilerplate/internal/tests/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setUpDB() *repository.DB {
	db, _ := gorm.Open("postgres", "postgres://hibban:postgres@localhost:5432/wdp_rnd?sslmode=disable")
	db.AutoMigrate(&domain.Message{}, &domain.User{}, &domain.Payment{})
	// defer db.Close()

	redisCache, err := cache.NewRedisCache("localhost:6379", "")
	if err != nil {
		panic(err)
	}

	store := repository.NewDB(db, redisCache)

	return store
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockUser := &domain.User{ID: "someID", Email: "test@example.com", Password: "hashedPassword"}
	mockRepo.On("CreateUser", mock.Anything, mock.Anything).Return(mockUser, nil)

	userService := usecase.NewUserService(mockRepo)
	user, err := userService.CreateUser("test@example.com", "password")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, mockUser, user)
	mockRepo.AssertExpectations(t)
}

func TestReadUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockUser := &domain.User{ID: "someID", Email: "test@example.com", Password: "hashedPassword"}
	mockRepo.On("ReadUser", mock.Anything).Return(mockUser, nil)

	userService := usecase.NewUserService(mockRepo)
	user, err := userService.ReadUser("someID")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, mockUser, user)
	mockRepo.AssertExpectations(t)

}

//func TestReadUser(t *testing.T) {
//	db := setUpDB()
//
//	email := "test@example.com"
//	password := "password"
//
//	user, err := db.CreateUser(email, password)
//	assert.NoError(t, err)
//	assert.NotNil(t, user)
//
//	cachedUser, err := db.ReadUser(user.ID)
//	assert.NoError(t, err)
//	assert.NotNil(t, cachedUser)
//	assert.Equal(t, user.ID, cachedUser.ID)
//	assert.Equal(t, user.Email, cachedUser.Email)
//	assert.Equal(t, user.Password, cachedUser.Password)
//
//	time.Sleep(time.Second * 3)
//
//	cachedUser, err = db.ReadUser(user.ID)
//	assert.Error(t, err)
//	assert.Nil(t, cachedUser)
//}

//func TestReadUsers(t *testing.T) {
//	db := setUpDB()
//
//	email := "test@example.com"
//	password := "hashedPassword"
//
//	user, err := db.CreateUser(email, password)
//	assert.NoError(t, err)
//	assert.NotNil(t, user)
//
//	users, err := db.ReadUsers()
//	assert.NoError(t, err)
//	assert.NotNil(t, users)
//	assert.NotEmpty(t, users)
//}
//
//func TestUpdateUser(t *testing.T) {
//	db := setUpDB()
//
//	email := "test@example.com"
//	password := "password"
//
//	user, err := db.CreateUser(email, password)
//	assert.NoError(t, err)
//	assert.NotNil(t, user)
//
//	newEmail := "new@example.com"
//	newPassword := "newpassword"
//
//	err = db.UpdateUser(user.ID, newEmail, newPassword)
//	assert.NoError(t, err)
//
//	cachedUser, err := db.ReadUser(user.ID)
//	assert.NoError(t, err)
//	assert.NotNil(t, cachedUser)
//	assert.Equal(t, newEmail, cachedUser.Email)
//	assert.NotEqual(t, password, cachedUser.Password)
//}
//
//func TestDeleteUser(t *testing.T) {
//	db := setUpDB()
//
//	email := "test@example.com"
//	password := "password"
//
//	user, err := db.CreateUser(email, password)
//	assert.NoError(t, err)
//	assert.NotNil(t, user)
//
//	err = db.DeleteUser(user.ID)
//	assert.NoError(t, err)
//
//	cachedUser, err := db.ReadUser(user.ID)
//	assert.Error(t, err)
//	assert.Nil(t, cachedUser)
//
//	users, err := db.ReadUsers()
//	assert.NoError(t, err)
//	assert.NotNil(t, users)
//	assert.Empty(t, users)
//}
//
//func TestCreateUserAlreadyExists(t *testing.T) {
//	db := setUpDB()
//
//	email := "test@example.com"
//	password := "password"
//
//	user, err := db.CreateUser(email, password)
//	assert.NoError(t, err)
//	assert.NotNil(t, user)
//
//	user, err = db.CreateUser(email, password)
//	assert.Error(t, err)
//	assert.Nil(t, user)
//	// assert.True(t, errors.Is(err, repository.ErrUserAlreadyExists))
//}
//
//func TestReadUserNotFound(t *testing.T) {
//	db := setUpDB()
//
//	user, err := db.ReadUser("nonexistent")
//	assert.Error(t, err)
//	assert.Nil(t, user)
//	// assert.True(t, errors.Is(err, repository.ErrUserNotFound))
//}
//
//func TestUpdateUserNotFound(t *testing.T) {
//	db := setUpDB()
//
//	err := db.UpdateUser("nonexistent", "new@example.com", "newpassword")
//	assert.Error(t, err)
//	// assert.True(t, errors.Is(err, repository.ErrUserNotFound))
//}
//
//func TestDeleteUserNotFound(t *testing.T) {
//	db := setUpDB()
//
//	err := db.DeleteUser("nonexistent")
//	assert.Error(t, err)
//	// assert.True(t, errors.Is(err, repository.ErrUserNotFound))
//}
