// Code generated by "stringer -type=VadType"; DO NOT EDIT.

package xf

import "strconv"

const _VadType_name = "VadBos"

var _VadType_index = [...]uint8{0, 6}

func (i VadType) String() string {
	if i < 0 || i >= VadType(len(_VadType_index)-1) {
		return "VadType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VadType_name[_VadType_index[i]:_VadType_index[i+1]]
}