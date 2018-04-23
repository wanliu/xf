package xf

//go:generate stringer -type=EventType
//go:generate stringer -type=StateType
//go:generate stringer -type=VadType

// 获取 常量名称与数字
// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1 }' | grep '^CMD'
// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1, $4 }' | grep '^CMD' | egrep -op '\d+';
type MsgType int

//go:generate stringer -type=MsgType

const (
	CmdBuildGrammar        MsgType = 16
	CmdCaeWriteDevinfo     MsgType = 1000
	CmdCaeWriteEncrypt     MsgType = 1001
	CmdCleanDialogHistory  MsgType = 21
	CmdGetState            MsgType = 1
	CmdQuerySyncStatus     MsgType = 24
	CmdReset               MsgType = 4
	CmdResetWakeup         MsgType = 8
	CmdResultValidationAck MsgType = 20
	CmdSendLog             MsgType = 12
	CmdSetBeam             MsgType = 9
	CmdSetParams           MsgType = 10
	CmdStart               MsgType = 5
	CmdStartRecord         MsgType = 22
	CmdStartSave           MsgType = 14
	CmdStartThrowAudio     MsgType = 18
	CmdStop                MsgType = 6
	CmdStopRecord          MsgType = 23
	CmdStopSave            MsgType = 15
	CmdStopThrowAudio      MsgType = 19
	CmdStopWrite           MsgType = 3
	CmdSync                MsgType = 13
	CmdUpdateLocalLexicon  MsgType = 17
	CmdUploadLexicon       MsgType = 11
	CmdWakeup              MsgType = 7
	CmdWrite               MsgType = 2
)

// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1 }' | grep '^EVENT'
// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1, $4 }' | grep '^EVENT' | egrep -op '\d+';

type EventType int

const (
	EventAudio              EventType = 9
	EventBindSucess         EventType = 7
	EventCaePlainText       EventType = 1000
	EventCmdReturn          EventType = 8
	EventConnectedToServer  EventType = 13
	EventError              EventType = 2
	EventPreSleep           EventType = 10
	EventPushMessage        EventType = 1001
	EventResult             EventType = 1
	EventServerDisconnected EventType = 14
	EventSleep              EventType = 5
	EventStartRecord        EventType = 11
	EventState              EventType = 3
	EventStopRecord         EventType = 12
	EventVad                EventType = 6
	EventWakeup             EventType = 4
)

// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1 }' | grep '^STATE'
// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1, $4 }' | grep '^STATE' | egrep -op '\d+';
type StateType int

const (
	StateIdle    StateType = 1
	StateReady             = 2
	StateWorking           = 3
)

// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1 }' | grep '^VAD'
// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1, $4 }' | grep '^VAD' | egrep -op '\d+';

type VadType int

const (
	VadBos VadType = 0
	VadEos         = 2
	VadVol         = 1
)
