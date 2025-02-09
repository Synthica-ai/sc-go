package repository

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ls := time.Date(2023, 1, 1, 5, 0, 0, 0, time.UTC)
	u, err := MockRepo.CreateUser(uuid.New(), "testcreateuser@stablecog.com", "cus_1234", &ls, nil)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, ls, *u.LastSignInAt)

	// Delete
	MockRepo.DB.User.DeleteOne(u).ExecX(MockRepo.Ctx)
}

func TestSetActiveProductID(t *testing.T) {
	u, err := MockRepo.CreateUser(uuid.New(), "testsetactiveproductid@stablecog.com", "cus_1234", nil, nil)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Nil(t, u.ActiveProductID)

	// Set
	err = MockRepo.SetActiveProductID(u.ID, "prod_1234", nil)
	assert.Nil(t, err)

	// Get user
	u, err = MockRepo.GetUser(u.ID)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "prod_1234", *u.ActiveProductID)

	// Delete
	MockRepo.DB.User.DeleteOne(u).ExecX(MockRepo.Ctx)
}

func TestUnsetActiveProductID(t *testing.T) {
	u, err := MockRepo.CreateUser(uuid.New(), "testunsetactiveproductid@stablecog.com", "cus_1234", nil, nil)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Nil(t, u.ActiveProductID)

	// Set
	err = MockRepo.SetActiveProductID(u.ID, "prod_1234", nil)
	assert.Nil(t, err)

	// Get user
	u, err = MockRepo.GetUser(u.ID)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "prod_1234", *u.ActiveProductID)

	// Unset
	changed, err := MockRepo.UnsetActiveProductID(u.ID, "prod_12345", nil)
	assert.Nil(t, err)
	assert.Equal(t, 0, changed)
	changed, err = MockRepo.UnsetActiveProductID(u.ID, "prod_1234", nil)
	assert.Nil(t, err)
	assert.Equal(t, 1, changed)

	// Get user
	u, err = MockRepo.GetUser(u.ID)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Nil(t, u.ActiveProductID)

	// Delete
	MockRepo.DB.User.DeleteOne(u).ExecX(MockRepo.Ctx)
}

func TestUpdateLastSeenAt(t *testing.T) {
	u, err := MockRepo.CreateUser(uuid.New(), "TestUpdateLastSeenAt@stablecog.com", "cus_1234", nil, nil)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Nil(t, u.ActiveProductID)

	// Set
	err = MockRepo.UpdateLastSeenAt(u.ID)
	assert.Nil(t, err)

	// Get user
	u2, err := MockRepo.GetUser(u.ID)
	assert.Nil(t, err)
	assert.NotNil(t, u2)
	assert.NotEqual(t, u.LastSeenAt, u2.LastSeenAt)

	// Delete
	MockRepo.DB.User.DeleteOne(u).ExecX(MockRepo.Ctx)
}
