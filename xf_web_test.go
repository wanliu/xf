package xf

import (
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestXFUI_Web(t *testing.T) {
	client := NewWebClient("5aaf5fe1", "629e95d1bd0846b289212576a26d3fb9")
	sess := client.NewSession("main", "user_01")

	if buf, err := sess.PostText("你好"); err != nil {
		t.Fatalf("error %s", err)
	} else {
		t.Logf("result: %s\n", buf)
	}

	time.Sleep(time.Second)
	log.Println("--------------------------")
	data, err := ioutil.ReadFile("test/16kVoice.pcm")
	if err != nil {
		log.Fatalf("打开测试文件失败 %s", err)
	}

	t.Logf("read wave %d length", len(data))

	sess.SetAudioFormat("16k", "raw")
	buf, err := sess.PostIAT(data)
	if err != nil {
		t.Fatalf("error %s", err)
	} else {
		t.Logf("result: %s\n", buf)
	}

	sess.SetAudioFormat("16k", "raw")
	buf, err = sess.PostVoice(data)
	if err != nil {
		t.Fatalf("error %s", err)
	} else {
		t.Logf("result: %s\n", buf)
	}
}
