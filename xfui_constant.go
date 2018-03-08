package xf

// 获取 常量名称与数字
// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1 }' | grep '^CMD'
// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1, $4 }' | grep '^CMD' | egrep -op '\d+';
type MsgType int

const (
	CmdBuildGrammar        MsgType = 16
	CmdCaeWriteDevinfo             = 1000
	CmdCaeWriteEncrypt             = 1001
	CmdCleanDialogHistory          = 21
	CmdGetState                    = 1
	CmdQuerySyncStatus             = 24
	CmdReset                       = 4
	CmdResetWakeup                 = 8
	CmdResultValidationAck         = 20
	CmdSendLog                     = 12
	CmdSetBeam                     = 9
	CmdSetParams                   = 10
	CmdStart                       = 5
	CmdStartRecord                 = 22
	CmdStartSave                   = 14
	CmdStartThrowAudio             = 18
	CmdStop                        = 6
	CmdStopRecord                  = 23
	CmdStopSave                    = 15
	CmdStopThrowAudio              = 19
	CmdStopWrite                   = 3
	CmdSync                        = 13
	CmdUpdateLocalLexicon          = 17
	CmdUploadLexicon               = 11
	CmdWakeup                      = 7
	CmdWrite                       = 2
)

// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1 }' | grep '^EVENT'
// cat TAGS | awk -F $'\t' 'NR>5 && $7 = "class:aiui::AIUIConstant" { print $1, $4 }' | grep '^EVENT' | egrep -op '\d+';
type EventType int

const (
	EventAudio              EventType = 9
	EventBindSucess                   = 7
	EventCaePlainText                 = 1000
	EventCmdReturn                    = 8
	EventConnectedToServer            = 13
	EventError                        = 2
	EventPreSleep                     = 10
	EventPushMessage                  = 1001
	EventResult                       = 1
	EventServerDisconnected           = 14
	EventSleep                        = 5
	EventStartRecord                  = 11
	EventState                        = 3
	EventStopRecord                   = 12
	EventVad                          = 6
	EventWakeup                       = 4
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
