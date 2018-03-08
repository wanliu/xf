package gen

import (
	"strings"

	"github.com/pkg/errors"
)

func trans_args(args string) (res []Arg, _ error) {
	ss := strings.Split(args, ",")

	ex := &ErrTransArgs{Msg: "解析参数失败! 包含错误如下:"}

	for _, s := range ss {
		if arg, err := trans_c(s); err == ErrArgNone {
			continue
		} else if err != nil {
			ex.Append(errors.Wrap(err, "处理: "+s))
		} else {
			res = append(res, arg)
		}
	}

	return res, ex.EmptyOrError()
}

func insert_arg(ws []string, idx int, arg string) []string {
	nws := make([]string, 0, len(ws)+1)
	if idx > len(ws) {
		idx = len(ws)
	}
	nws = append(nws, ws[:idx]...)
	nws = append(nws, arg)
	nws = append(nws, ws[idx:]...)
	return nws
}

func trans_c(raw string) (Arg, error) {
	s := strings.TrimSpace(raw)
	ws := strings.Split(s, " ")
	if len(ws) == 0 {
		return Arg{}, ErrArgNone
	}

	for i, word := range ws {
		if len(word) == 0 {
			return Arg{}, ErrArgNone
		} else if word[0] == '*' {
			ws[i] = ws[i][len(ws[i])-1:]
			ws = insert_arg(ws, i, "*")
			break
			// ws = append(ws, )
		} else if word[len(word)-1] == '*' {
			ws[i] = ws[i][:len(ws[i])-1]
			ws = insert_arg(ws, i+1, "*")
			break
		} else if strings.Index(word, "*") >= 0 {
			pos := strings.Index(word, "*")
			ws = insert_arg(ws, i+1, ws[i][pos+1:])
			ws = insert_arg(ws, i+1, "*")
			ws[i] = ws[i][:pos]
			break
		}
	}

	for i := 0; i < len(ws); i += 1 {
		key := strings.Join(ws[:i+1], " ")
		// log.Printf("types: %s, %#v", key, ws)
		_, ok := Types[key]

		if ok {
			if typ, ok := Types[key].(string); ok {
				return Arg{
					ArgName: ws[len(ws)-1],
					Type:    typ,
					rawType: key,
				}, nil
			} else if typFunc, ok := Types[key].(ConvFunc); ok {
				return Arg{
					ArgName: ws[len(ws)-1],
					Type:    typFunc(key),
					rawType: key,
				}, nil
			}
		}
	}

	return Arg{}, ErrInvalidType(raw)
}

func trans_typ(raw string) (Type, error) {
	raw = strings.Replace(raw, "MSPAPI", "", 1)
	if arg, err := trans_c(raw); err != nil {
		return Type{}, ErrInvalidType(raw)
	} else {
		return Type{Type: arg.Type, rawType: arg.rawType}, nil
	}
}
