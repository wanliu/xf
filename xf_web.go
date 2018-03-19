package xf

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Scene struct {
	Scene  string `json:"scene"`
	UserId string `json:"userid"`
}

type IATScene struct {
	AUF   string `json:"auf"`
	AUE   string `json:"aue"`
	Scene string `json:"scene"`
}

type VoiceScene struct {
	AUF    string `json:"auf"`
	AUE    string `json:"aue"`
	Scene  string `json:"scene"`
	UserId string `json:"userid"`
}

type WebClient struct {
	AppId   string
	AppKey  string
	apiHost string
	client  *http.Client
}

type WebSession struct {
	Scene       string
	UserId      string
	client      *WebClient
	form        url.Values
	sampleRate  string
	audioFormat string
	curType     ParamType
}

type ApiURL string

const (
	ApiText  = "v1/aiui/v1/text_semantic"
	ApiIAT   = "v1/aiui/v1/iat"
	ApiVoice = "v1/aiui/v1/voice_semantic"
)

type ParamType int

const (
	TextParam ParamType = iota
	IATParam
	VoiceParam
)

func btoa(buf []byte) string {
	return base64.StdEncoding.EncodeToString(buf)
}

func Query(appId, apikey, scene, user_id, text string) ([]byte, error) {
	var client http.Client
	form := url.Values{}
	form.Add("text", btoa([]byte(text)))

	req, err := http.NewRequest("POST", "https://api.xfyun.cn/v1/aiui/v1/text_semantic", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	p := Scene{
		Scene:  scene,
		UserId: user_id,
	}

	buf, err := json.Marshal(&p)
	if err != nil {
		return nil, err
	}

	t := time.Now()
	cur_time := fmt.Sprintf("%d", t.Unix())
	param := btoa(buf)
	h := md5.New()
	io.WriteString(h, apikey+cur_time+param+form.Encode())
	log.Printf("%s", appId+cur_time+param+form.Encode())
	check_sum := fmt.Sprintf("%x", h.Sum(nil))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("X-Appid", appId)
	req.Header.Set("X-CurTime", cur_time)
	req.Header.Set("X-Param", param)
	req.Header.Set("X-CheckSum", check_sum)
	log.Printf("Content-Type: %s\n", "application/x-www-form-urlencoded; charset=utf-8")
	log.Printf("X-Appid: %s\n", appId)
	log.Printf("X-CurTime: %s\n", cur_time)
	log.Printf("Param: %s\n", buf)
	log.Printf("X-Param: %s\n", param)
	log.Printf("Body: %s\n", form.Encode())
	log.Printf("X-CheckSum: %s\n", check_sum)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func NewWebClient(appid, appkey string) *WebClient {
	return &WebClient{
		AppId:   appid,
		AppKey:  appkey,
		apiHost: "https://api.xfyun.cn/",
		client:  &http.Client{},
	}
}

func (cli *WebClient) NewSession(scene, user_id string) *WebSession {
	return &WebSession{
		Scene:  scene,
		UserId: user_id,
		client: cli,
	}
}

func (sess *WebSession) PostText(text string) ([]byte, error) {
	sess.setParamType(TextParam)
	sess.form = url.Values{}
	sess.form.Set("text", btoa([]byte(text)))

	req, err := sess.NewReq(sess.apiUrl(), strings.NewReader(sess.XBody()))

	if err != nil {
		return nil, err
	}

	cur_time := sess.XCurTime()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("X-Appid", sess.client.AppId)
	req.Header.Set("X-CurTime", cur_time)
	req.Header.Set("X-Param", sess.XParam())
	req.Header.Set("X-CheckSum", sess.XCheckSum(cur_time))
	log.Printf("X-Appid: %s\n", sess.client.AppId)
	log.Printf("X-CurTime: %s\n", cur_time)
	log.Printf("X-Param: %s\n", sess.XParam())
	// log.Printf("Body: %s\n", sess.XBody())
	log.Printf("X-CheckSum: %s\n", sess.XCheckSum(cur_time))
	log.Printf("API URL: %s", sess.apiUrl())
	resp, err := sess.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (sess *WebSession) PostIAT(buf []byte) ([]byte, error) {
	sess.setParamType(IATParam)

	sess.form = url.Values{}
	sess.form.Set("data", btoa(buf))

	req, err := sess.NewReq(sess.apiUrl(), strings.NewReader(sess.XBody()))

	if err != nil {
		return nil, err
	}

	cur_time := sess.XCurTime()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("X-Appid", sess.client.AppId)
	req.Header.Set("X-CurTime", cur_time)
	req.Header.Set("X-Param", sess.XParam())
	req.Header.Set("X-CheckSum", sess.XCheckSum(cur_time))
	log.Printf("X-Appid: %s\n", sess.client.AppId)
	log.Printf("X-CurTime: %s\n", cur_time)
	log.Printf("X-Param: %s\n", sess.XParam())
	// log.Printf("Body: %s\n", sess.XBody())
	log.Printf("X-CheckSum: %s\n", sess.XCheckSum(cur_time))
	log.Printf("API URL: %s", sess.apiUrl())
	resp, err := sess.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (sess *WebSession) PostVoice(buf []byte) ([]byte, error) {
	sess.setParamType(VoiceParam)

	sess.form = url.Values{}
	sess.form.Set("data", btoa(buf))

	req, err := sess.NewReq(sess.apiUrl(), strings.NewReader(sess.XBody()))

	if err != nil {
		return nil, err
	}

	cur_time := sess.XCurTime()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("X-Appid", sess.client.AppId)
	req.Header.Set("X-CurTime", cur_time)
	req.Header.Set("X-Param", sess.XParam())
	req.Header.Set("X-CheckSum", sess.XCheckSum(cur_time))
	log.Printf("X-Appid: %s\n", sess.client.AppId)
	log.Printf("X-CurTime: %s\n", cur_time)
	log.Printf("X-Param: %s\n", sess.XParam())
	// log.Printf("Body: %s\n", sess.XBody())
	log.Printf("X-CheckSum: %s\n", sess.XCheckSum(cur_time))
	log.Printf("API URL: %s", sess.apiUrl())
	resp, err := sess.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (sess *WebSession) NewReq(path string, form io.Reader) (*http.Request, error) {
	return http.NewRequest("POST", path, form)
}

func (sess *WebSession) apiUrl() string {
	var api string
	switch sess.curType {
	case TextParam:
		api = ApiText
	case IATParam:
		api = ApiIAT
	case VoiceParam:
		api = ApiVoice
	}
	return fmt.Sprintf("%s%s", sess.client.apiHost, api)
}

func (sess *WebSession) XCurTime() string {
	t := time.Now()
	return fmt.Sprintf("%d", t.Unix())
}

func (sess *WebSession) XBody() string {
	switch sess.curType {
	case TextParam:
		return fmt.Sprintf("text=%s", sess.form.Get("text"))
	case IATParam, VoiceParam:
		return fmt.Sprintf("data=%s", sess.form.Get("data"))
	}
	return ""
}

func (sess *WebSession) XCheckSum(time string) string {
	h := md5.New()
	s := fmt.Sprintf("%s%s%s%s", sess.client.AppKey, time, sess.XParam(), sess.XBody())
	io.WriteString(h, s)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (sess *WebSession) setParamType(pt ParamType) {
	sess.curType = pt
}

func (sess *WebSession) SetAudioFormat(sampleRate, format string) {
	sess.audioFormat = format
	sess.sampleRate = sampleRate
}

func (sess *WebSession) XParam() string {
	var p interface{}
	// log.Printf("curType: %d", sess.curType)
	switch sess.curType {
	case TextParam:
		p = Scene{
			Scene:  sess.Scene,
			UserId: sess.UserId,
		}
	case IATParam:
		p = IATScene{
			AUF:   sess.sampleRate,
			AUE:   sess.audioFormat,
			Scene: sess.Scene,
		}
	case VoiceParam:
		p = VoiceScene{
			AUF:    sess.sampleRate,
			AUE:    sess.audioFormat,
			Scene:  sess.Scene,
			UserId: sess.UserId,
		}
	}
	// log.Printf("params %#v", p)

	buf, _ := json.Marshal(&p)
	// log.Printf("params json: %s", buf)
	return btoa(buf)
}
