package authapi

import (
	"context"

	"github.com/qnap/core-sdk-for-go/services/auth"
)

// AuthClientAPI contains the set of methods on the VirtualMachinesClient type.
type AuthClientAPI interface {
	Login(ctx context.Context, username string, password string) (auth.LoginResponse, error)
	Logout(ctx context.Context, sid string) error
}

var _ AuthClientAPI = (*auth.Client)(nil)
