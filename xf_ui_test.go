package xf

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestXFUI(t *testing.T) {
	li := CreateListener(func(evt *Event) {

	})

	file, err := os.Open("aiui.cfg") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(file)
	var cfg map[string]interface{}
	for {
		if err := dec.Decode(&cfg); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	cfg["login"] = map[string]interface{}{
		"appid": "58c51121",
	}

	buf, err := json.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	agent := CreateAgent(string(buf), li)
	t.Logf("Listener %#v", li)
	t.Logf("Agent %#v", agent)
	// agent.Start()
	time.Sleep(time.Second)
	agent.Weakup()

	msg := buildMessageText(CmdWrite, "你好")
	agent.SendMessage(msg)
	msg.Destroy()

	time.Sleep(3 * time.Second)
}

func TestXFUI_Wav(t *testing.T) {
	li := CreateListener(func(evt *Event) {

	})

	cfile, err := os.Open("aiui.cfg") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(cfile)
	var cfg map[string]interface{}
	for {
		if err := dec.Decode(&cfg); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	cfg["login"] = map[string]interface{}{
		"appid": "58c51121",
	}

	buf, err := json.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	agent := CreateAgent(string(buf), li)
	time.Sleep(time.Second)
	agent.Weakup()

	file, err := os.Open("test/test.pcm")
	if err != nil {
		log.Fatalf("打开测试文件失败 %s", err)
	}

	var pcm = make([]byte, 1279)
	for _, err := file.Read(pcm); err != io.EOF; _, err = file.Read(pcm) {
		msg := buildMessageBytes(CmdWrite, "data_type=audio,sample_rate=16000", pcm)
		agent.SendMessage(msg)
		msg.Destroy()
		time.Sleep(10 * time.Millisecond)
	}

	msg := buildMessageBytes(CmdStopWrite, "data_type=audio,sample_rate=16000", nil)
	agent.SendMessage(msg)
	msg.Destroy()
	time.Sleep(3 * time.Second)
}
