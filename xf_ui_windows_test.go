package xf

import (
	"testing"
)

func TestWinXF_UI(t *testing.T) {
	msg := NewMessage(CmdStart)
	t.Logf("msg: %#v", msg)
}
