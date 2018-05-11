package xf

/*
#cgo CFLAGS:-g -Wall -I./include
#cgo LDFLAGS:-L./libs/x64 -lxf_ui -laiui -lstdc++
#include "xf_ui.h"

int callEventListener(Event *evt); // Forward declaration.
*/
import "C"
import (
	"unsafe"
)

type Event struct {
	event *C.Event
}

type Listener struct {
	listener *C.Listener
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

//export goEventListner
func goEventListner(idx C.int, v *C.Event) {
	fn := lookup(int(idx))
	fn(&Event{event: v})
}

func NewListener(handle func(*Event)) *Listener {
	var listener Listener
	i := register(handle)
	li := C.createListener(C.int(i), (C.CallbackFcn)(unsafe.Pointer(C.callEventListener)))
	listener.listener = li

	return &listener
}

func NewAgent(params string, li *Listener) *Agent {
	agent := C.createAgent(C.CString(params), li.listener)
	return &Agent{
		agent:  agent,
		Events: make(chan *Event, 0),
	}
}

func (agt *Agent) Start() {
	msg := NewMessage(CmdStart)
	agt.SendMessage(msg)
	msg.Destroy()
}

func (agt *Agent) Stop() {
	msg := NewMessage(CmdStop)
	agt.SendMessage(msg)
	msg.Destroy()
}

func (agt *Agent) Weakup() {
	msg := NewMessage(CmdWakeup)
	agt.SendMessage(msg)
	msg.Destroy()
}

func (agt *Agent) Reset() {
	msg := NewMessage(CmdResetWakeup)
	agt.SendMessage(msg)
	msg.Destroy()
}

func (agt *Agent) SendMessage(msg *Message) error {
	C.sendMessage(agt.agent, msg.message)
	return nil
}

func NewMessage(msgType MsgType) *Message {

	return &Message{
		message: C.buildMessage(C.int(msgType), 0, 0, C.CString(""), nil, 0),
	}
}

func NewMessageText(msgType MsgType, text string) *Message {
	data := []byte(text)
	return &Message{
		message: C.buildMessage(C.int(msgType), 0, 0, C.CString("data_type=text"), C.CBytes(data), C.int(len(data))),
	}
}

func NewMessageBytes(msgType MsgType, params string, buf []byte) *Message {
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
