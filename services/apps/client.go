package apps

import "github.com/Azure/go-autorest/autorest"

const (
	// DefaultBaseURI is the default URI used for the service Apps
	DefaultBaseURI = "/cgi-bin/application/appRequest.cgi"
)

// BaseClient is the base client for Apps.
type BaseClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}
