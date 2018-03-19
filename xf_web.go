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

type WebClient struct {
	AppId   string
	AppKey  string
	apiHost string
	client  *http.Client
}

type WebSession struct {
	Scene  string
	UserId string
	client *WebClient
	form   url.Values
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

type ApiURL string

const (
	ApiText   = "v1/aiui/v1/text_semantic"
	ApiSpeech = "v1/aiui/v1/iat"
)

func (sess *WebSession) PostText(text string) ([]byte, error) {
	sess.form = url.Values{}
	sess.form.Add("text", btoa([]byte(text)))

	req, err := sess.NewReq(sess.apiUrl(ApiText), strings.NewReader(sess.XBody()))

	if err != nil {
		return nil, err
	}

	cur_time := sess.XCurTime()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("X-Appid", sess.client.AppId)
	req.Header.Set("X-CurTime", cur_time)
	req.Header.Set("X-Param", string(sess.XParam()))
	req.Header.Set("X-CheckSum", sess.XCheckSum(cur_time))
	log.Printf("X-Appid: %s\n", sess.client.AppId)
	log.Printf("X-CurTime: %s\n", cur_time)
	log.Printf("X-Param: %s\n", sess.XParam())
	log.Printf("Body: %s\n", sess.XBody())
	log.Printf("X-CheckSum: %s\n", sess.XCheckSum(cur_time))
	log.Printf("API URL: %s", sess.apiUrl(ApiText))
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

func (sess *WebSession) apiUrl(api ApiURL) string {
	return fmt.Sprintf("%s%s", sess.client.apiHost, api)
}

func (sess *WebSession) XCurTime() string {
	t := time.Now()
	return fmt.Sprintf("%d", t.Unix())
}

func (sess *WebSession) XBody() string {
	return sess.form.Encode()
}

func (sess *WebSession) XCheckSum(time string) string {
	h := md5.New()
	s := fmt.Sprintf("%s%s%s%s", sess.client.AppKey, time, sess.XParam(), sess.XBody())
	io.WriteString(h, s)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (sess *WebSession) XParam() string {
	p := Scene{
		Scene:  sess.Scene,
		UserId: sess.UserId,
	}

	buf, _ := json.Marshal(&p)
	return btoa(buf)
}
