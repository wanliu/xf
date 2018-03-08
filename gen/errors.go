package gen

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrArgNone = errors.New("Arg 是空的")

type ErrTransArgs struct {
	Msg      string
	children []error
}

func (ta *ErrTransArgs) Append(err error) {
	ta.children = append(ta.children, err)
}

func (ta *ErrTransArgs) Error() string {
	msg := fmt.Sprintf("参数错误: %s", ta.Msg)
	for _, err := range ta.children {
		msg += fmt.Sprintf("%+v\n", err)
	}
	return msg
}

func (ta *ErrTransArgs) EmptyOrError() error {
	if len(ta.children) == 0 {
		return nil
	} else {
		return ta
	}
}

type ErrorWithStack struct {
}

func ErrInvalidType(raw string) error {
	return fmt.Errorf("无效的转换类型字符: %s", raw)
}
