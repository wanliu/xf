package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/wanliu/xf"
)

func main() {
	if err := xf.MSPLogin("appid = 58c51121, work_dir = ."); err != nil {
		t.Fatalf("login failed", err)
	}

	grammar, err := ioutil.ReadAll("test.anbf")
	if err != nil {
		t.Fatalf("load anbf file error: %s", err)
	}

	grammarId, err := xf.MSPUploadData("usergram", grammar, "dtt = abnf, sub = asr")
	if err != nil {
		t.Fatalf("MSPUploadData: %s", err)
	}

	var params = "sub = iat, domain = iat, language = zh_cn, accent = mandarin, sample_rate = 16000, result_type = plain, result_encoding = utf8"

	sessionId, err := xf.QISRSessionBegin("", params)
	if err != nil {
		t.Fatalf("session begin failed, error: %s", err)
	}

	t.Logf("sessionId: %+v", sessionId)
	f, err := os.Open("test.wav")
	if err != nil {
		t.Fatalf("open wave file failed %s", err)
	}

	r := bufio.NewReader(f)
	buf := make([]byte, 10*FRAME_LEN)

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

			ep_stat, _, err := xf.QISRAudioWrite(sessionId, rbuf, aud_stat)
			if err != nil {
				t.Logf("Audio Write Failed %s", err)
				break
			}

			if ep_stat == MSP_EP_AFTER_SPEECH {
				break
			}

			time.Sleep(200 * time.Millisecond)
		}

		_, _, err := xf.QISRAudioWrite(sessionId, nil, MSP_AUDIO_SAMPLE_LAST)
		return err
	}

	{
		if err := sendData(); err != nil {
			t.Logf("发送数据失败: %s", err)
			goto iat_exit
		}

		rslt, err := get_audio_result(sessionId)
		if err != nil {
			goto iat_exit
		}

		t.Logf("result: %s", rslt)
	}
iat_exit:

	if err := xf.QISRSessionEnd(sessionId, ""); err != nil {
		t.Fatalf("session end failed, error: %s", err)
	}

	if err := xf.MSPLogout(); err != nil {
		t.Fatalf("logout failed", err)
	}
}
