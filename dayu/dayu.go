package dayu

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var (
	appKey     string
	appSecret  string
	apiURL     string
	signName   string
	smsID      string
	sendSms    string
	callTTS    string
	callVoice  string
	callDouble string
	msgConsume string
	msgConfirm string
)

type Cfg struct {
	AppKey     string
	AppSecret  string
	UseHTTPS   bool
	HTTPURL    string
	HTTPSURL   string
	SignName   string
	SmsID      string
	SendSms    string
	CallTTS    string
	CallVoice  string
	CallDouble string
	MsgConsume string
	MsgConfirm string
}

func LoadCfg(cfg *Cfg) {

	appKey = cfg.AppKey
	appSecret = cfg.AppSecret
	if cfg.UseHTTPS {
		apiURL = cfg.HTTPSURL
	} else {
		apiURL = cfg.HTTPURL
	}
	signName = cfg.SignName
	smsID = cfg.SmsID
	sendSms = cfg.SendSms
	callTTS = cfg.CallTTS
	callVoice = cfg.CallVoice
	callDouble = cfg.CallDouble
	msgConsume = cfg.MsgConsume
	msgConfirm = cfg.MsgConfirm
}
func signBody(m map[string]string) (reader io.Reader, size int64) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	v := url.Values{}
	signString := appSecret
	for _, k := range keys {
		v.Set(k, m[k])
		signString += k + m[k]
	}
	signString += appSecret
	signByte := md5.Sum([]byte(signString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	v.Set("sign", sign)
	return ioutil.NopCloser(strings.NewReader(v.Encode())), int64(len(v.Encode()))
}
func doPost(m map[string]string) (success bool, response string) {
	body, size := signBody(m)
	client := &http.Client{}
	var req *http.Request
	var err error
	req, err = http.NewRequest("POST", apiURL, body)
	if err != nil {
		return false, err.Error()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = size
	resp, err := client.Do(req)
	if err != nil {
		response = err.Error()
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	response = string(data)
	if strings.Contains(response, "success") {
		return true, response
	}
	return false, response
}
func SendSMS(rec_num, sms_param string) (success bool, response string) {
	// || sms_free_sign_name == "" || sms_template_code == "" {
	if rec_num == "" {
		return false, "Parameter not complete"
	}
	params := make(map[string]string)
	params["app_key"] = appKey
	params["format"] = "json"
	params["method"] = sendSms
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["sms_type"] = "normal"
	params["sms_free_sign_name"] = signName
	params["rec_num"] = rec_num
	params["sms_template_code"] = smsID
	params["sms_param"] = sms_param
	return doPost(params)
}
func Consume() (success bool, response string) {
	params := make(map[string]string)
	params["app_key"] = appKey
	params["format"] = "json"
	params["method"] = msgConsume
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["quantity"] = "200"
	return doPost(params)
}
func Confirm(sids string) (success bool, response string) {
	params := make(map[string]string)
	params["app_key"] = appKey
	params["format"] = "json"
	params["method"] = msgConfirm
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["s_message_ids"] = sids
	return doPost(params)
}
func CallDouble(caller_num, caller_show_num, called_num, called_show_num string) (success bool, response string) {
	if caller_num == "" || caller_show_num == "" || called_num == "" || called_show_num == "" {
		return false, "Parameter not complete"
	}
	params := make(map[string]string)
	params["app_key"] = appKey
	params["format"] = "json"
	params["method"] = callVoice
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["caller_num"] = caller_num
	params["caller_show_num"] = caller_show_num
	params["called_num"] = called_num
	params["called_show_num"] = called_show_num
	return doPost(params)
}
func CallTTS(called_num, called_show_num, tts_code, tts_param string) (success bool, response string) {
	if called_num == "" || called_show_num == "" || tts_code == "" {
		return false, "Parameter not complete"
	}
	params := make(map[string]string)
	params["app_key"] = appKey
	params["format"] = "json"
	params["method"] = callTTS
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["called_show_num"] = called_show_num
	params["called_num"] = called_num
	params["tts_code"] = tts_code
	params["tts_param"] = tts_param
	return doPost(params)
}
func CallVoice(called_num, called_show_num, voice_code string) (success bool, response string) {
	if called_num == "" || called_show_num == "" || voice_code == "" {
		return false, "Parameter not complete"
	}
	params := make(map[string]string)
	params["app_key"] = appKey
	params["format"] = "json"
	params["method"] = callVoice
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["called_show_num"] = called_show_num
	params["called_num"] = called_num
	params["voice_code"] = voice_code
	return doPost(params)
}
