package xf

/*
#cgo CFLAGS:-g -Wall -I./include
#cgo LDFLAGS:-L./libs/x64 -lxf_ui -laiui -lstdc++
#include "xf_ui.h"

int callEventListener(Event *evt); // Forward declaration.
*/
import "C"
import (
	"bytes"
	"encoding/json"
	"log"
	"unsafe"
)

type Event struct {
	event *C.Event
}

type Listener struct {
	listener *C.Listener
	agent    *Agent
}

type Agent struct {
	agent  *C.Agent
	Events chan *Event
}

type Message struct {
	message *C.Message
}

type Buffer struct {
	buffer *C.MessageBuffer
}

type EventHandler func(evt *Event) error

type Map map[string]interface{}

//export goEventListner
func goEventListner(v *C.Event) int {
	var evt = &Event{
		event: v,
	}

	// handler := C.getEventHandler(v)
	// t := reflect.TypeOf(Listener{})
	// vv := reflect.NewAt(t, handler)
	// log.Printf("event handler %v %s", vv, vv.Type().String())
	// listener, ok := vv.Interface().(*Listener)
	// if !ok {
	// 	log.Printf("listener convert failed")
	// }

	// log.Printf("listener agent %# v", listener)
	// if listener.agent != nil {
	// 	listener.agent.Events <- evt
	// }

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
				return 0
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
	return 0
}

func CreateListener(handle func(evt *Event)) *Listener {
	var listener Listener
	li := C.createListener(unsafe.Pointer(&listener), (C.CallbackFcn)(unsafe.Pointer(C.callEventListener)))
	listener.listener = li

	return &listener
}

func CreateAgent(params string, listener *Listener) *Agent {
	agent := C.createAgent(C.CString(params), listener.listener)
	var a = Agent{
		agent:  agent,
		Events: make(chan *Event, 0),
	}
	listener.agent = &a
	return &a
}

func NewListener() *Listener {
	var listener Listener
	li := C.createListener(unsafe.Pointer(&listener), (C.CallbackFcn)(unsafe.Pointer(C.callEventListener)))
	listener.listener = li

	return &listener
}

func NewAgent(params string) *Agent {
	listener := NewListener()
	agent := C.createAgent(C.CString(params), listener.listener)
	var a = Agent{
		agent:  agent,
		Events: make(chan *Event, 0),
	}
	listener.agent = &a
	return &a
}

func (agt *Agent) Start() {
	msg := buildMessage(CmdStart)
	agt.SendMessage(msg)
	msg.Destroy()
}

func (agt *Agent) Stop() {
	msg := buildMessage(CmdStop)
	agt.SendMessage(msg)
	msg.Destroy()
}

func (agt *Agent) Weakup() {
	msg := buildMessage(CmdWakeup)
	agt.SendMessage(msg)
	msg.Destroy()
}

func (agt *Agent) SendMessage(msg *Message) error {
	C.sendMessage(agt.agent, msg.message)
	return nil
}

func buildMessage(msgType MsgType) *Message {

	return &Message{
		message: C.buildMessage(C.int(msgType), 0, 0, C.CString(""), nil, 0),
	}
}

func buildMessageText(msgType MsgType, text string) *Message {
	data := []byte(text)
	return &Message{
		message: C.buildMessage(C.int(msgType), 0, 0, C.CString("data_type=text"), C.CBytes(data), C.int(len(data))),
	}
}

func buildMessageBytes(msgType MsgType, params string, buf []byte) *Message {
	return &Message{
		message: C.buildMessage(C.int(msgType), 0, 0, C.CString(params), C.CBytes(buf), C.int(len(buf))),
	}
}

func (msg *Message) Destroy() {
	C.destroyMessage(msg.message)
}

func (evt *Event) EventType() EventType {
	return (EventType)(C.getEventType(evt.event))
}

func (evt *Event) Info() string {
	return C.GoString(C.getEventInfo(evt.event))
}

func (evt *Event) Arg1() int {
	return int(C.getEventArg1(evt.event))
}

func (evt *Event) Arg2() int {
	return int(C.getEventArg2(evt.event))
}

func (evt *Event) Data() *DataBundle {
	var bundle DataBundle
	bundle.bundle = C.getEventData(evt.event)
	return &bundle
}

type DataBundle struct {
	bundle *C.DataBundle
}

func CBool(v C.int) bool {
	if v > 0 {
		return true
	} else {
		return false
	}
}

func Bool2Int(v bool) C.int {
	if v {
		return 1
	} else {
		return 0
	}
}

func (db *DataBundle) Close() {
	C.destroyDataBundle(db.bundle)
}

func (db *DataBundle) Remove(key string) bool {
	return CBool(C.DataBundleRemove(db.bundle, C.CString(key)))
}

func (db *DataBundle) PutInt(key string, val int, replace bool) bool {
	return CBool(C.DataBundlePutInt(db.bundle, C.CString(key), C.int(val), Bool2Int(replace)))
}

func (db *DataBundle) GetInt(key string, defaultVal int) int {
	return (int)(C.DataBundleGetInt(db.bundle, C.CString(key), C.int(defaultVal)))
}

func (db *DataBundle) PutLong(key string, val int, replace bool) bool {
	return CBool(C.DataBundlePutLong(db.bundle, C.CString(key), C.long(val), Bool2Int(replace)))
}

func (db *DataBundle) GetLong(key string, defaultVal int) int {
	return (int)(C.DataBundleGetLong(db.bundle, C.CString(key), C.long(defaultVal)))
}

func (db *DataBundle) PutString(key string, val string, replace bool) bool {
	return CBool(C.DataBundlePutString(db.bundle, C.CString(key), C.CString(val), Bool2Int(replace)))
}

func (db *DataBundle) GetString(key string, defaultVal string) string {
	return C.GoString((C.DataBundleGetString(db.bundle, C.CString(key), C.CString(defaultVal))))
}

func (db *DataBundle) PutBinary(key string, val *Buffer, replace bool) bool {
	return CBool(C.DataBundlePutBinary(db.bundle, C.CString(key), (*C.MessageBuffer)(val.buffer), Bool2Int(replace)))
}

func (db *DataBundle) GetBinary(key string) *Buffer {
	var buffer Buffer
	buffer.buffer = C.DataBundleGetBinary(db.bundle, C.CString(key))
	return &buffer
}

func (buf *Buffer) Data() []byte {
	size := C.BufferGetSize(buf.buffer)
	return C.GoBytes(C.BufferGetData(buf.buffer), size)
}

func RegisterListener() {

}
