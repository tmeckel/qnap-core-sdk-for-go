package auth

import (
	"encoding/xml"

	"github.com/Azure/go-autorest/autorest"
)

const fqdn = "github.com/qnap/core-sdk-for-go/services/login"

type qdocRoot struct {
	XMLName      xml.Name `xml:"QDocRoot"`
	Text         string   `xml:",chardata"`
	Version      string   `xml:"version,attr"`
	DoQuick      string   `xml:"doQuick"`
	IsBooting    string   `xml:"is_booting"`
	MediaReady   string   `xml:"mediaReady"`
	ShutdownInfo struct {
		Text      string `xml:",chardata"`
		Type      string `xml:"type"`
		Timestamp string `xml:"timestamp"`
		Duration  string `xml:"duration"`
	} `xml:"shutdown_info"`
	SMBFW           string `xml:"SMBFW"`
	HeroModel       string `xml:"hero_model"`
	QtsModeType     string `xml:"qts_mode_type"`
	AuthPassed      int    `xml:"authPassed"`
	AuthSid         string `xml:"authSid"`
	PwStatus        string `xml:"pw_status"`
	IsAdmin         string `xml:"isAdmin"`
	Username        string `xml:"username"`
	Groupname       string `xml:"groupname"`
	Ts              string `xml:"ts"`
	FwNotice        string `xml:"fwNotice"`
	SUID            string `xml:"SUID"`
	Title           string `xml:"title"`
	Content         string `xml:"content"`
	PsType          string `xml:"psType"`
	StandardMassage string `xml:"standard_massage"`
	StandardColor   string `xml:"standard_color"`
	StandardSize    string `xml:"standard_size"`
	StandardBgStyle string `xml:"standard_bg_style"`
	ShowVersion     string `xml:"showVersion"`
	ShowLink        string `xml:"show_link"`
	Cuid            string `xml:"cuid"`
}

type LoginResponse struct {
	autorest.Response
	AuthPassed int
	Sid        string
	IsAdmin    bool
	Username   string
	Groupname  string
}
