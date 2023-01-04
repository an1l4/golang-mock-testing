package user_test

import (
	"testing"

	"github.com/an1l4/mocktest/mocks"
	"github.com/an1l4/mocktest/user"
	"github.com/golang/mock/gomock"
)

func TestUse(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	mockUser := mocks.NewMockIUserInterface(mockctrl)
	testUser := &user.User{IUser: mockUser}

	mockUser.EXPECT().AddUser(1, "sample test").Return(nil).Times(1)
	testUser.Use()

}
