package apps

import (
	"encoding/xml"
	"strings"

	"github.com/Azure/go-autorest/autorest"
)

const fqdn = "github.com/qnap/apps-sdk-for-go/services/apps"

func filterNullString(s string) string {
	if strings.EqualFold(s, "null") {
		return ""
	}
	return s
}

type qdocAppList struct {
	autorest.Response `xml:"-"`
	XMLName           xml.Name `xml:"QDocRoot"`
	Version           string   `xml:"version,attr"`
	AuthPassed        int      `xml:"authPassed"`
	Func              struct {
		Name       string `xml:"name"`
		OwnContent struct {
			QItem []struct {
				Name string `xml:"name"`
				Attr struct {
					DisplayName        string `xml:"displayName"`
					QPKGFile           string `xml:"QPKGFile"`
					Date               string `xml:"date"`
					Version            string `xml:"version"`
					Build              string `xml:"build"`
					InstallPath        string `xml:"installPath"`
					ConfigPath         string `xml:"configPath"`
					ShellPath          string `xml:"shellPath"`
					Shell              string `xml:"shell"`
					Enable             string `xml:"enable"`
					Installed          string `xml:"installed"`
					ServPort           string `xml:"servPort"`
					Unofficial         string `xml:"unofficial"`
					IncompleteConf     string `xml:"incomplete_conf"`
					WebPort            int    `xml:"webPort"`
					WebSSLPort         int    `xml:"webSSLPort"`
					WebUI              string `xml:"webUI"`
					Provider           string `xml:"provider"`
					Author             string `xml:"author"`
					Visible            string `xml:"visible"`
					ForceVisible       string `xml:"force_visible"`
					TaskInfo           string `xml:"task_info"`
					SysApp             int    `xml:"sysApp"`
					Desktop            string `xml:"desktop"`
					Class              string `xml:"class"`
					Store              string `xml:"store"`
					Status             string `xml:"status"`
					UserDataPath       string `xml:"userDataPath"`
					OpenIn             string `xml:"openIn"`
					AddOn              string `xml:"addOn"`
					LoginScreen        string `xml:"loginScreen"`
					VolumeSelect       string `xml:"volume_select"`
					AppRoute           string `xml:"app_route"`
					AppRouteRule       string `xml:"app_route_rule"`
					BootRunStatus      string `xml:"boot_run_status"`
					ShutdownStopStatus string `xml:"shutdown_stop_status"`
					FwVerMax           string `xml:"fw_ver_max"`
					FwVerMin           string `xml:"fw_ver_min"`
					HaveToUpdate       string `xml:"have_to_update"`
					CodeSigningStatus  string `xml:"code_signing_status"`
					DepCnt             string `xml:"dep_cnt"`
					DepList            string `xml:"dep_list"`
				} `xml:"attr"`
			} `xml:"qItem"`
		} `xml:"ownContent"`
	} `xml:"func"`
	Result string `xml:"result"`
}

type Application struct {
	ID          string `xml:"id" json:"id" yaml:"id"`
	DisplayName string `xml:"displayName" json:"displayName" yaml:"displayName"`
}

type ApplicationState struct {
	Application        `yaml:",inline"`
	Date               string `xml:"date" json:"date" yaml:"date"`
	Version            string `xml:"version" json:"version" yaml:"verion"`
	Build              string `xml:"build" json:"build" yaml:"build"`
	Status             string `xml:"status" json:"status" yaml:"status"`
	BootRunStatus      string `xml:"bootRunStatus" json:"bootRunStatus" yaml:"bootRunStatus"`
	ShutdownStopStatus string `xml:"shutdownStopStatus" json:"shutdownStopStatus" yaml:"shutdownStopStatus"`
	Enabled            bool   `xml:"enabled" json:"enabled" yaml:"enabled"`
	Installed          bool   `xml:"installed" json:"installed" yaml:"installed"`
}

