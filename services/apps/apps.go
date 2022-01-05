package apps

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
)

type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

func (client Client) List(ctx context.Context) (result ListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apps.Client", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "List", resp, "Failure responding to request")
		return
	}

	return
}

func (client Client) ListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithFormData(url.Values{
			"subfunc": []string{"qpkg"},
			"apply":   []string{"10"},
			"action":  []string{"reload"},
			"lang":    []string{"eng"},
		}))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

func (client Client) ListResponder(resp *http.Response) (result ListResponse, err error) {
	var doc qdocAppList
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingXML(&doc),
		autorest.ByClosing())

	if err != nil {
		return
	}

	if doc.AuthPassed != 1 {
		err = fmt.Errorf("unauthorized")
		return
	}

	result.Apps = make([]ApplicationDetails, len(doc.Func.OwnContent.QItem))
	for i, item := range doc.Func.OwnContent.QItem {
		result.Apps[i] = ApplicationDetails{
			ApplicationState: ApplicationState{
				Application: Application{
					ID:          item.Name,
					DisplayName: item.Attr.DisplayName,
				},
				Date:               item.Attr.Date,
				Version:            item.Attr.Version,
				Build:              item.Attr.Build,
				Status:             item.Attr.Status,
				BootRunStatus:      item.Attr.BootRunStatus,
				ShutdownStopStatus: item.Attr.ShutdownStopStatus,
				Enabled:            item.Attr.Enable == "TRUE",
				Installed:          item.Attr.Installed == "1",
			},
			QPKGFile:          filterNullString(item.Attr.QPKGFile),
			InstallPath:       filterNullString(item.Attr.InstallPath),
			ConfigPath:        filterNullString(item.Attr.ConfigPath),
			ShellPath:         filterNullString(item.Attr.ShellPath),
			Shell:             filterNullString(item.Attr.Shell),
			ServPort:          filterNullString(item.Attr.ServPort),
			Unofficial:        filterNullString(item.Attr.Unofficial),
			IncompleteConf:    filterNullString(item.Attr.IncompleteConf),
			WebPort:           item.Attr.WebPort,
			WebSSLPort:        item.Attr.WebSSLPort,
			WebUI:             filterNullString(item.Attr.WebUI),
			Provider:          filterNullString(item.Attr.Provider),
			Author:            filterNullString(item.Attr.Author),
			Visible:           filterNullString(item.Attr.Visible),
			ForceVisible:      filterNullString(item.Attr.ForceVisible),
			TaskInfo:          filterNullString(item.Attr.TaskInfo),
			SysApp:            item.Attr.SysApp == 1,
			Desktop:           filterNullString(item.Attr.Desktop),
			Class:             filterNullString(item.Attr.Class),
			Store:             filterNullString(item.Attr.Store),
			UserDataPath:      filterNullString(item.Attr.UserDataPath),
			OpenIn:            filterNullString(item.Attr.OpenIn),
			AddOn:             filterNullString(item.Attr.AddOn),
			LoginScreen:       filterNullString(item.Attr.LoginScreen),
			VolumeSelect:      filterNullString(item.Attr.VolumeSelect),
			AppRoute:          filterNullString(item.Attr.AppRoute),
			AppRouteRule:      filterNullString(item.Attr.AppRouteRule),
			FwVerMax:          filterNullString(item.Attr.FwVerMax),
			FwVerMin:          filterNullString(item.Attr.FwVerMin),
			CodeSigningStatus: filterNullString(item.Attr.CodeSigningStatus),
			DepCnt:            filterNullString(item.Attr.DepCnt),
			DepList:           filterNullString(item.Attr.DepList),
		}
	}
	result.Response = autorest.Response{Response: resp}
	return
}

func (client Client) ListStates(ctx context.Context) (result StatesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListStates")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListStatesPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "ListStates", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListStatesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apps.Client", "ListStates", resp, "Failure sending request")
		return
	}

	result, err = client.ListStatesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "ListStates", resp, "Failure responding to request")
		return
	}

	return
}

func (client Client) ListStatesPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithFormData(url.Values{
			"subfunc": []string{"qpkg"},
			"apply":   []string{"10"},
			"action":  []string{"reload"},
			"lang":    []string{"eng"},
		}))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (client Client) ListStatesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

func (client Client) ListStatesResponder(resp *http.Response) (result StatesResponse, err error) {
	var doc qdocAppList
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingXML(&doc),
		autorest.ByClosing())

	if err != nil {
		return
	}

	if doc.AuthPassed != 1 {
		err = fmt.Errorf("unauthorized")
		return
	}

	result.AppStates = make([]ApplicationState, len(doc.Func.OwnContent.QItem))
	for i, item := range doc.Func.OwnContent.QItem {
		result.AppStates[i] = ApplicationState{
			Application: Application{
				ID:          item.Name,
				DisplayName: item.Attr.DisplayName,
			},
			Date:               item.Attr.Date,
			Version:            item.Attr.Version,
			Build:              item.Attr.Build,
			Status:             item.Attr.Status,
			BootRunStatus:      item.Attr.BootRunStatus,
			ShutdownStopStatus: item.Attr.ShutdownStopStatus,
			Enabled:            item.Attr.Enable == "TRUE",
			Installed:          item.Attr.Installed == "1",
		}
	}
	result.Response = autorest.Response{Response: resp}
	return
}

