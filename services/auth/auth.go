package auth

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
)

type Client struct {
	BaseClient
}

// NewClientClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

func (client Client) Login(ctx context.Context, username string, password string) (result LoginResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Login")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}

	req, err := client.LoginPreparer(ctx, username, password)
	if err != nil {
		err = autorest.NewErrorWithError(err, "core.Client", "Login", nil, "Failure preparing request")
		return
	}

	resp, err := client.LoginSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "core.Client", "Login", resp, "Failure sending request")
		return
	}

	result, err = client.LoginResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "core.Client", "Login", resp, "Failure responding to request")
		return
	}

	return
}

func (client Client) LoginPreparer(ctx context.Context, username string, password string) (*http.Request, error) {
	secret := b64.StdEncoding.EncodeToString([]byte(password))

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI+"/authLogin.cgi"),
		autorest.WithFormData(url.Values{
			"user": []string{username},
			"pwd":  []string{secret},
		}))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (client Client) LoginSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

func (client Client) LoginResponder(resp *http.Response) (result LoginResponse, err error) {
	var doc qdocRoot

	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingXML(&doc),
		autorest.ByClosing())

	if err != nil {
		return
	}

	if doc.AuthPassed != 1 {
		return result, fmt.Errorf("login failed")
	}

	result.Sid = doc.AuthSid
	result.IsAdmin = doc.IsAdmin == "1"
	result.Username = doc.Username
	result.Groupname = doc.Groupname

	return
}

func (client Client) Logout(ctx context.Context, sid string) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Login")
		defer func() {
			sc := -1
			tracing.EndSpan(ctx, sc, err)
		}()
	}

	req, err := client.LogoutPreparer(ctx, sid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "core.Client", "Logout", nil, "Failure preparing request")
		return
	}

	resp, err := client.LogoutSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "core.Client", "Logout", resp, "Failure sending request")
		return
	}

	err = client.LogoutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "core.Client", "Logout", resp, "Failure responding to request")
		return
	}

	return
}

func (client Client) LogoutPreparer(ctx context.Context, sid string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"sid": autorest.Encode("query", sid),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI+"/authLogout.cgi"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithFormData(url.Values{
			"logout": []string{"1"},
		}))

	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (client Client) LogoutSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

func (client Client) LogoutResponder(resp *http.Response) (err error) {
	var doc qdocRoot
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingXML(&doc),
		autorest.ByClosing())

	if err != nil {
		return
	}

	if doc.AuthPassed != 1 {
		return fmt.Errorf("logout failed")
	}

	return
}
