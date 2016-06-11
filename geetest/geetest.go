package geetest

import (
	"github.com/GeeTeam/GtGoSdk"
)

const (
	GEETEST_ID  = "242609955f2a65b5b2ad643ad55152f5"
	GEETEST_KEY = "75f37ba1d3db62977e45842162d9adea"
)

type Captcha struct {
	GeetestID string `json:"geetestId,omitempty"`
	CaptchaID string `json:"captchaId"`
	Mode      int    `json:"mode"` // 1 indicates normal mode and 0 indicates fallback mode
	Key       string `json:"key,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

func NewCaptcha(userID string) *Captcha {
	lib := GtGoSdk.GeetestLib(GEETEST_KEY, GEETEST_ID)
	mode := lib.PreProcess(userID)
	return &Captcha{
		CaptchaID: lib.GetResponseMap()["challenge"].(string),
		GeetestID: GEETEST_ID,
		Mode:      mode,
	}
}

func (captcha *Captcha) Validate(userID string) bool {
	lib := GtGoSdk.GeetestLib(GEETEST_KEY, GEETEST_ID)
	if captcha.Mode == 1 {
		return lib.SuccessValidate(captcha.CaptchaID, captcha.Hash, captcha.Key, userID)
	}
	return lib.FailbackValidate(captcha.CaptchaID, captcha.Hash, captcha.Key)
}
