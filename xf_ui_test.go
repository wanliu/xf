package xf

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func receiveEvent(evt *Event) {
	// log.Printf("evt: %#v", evt)

	switch evt.EventType() {
	case EventState:
		state := (StateType)(evt.Arg1())
		switch state {
		case StateIdle:
			log.Printf("EventState: %s\n", "IDLE")
		case StateReady:
			log.Printf("EventState: %s\n", "READY")
		case StateWorking:
			log.Printf("EventState: %s\n", "WORKING")
		default:
			log.Printf("InvalidState %d", state)
		}
	case EventVad:
		vad := (VadType)(evt.Arg1())

		switch vad {
		case VadBos:
			log.Printf("VadState: %s\n", "检测到开始")
		case VadEos:
			log.Printf("VadState: %s\n", "检测到结束")
		case VadVol:
			//
		}
	case EventResult:
		info := evt.Info()
		var (
			result AIUIResult
			data   AIUIResultData
		)

		if err := json.Unmarshal([]byte(info), &result); err != nil {
			log.Printf("解析 Result 错误: %s", err)
		}

		data = result.Data[0]

		if data.Params.Sub == "nlp" {
			contentId := data.Content[0].CNT_ID
			if len(contentId) == 0 {
				log.Printf("missing contentId")
			}

			buffer := evt.Data().GetBinary(contentId)
			var out bytes.Buffer
			var buff = buffer.Data()
			buff = buff[:len(buff)-1]
			if err := json.Indent(&out, buff, "", "  "); err != nil {
				log.Printf("格式化内容错误: %s", err)
			}

			log.Printf("output: %s", out.String())
		}
	case EventError:
		log.Printf("Error Code : %d", evt.Arg1())
	default:
		log.Printf("消息 EventType: %d", evt.EventType())
	}
}

func TestXFUI(t *testing.T) {
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

	listener := NewListener(receiveEvent)
	log.Printf("Listener %#v", listener)

	agent := NewAgent(string(buf), listener)
	log.Printf("Agent %#v", agent)
	// agent.Start()
	time.Sleep(time.Second)
	agent.Weakup()

	msg := NewMessageText(CmdWrite, "你好")
	agent.SendMessage(msg)
	msg.Destroy()
	time.Sleep(3 * time.Second)
}

func TestXFUI_Wav(t *testing.T) {
	li := NewListener(receiveEvent)

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
		"appid": appId,
	}

	buf, err := json.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	agent := NewAgent(string(buf), li)
	time.Sleep(time.Second)
	agent.Weakup()

	file, err := os.Open("test/test.pcm")
	if err != nil {
		log.Fatalf("打开测试文件失败 %s", err)
	}

	var pcm = make([]byte, 1279)
	for _, err := file.Read(pcm); err != io.EOF; _, err = file.Read(pcm) {
		msg := NewMessageBytes(CmdWrite, "data_type=audio,sample_rate=16000", pcm)
		agent.SendMessage(msg)
		msg.Destroy()
		time.Sleep(10 * time.Millisecond)
	}

	msg := NewMessageBytes(CmdStopWrite, "data_type=audio,sample_rate=16000", nil)
	agent.SendMessage(msg)
	msg.Destroy()
	time.Sleep(3 * time.Second)
}

func init() {
	if a := os.Getenv("APPID"); len(a) > 0 {
		appId = a
	}
}
