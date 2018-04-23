// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package xf

import "strconv"

const (
	_EventType_name_0 = "EventResultEventErrorEventStateEventWakeupEventSleepEventVadEventBindSucessEventCmdReturnEventAudioEventPreSleepEventStartRecordEventStopRecordEventConnectedToServerEventServerDisconnected"
	_EventType_name_1 = "EventCaePlainTextEventPushMessage"
)

var (
	_EventType_index_0 = [...]uint8{0, 11, 21, 31, 42, 52, 60, 75, 89, 99, 112, 128, 143, 165, 188}
	_EventType_index_1 = [...]uint8{0, 17, 33}
)

func (i EventType) String() string {
	switch {
	case 1 <= i && i <= 14:
		i -= 1
		return _EventType_name_0[_EventType_index_0[i]:_EventType_index_0[i+1]]
	case 1000 <= i && i <= 1001:
		i -= 1000
		return _EventType_name_1[_EventType_index_1[i]:_EventType_index_1[i+1]]
	default:
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}