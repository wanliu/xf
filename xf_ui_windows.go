package xf

import (
	"C"
	"syscall"
	"unsafe"
)

type Event struct {
	event uintptr
}

type Listener struct {
	clListener uintptr
}

type Agent struct {
	agent  uintptr
	Events chan *Event
}

type Message struct {
	message uintptr
}

type Buffer struct {
	buffer uintptr
}

type DataBundle struct {
	bundle uintptr
}

var (
	procCreateListener      = xunfei.NewProc("xfuiCreateListener")
	procCreateAgent         = xunfei.NewProc("xfuiCreateAgent")
	procBuildMessage        = xunfei.NewProc("xfuiBuildMessage")
	procSendMessage         = xunfei.NewProc("xfuiSendMessage")
	procDestroyMessage      = xunfei.NewProc("xfuiDestroyMessage")
	procGetEventType        = xunfei.NewProc("xfuiGetEventType")
	procGetEventInfo        = xunfei.NewProc("xfuiGetEventInfo")
	procGetEventArg1        = xunfei.NewProc("xfuiGetEventArg1")
	procGetEventArg2        = xunfei.NewProc("xfuiGetEventArg2")
	procGetEventData        = xunfei.NewProc("xfuiGetEventData")
	procDestroyDataBundle   = xunfei.NewProc("xfuiDestroyDataBundle")
	procDataBundleRemove    = xunfei.NewProc("xfuiDataBundleRemove")
	procDataBundlePutInt    = xunfei.NewProc("xfuiDataBundlePutInt")
	procDataBundleGetInt    = xunfei.NewProc("xfuiDataBundleGetInt")
	procDataBundlePutLong   = xunfei.NewProc("xfuiDataBundlePutLong")
	procDataBundleGetLong   = xunfei.NewProc("xfuiDataBundleGetLong")
	procDataBundlePutString = xunfei.NewProc("xfuiDataBundlePutString")
	procDataBundleGetString = xunfei.NewProc("xfuiDataBundleGetString")
	procDataBundlePutBinary = xunfei.NewProc("xfuiDataBundlePutBinary")
	procDataBundleGetBinary = xunfei.NewProc("xfuiDataBundleGetBinary")
	procAllocBuffer         = xunfei.NewProc("xfuiAllocBuffer")
	procDeallocBuffer       = xunfei.NewProc("xfuiDeallocBuffer")
	procBufferGetData       = xunfei.NewProc("xfuiBufferGetData")
	procBufferGetSize       = xunfei.NewProc("xfuiBufferGetSize")
)

func eventListnerCallback(idx, v uintptr) uintptr {
	fn := lookup(int(idx))
	fn(&Event{event: v})
	return 0
}

func NewListener(handle func(*Event)) *Listener {
	var listener Listener
	i := register(handle)
	if _, _, err := procCreateListener.Call(
		uintptr(i),
		syscall.NewCallback(eventListnerCallback),
		uintptr(unsafe.Pointer(&listener))); err != nil {
	}
	// li := C.createListener(C.int(i), (C.CallbackFcn)(unsafe.Pointer(C.callEventListener)))
	// listener.listener = li

	return &listener
}

func NewAgent(params string, li *Listener) *Agent {
	var agent Agent

	procCreateAgent.Call(
		uintptr(unsafeString(params)),
		uintptr(unsafe.Pointer(li)),
		uintptr(unsafe.Pointer(&agent)))
	return &agent
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
	procSendMessage.Call(uintptr(unsafe.Pointer(agt)), uintptr(unsafe.Pointer(msg)))
	return nil
}

func NewMessage(msgType MsgType) *Message {
	var msg Message

	_, _, _ = procBuildMessage.Call(
		uintptr(msgType),
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringBytePtr(""))),
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(&msg)))
	return &msg
}

func NewMessageText(msgType MsgType, text string) *Message {
	var (
		msg  Message
		data = []byte(text)
	)

	procBuildMessage.Call(
		uintptr(msgType),
		uintptr(0),
		uintptr(0),
		uintptr(unsafeString("data_type=text")),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&msg)))
	return &msg
}

