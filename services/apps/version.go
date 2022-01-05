package apps

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "QNAP-APPS-SDK-For-Go/v1.0.0-beta services"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return "v1.0.0-beta"
}
