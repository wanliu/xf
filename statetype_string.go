// Code generated by "stringer -type=StateType"; DO NOT EDIT.

package xf

import "strconv"

const _StateType_name = "StateIdle"

var _StateType_index = [...]uint8{0, 9}

func (i StateType) String() string {
	i -= 1
	if i < 0 || i >= StateType(len(_StateType_index)-1) {
		return "StateType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _StateType_name[_StateType_index[i]:_StateType_index[i+1]]
}