func NewMessageBytes(msgType MsgType, params string, buf []byte) *Message {
	var (
		msg  Message
		bufp uintptr
	)

	if len(buf) > 0 {
		bufp = uintptr(unsafe.Pointer(&buf[0]))
	}
	procBuildMessage.Call(
		uintptr(msgType),
		uintptr(0),
		uintptr(0),
		uintptr(unsafeString(params)),
		bufp,
		uintptr(len(buf)),
		uintptr(unsafe.Pointer(&msg)))

	return &msg
}

func CBool(v uintptr) bool {
	if v > 0 {
		return true
	} else {
		return false
	}
}

func Bool2Int(v bool) uintptr {
	if v {
		return 1
	} else {
		return 0
	}
}

func (msg *Message) Destroy() {
	procDestroyMessage.Call(uintptr(unsafe.Pointer(msg)))
}

func (evt *Event) EventType() EventType {
	r1, _, _ := procGetEventType.Call(
		evt.event,
	)

	return EventType(int(r1))
}

func (evt *Event) Info() string {
	r1, _, _ := procGetEventInfo.Call(evt.event)
	return uintptrToString(r1)
}

func (evt *Event) Arg1() int {
	r1, _, _ := procGetEventArg1.Call(evt.event)
	return int(r1)
}

func (evt *Event) Arg2() int {
	r1, _, _ := procGetEventArg2.Call(evt.event)
	return int(r1)
}

func (evt *Event) Data() *DataBundle {
	var bundle DataBundle
	procGetEventData.Call(evt.event, uintptr(unsafe.Pointer(&bundle)))
	return &bundle
}

func (db *DataBundle) Close() {
	procDestroyDataBundle.Call(uintptr(unsafe.Pointer(&db.bundle)))
}

func (db *DataBundle) Remove(key string) bool {
	r1, _, _ := procDataBundleRemove.Call(
		uintptr(unsafe.Pointer(&db.bundle)),
		uintptr(unsafeString(key)))
	return CBool(r1)
}

func (db *DataBundle) PutInt(key string, val int, replace bool) bool {
	r1, _, _ := procDataBundlePutInt.Call(
		db.bundle,
		uintptr(unsafeString(key)),
		uintptr(val),
		Bool2Int(replace))
	return CBool(r1)
}

func (db *DataBundle) GetInt(key string, defaultVal int) int {
	r1, _, _ := procDataBundleGetInt.Call(
		db.bundle,
		uintptr(unsafeString(key)),
		uintptr(defaultVal))
	return int(r1)
}

func (db *DataBundle) PutLong(key string, val int, replace bool) bool {
	r1, _, _ := procDataBundlePutLong.Call(
		db.bundle,
		uintptr(unsafeString(key)),
		uintptr(val),
		Bool2Int(replace))
	return CBool(r1)
}

func (db *DataBundle) GetLong(key string, defaultVal int) int {
	r1, _, _ := procDataBundleGetLong.Call(
		db.bundle,
		uintptr(unsafeString(key)),
		uintptr(defaultVal))
	return int(r1)
}

func (db *DataBundle) PutString(key string, val string, replace bool) bool {
	r1, _, _ := procDataBundlePutString.Call(
		db.bundle,
		uintptr(unsafeString(key)),
		uintptr(unsafeString(val)),
		Bool2Int(replace))
	return CBool(r1)
}

func (db *DataBundle) GetString(key string, defaultVal string) string {
	r1, _, _ := procDataBundleGetString.Call(
		db.bundle,
		uintptr(unsafeString(key)),
		uintptr(unsafeString(defaultVal)),
	)
	return uintptrToString(r1)
}

func (db *DataBundle) PutBinary(key string, val *Buffer, replace bool) bool {
	r1, _, _ := procDataBundlePutBinary.Call(
		db.bundle,
		uintptr(unsafeString(key)),
		val.buffer,
		uintptr(unsafe.Pointer(val.buffer)),
		Bool2Int(replace))

	return CBool(r1)
}

func (db *DataBundle) GetBinary(key string) *Buffer {
	var buffer Buffer
	procDataBundleGetBinary.Call(
		uintptr(unsafe.Pointer(db)),
		uintptr(unsafeString(key)),
		uintptr(unsafe.Pointer(&buffer)))
	return &buffer
}

func (buf *Buffer) Data() []byte {
	size, _, _ := procBufferGetSize.Call(uintptr(unsafe.Pointer(buf)))
	r1, _, _ := procBufferGetData.Call(uintptr(unsafe.Pointer(buf)), uintptr(size))
	return uintptrToBytes(r1, int(size))
}
