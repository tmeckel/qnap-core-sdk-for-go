package appsapi

import (
	"context"

	"github.com/qnap/core-sdk-for-go/services/apps"
)

// AppsClientAPI contains the set of methods on the VirtualMachinesClient type.
type AppsClientAPI interface {
	List(ctx context.Context) (apps.ListResponse, error)
	ListStates(ctx context.Context) (apps.StatesResponse, error)
	Start(ctx context.Context, qname string, dontWait bool) error
	Stop(ctx context.Context, qname string, dontWait bool) error
}

var _ AppsClientAPI = (*apps.Client)(nil)
