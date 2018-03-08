package main

/*
#include <stdio.h>
#include "xf_ui.h"

// The gateway function
void callEventListener(Event *evt)
{
	int goEventListner(Event*);
	goEventListner(evt);
}

*/
import "C"
import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/wanliu/xf"
)

//export goEventListner
func goEventListner(v *C.Event) int {
	return 0
}

func main() {
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

	agent := xf.NewAgent(string(buf))
	log.Printf("Agent %#v", agent)
	// agent.Start()
	time.Sleep(time.Second)
	agent.Weakup()

	msg := xf.NewMessageText(xf.CmdWrite, "你好")
	agent.SendMessage(msg)
	msg.Destroy()
}
