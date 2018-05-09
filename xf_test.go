package xf

import (
	"bufio"
	"log"
	"os"
	"testing"
	"time"
	"fmt"
)

const (
	BUFFER_SIZE = 4096
	FRAME_LEN   = 640
	HINTS_SIZE  = 100
)

var appId = "58c51121"

func TestXF(t *testing.T) {
	params := fmt.Sprintf("appid = %s, work_dir = .", appId)
	if err := MSPLogin(params); err != nil {
		t.Fatalf("login failed %s", err)
	}

	if err := MSPLogout(); err != nil {
		t.Fatalf("logout failed %s", err)
	}
}

func get_audio_result(sessionId string, status int) (string, error) {
	var rslt = ""
	for {
		res, stat, err := QISRGetResult(sessionId, status, 0)
		if err != nil {
			return "", err
		}

		if stat == MSP_REC_STATUS_COMPLETE {
			break
		}

		rslt += res

		time.Sleep(150 * time.Millisecond)
	}

	return rslt, nil
}

func TestQISR(t *testing.T) {
	loginParams := fmt.Sprintf("appid = %s, work_dir = .", appId)

	if err := MSPLogin(loginParams); err != nil {
		t.Fatalf("login failed %s", err)
	}

	var params = "sub = iat, domain = iat, language = zh_cn, accent = mandarin, sample_rate = 16000, result_type = plain, result_encoding = utf8"

	sessionId, err := QISRSessionBegin("", params)
	if err != nil {
		t.Fatalf("session begin failed, error: %s", err)
	}

	t.Logf("sessionId: %+v", sessionId)
	f, err := os.Open("res/wav/iflytek02.wav")
	if err != nil {
		t.Fatalf("open wave file failed %s", err)
	}

	r := bufio.NewReader(f)
	buf := make([]byte, 10*FRAME_LEN)
	var res_stat int
	sendData := func() error {
		var aud_stat int
		for i := 0; ; i += 1 {
			n, err := r.Read(buf)
			rbuf := buf[:n]
			if err != nil {
				break
			}

			log.Printf("read %d\n", i)

			aud_stat = MSP_AUDIO_SAMPLE_CONTINUE
			if 0 == i {
				aud_stat = MSP_AUDIO_SAMPLE_FIRST
			}

			ep_stat, _, err := QISRAudioWrite(sessionId, rbuf, aud_stat)
			if err != nil {
				t.Logf("Audio Write Failed %s", err)
				break
			}

			if ep_stat == MSP_EP_AFTER_SPEECH {
				break
			}

			time.Sleep(200 * time.Millisecond)
		}

		_, res_stat, err = QISRAudioWrite(sessionId, nil, MSP_AUDIO_SAMPLE_LAST)
		return err
	}

	{
		if err := sendData(); err != nil {
			t.Logf("发送数据失败: %s", err)
			goto iat_exit
		}

		rslt, err := get_audio_result(sessionId, res_stat)
		if err != nil {
			goto iat_exit
		}

		t.Logf("result: %s", rslt)
	}
iat_exit:

	if err := QISRSessionEnd(sessionId, ""); err != nil {
		t.Fatalf("session end failed, error: %s", err)
	}

	if err := MSPLogout(); err != nil {
		t.Fatalf("logout failed %s", err)
	}
}


func init() {
	if a := os.Getenv("APPID"); len(a) > 0 {
		appId = a
	}
}