type ApplicationDetails struct {
	ApplicationState  `yaml:",inline"`
	QPKGFile          string `xml:"qpkgFile" json:"qpkgFile" yaml:"qpkgFile"`
	InstallPath       string `xml:"installPath" json:"installPath" yaml:"installPath"`
	ConfigPath        string `xml:"configPath,omitempty" json:"configPath,omitempty" yaml:"configPath,omitempty"`
	ShellPath         string `xml:"shellPath,omitempty" json:"shellPath,omitempty" yaml:"shellPath,omitempty"`
	Shell             string `xml:"shell,omitempty" json:"shell,omitempty" yaml:"shell,omitempty"`
	ServPort          string `xml:"-" json:"-" yaml:"-"`
	Unofficial        string `xml:"-" json:"-" yaml:"-"`
	IncompleteConf    string `xml:"-" json:"-" yaml:"-"`
	WebPort           int    `xml:"webPort" json:"webPort" yaml:"webPort"`
	WebSSLPort        int    `xml:"webSSLPort" json:"webSSLPort" yaml:"webSSLPort"`
	WebUI             string `xml:"-" json:"-" yaml:"-"`
	Provider          string `xml:"provider,omitempty" json:"provider,omitempty" yaml:"provider,omitempty"`
	Author            string `xml:"author" json:"author" yaml:"author"`
	Visible           string `xml:"-" json:"-" yaml:"-"`
	ForceVisible      string `xml:"-" json:"-" yaml:"-"`
	TaskInfo          string `xml:"-" json:"-" yaml:"-"`
	SysApp            bool   `xml:"sysApp" json:"sysApp" yaml:"sysApp"`
	Desktop           string `xml:"-" json:"-" yaml:"-"`
	Class             string `xml:"-" json:"-" yaml:"-"`
	Store             string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
	UserDataPath      string `xml:"userDataPath,omitempty" json:"userDataPath,omitempty" yaml:"userDataPath,omitempty"`
	OpenIn            string `xml:"-" json:"-" yaml:"-"`
	AddOn             string `xml:"-" json:"-" yaml:"-"`
	LoginScreen       string `xml:"-" json:"-" yaml:"-"`
	VolumeSelect      string `xml:"-" json:"-" yaml:"-"`
	AppRoute          string `xml:"-" json:"-" yaml:"-"`
	AppRouteRule      string `xml:"-" json:"-" yaml:"-"`
	FwVerMax          string `xml:"fwVerMax,omitempty" json:"fwVerMax,omitempty" yaml:"fwVerMax,omitempty"`
	FwVerMin          string `xml:"fwVerMin,omitempty" json:"fwVerMin,omitempty" yaml:"fwVerMin,omitempty"`
	CodeSigningStatus string `xml:"-" json:"-" yaml:"-"`
	DepCnt            string `xml:"-" json:"-" yaml:"-"`
	DepList           string `xml:"-" json:"-" yaml:"-"`
}

type ApplicationUpdateInfo struct {
	Application
	AvailableVersion string
	InstalledVersion string
}

type ListResponse struct {
	autorest.Response `xml:"-" json:"-" yaml:"-"`
	Apps              []ApplicationDetails
}

type StatesResponse struct {
	autorest.Response `xml:"-" json:"-" yaml:"-"`
	AppStates         []ApplicationState
}

type UpdateListResponse struct {
	autorest.Response `xml:"-" json:"-" yaml:"-"`
	Apps              []ApplicationUpdateInfo
}

type qdocAppOp struct {
	XMLName    xml.Name `xml:"QDocRoot" json:"qdocroot,omitempty"`
	AuthPassed int      `xml:"authPassed"`
}

type qdocApplicationStatus struct {
	XMLName    xml.Name `xml:"QDocRoot"`
	AuthPassed int      `xml:"authPassed"`
	Func       struct {
		Name       string `xml:"name"`
		OwnContent struct {
			App struct {
				Name            string `xml:"name"`
				DisplayName     string `xml:"display_name"`
				Filename        string `xml:"filename"`
				Store           string `xml:"store"`
				OpCode          string `xml:"op_code"`
				StCode          string `xml:"st_code"`
				Class           string `xml:"class"`
				Category        string `xml:"category"`
				Version         string `xml:"version"`
				DownloadPercent string `xml:"downloadPercent"`
				Operation       string `xml:"operation"`
				IsUpdate        string `xml:"isUpdate"`
			} `xml:"app"`
		} `xml:"ownContent"`
	} `xml:"func"`
}

type applicationTaskStatusReponse struct {
	autorest.Response
	IsRunning       bool
	Name            string
	DisplayName     string
	Filename        string
	Store           string
	OpCode          string
	StCode          string
	Class           string
	Category        string
	Version         string
	DownloadPercent string
	Operation       string
	IsUpdate        bool
}
