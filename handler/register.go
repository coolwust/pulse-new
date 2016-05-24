package handler

import (
	"net/http"
	"html/template"
	"io"
	"github.com/GeeTeam/GtGoSdk"
)

const (
	GEETEST_ID = "242609955f2a65b5b2ad643ad55152f5"
	GEETEST_KEY = "75f37ba1d3db62977e45842162d9adea"
)

func Register(w http.ResponseWriter, req *http.Request) {
	template.Must(template.ParseFiles("tmpl/register.tmpl")).Execute(w, nil)
}

func GeetestInit(w http.ResponseWriter, _ *http.Request) {
	gt := GtGoSdk.GeetestLib(GEETEST_KEY, GEETEST_ID)
	gt.PreProcess("")
	io.WriteString(w, gt.GetResponseStr())
}