func (client Client) Start(ctx context.Context, qname string, dontWait bool) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Start")
		defer func() {
			sc := -1
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, qname)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "Start", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "Start", resp, "Failure sending request")
		return
	}

	err = client.StartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "Start", resp, "Failure responding to request")
		return
	}

	if !dontWait {
		success := false
		for i := 0; i < 20; i++ {
			stat, ierr := client.getAppTaskStatus(ctx)
			if ierr != nil {
				err = ierr
				return
			}

			if !stat.IsRunning {
				success = true
				break
			}

			time.Sleep(time.Duration(5) * time.Second)
		}

		if !success {
			err = fmt.Errorf("failed to wait for application %s getting started", qname)
			return
		}
	}

	return
}

func (client Client) StartPreparer(ctx context.Context, qname string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithFormData(url.Values{
			"subfunc": []string{"qpkg"},
			"apply":   []string{"3"},
			"block":   []string{"0"},
			"qname":   []string{qname},
		}))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (client Client) StartSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

func (client Client) StartResponder(resp *http.Response) (err error) {
	var doc qdocAppOp
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingXML(&doc),
		autorest.ByClosing())

	if err != nil {
		return
	}

	if doc.AuthPassed != 1 {
		err = fmt.Errorf("unauthorized")
		return
	}

	return
}

func (client Client) Stop(ctx context.Context, qname string, dontWait bool) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Stop")
		defer func() {
			sc := -1
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, qname)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "Stop", resp, "Failure sending request")
		return
	}

	err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "Stop", resp, "Failure responding to request")
		return
	}

	if !dontWait {
		success := false
		for i := 0; i < 20; i++ {
			stat, ierr := client.getAppTaskStatus(ctx)
			if ierr != nil {
				err = ierr
				return
			}

			if !stat.IsRunning {
				success = true
				break
			}

			time.Sleep(time.Duration(5) * time.Second)
		}

		if !success {
			err = fmt.Errorf("failed to wait for application %s getting stopped", qname)
			return
		}
	}

	return
}

func (client Client) StopPreparer(ctx context.Context, qname string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithFormData(url.Values{
			"subfunc": []string{"qpkg"},
			"apply":   []string{"4"},
			"block":   []string{"0"},
			"qname":   []string{qname},
		}))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (client Client) StopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

func (client Client) StopResponder(resp *http.Response) (err error) {
	var doc qdocAppOp
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingXML(&doc),
		autorest.ByClosing())

	if err != nil {
		return
	}

	if doc.AuthPassed != 1 {
		err = fmt.Errorf("unauthorized")
		return
	}

	return
}

func (client Client) getAppTaskStatus(ctx context.Context) (result applicationTaskStatusReponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.getAppTaskStatus")
		defer func() {
			sc := -1
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.getAppTaskStatusPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "getAppTaskStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.getAppTaskStatusSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "getAppTaskStatus", resp, "Failure sending request")
		return
	}

	result, err = client.getAppTaskStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apps.Client", "getAppTaskStatus", resp, "Failure responding to request")
		return
	}

	return
}

func (client Client) getAppTaskStatusPreparer(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"subfunc":            "qpkg",
		"apply":              "1",
		"getstatus":          "2",
		"get_operating_task": "1",
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithQueryParameters(queryParameters))

	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

func (client Client) getAppTaskStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

func (client Client) getAppTaskStatusResponder(resp *http.Response) (result applicationTaskStatusReponse, err error) {
	var doc qdocApplicationStatus
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingXML(&doc),
		autorest.ByClosing())

	if err != nil {
		return
	}

	if doc.AuthPassed != 1 {
		err = fmt.Errorf("unauthorized")
		return
	}

	result.IsRunning = doc.Func.OwnContent.App.Name != ""

	result.Category = doc.Func.OwnContent.App.Category
	result.Class = doc.Func.OwnContent.App.Class
	result.DisplayName = doc.Func.OwnContent.App.DisplayName
	result.DownloadPercent = doc.Func.OwnContent.App.DownloadPercent
	result.Filename = doc.Func.OwnContent.App.Filename
	result.IsUpdate = doc.Func.OwnContent.App.IsUpdate == "1"
	result.Name = doc.Func.OwnContent.App.Name
	result.OpCode = doc.Func.OwnContent.App.OpCode
	result.Operation = doc.Func.OwnContent.App.Operation
	result.StCode = doc.Func.OwnContent.App.StCode
	result.Store = doc.Func.OwnContent.App.Store

	return
}
