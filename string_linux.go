package xf

import "C"

type NString struct {
	Value  string
	isHave bool
}

func (ns *NString) CString() *C.char {
	if ns.isHave {
		return C.CString(ns.Value)
	} else {
		return nil
	}
}

func NullString(val string) *NString {
	return &NString{Value: val, isHave: true}
}

func Empty2Null(val string) *C.char {
	if len(val) == 0 {
		return nil
	}

	return C.CString(val)
}
