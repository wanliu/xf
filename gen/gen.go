package gen

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
	"strings"
)

type Tag struct {
	Name   string
	Args   []Arg
	Ret    Type
	source string
}

type Arg struct {
	ArgName string
	Type    string
	rawType string
}

type ValType struct {
	ValName string
	Value   string
}

type Type struct {
	Type    string
	rawType string
}

const SkipHeader = 6

var Types = map[string]interface{}{
	"int":                      "int64",
	"const char":               "string",
	"void *":                   "[]byte",
	"char *":                   "string",
	"char":                     "string",
	"int *":                    "*int",
	"unsigned int":             "uint64",
	"unsigned int *":           "*uint64",
	"unsigned int64":           "uint64",
	"unsigned int64 *":         "*uint64",
	"const wchar_t *":          "string",
	"wchar_t *":                "string",
	"const char*":              "string",
	"const char *":             "string",
	"const void *":             "string",
	"DownloadStatusCB":         "DownloadStatusCB",
	"DownloadResultCB":         "DownloadResultCB",
	"GrammarCallBack":          "GrammarCallBack",
	"LexiconCallBack":          "LexiconCallBack",
	"NLPSearchCB":              "NLPSearchCB",
	"recog_result_ntf_handler": "recog_result_ntf_handler",
	"recog_status_ntf_handler": "recog_status_ntf_handler",
	"recog_error_ntf_handler":  "recog_error_ntf_handler",
	"tts_result_ntf_handler":   "tts_result_ntf_handler",
	"tts_error_ntf_handler":    "tts_error_ntf_handler",
	"tts_status_ntf_handler":   "tts_status_ntf_handler",
	"msp_status_ntf_handler":   "msp_status_ntf_handler",
}

var TypeConverts = map[string]interface{}{
	"void *":          "C.CBytes(%v)",
	"char *":          "C.CString(%v)",
	"const char *":    "C.CString(%v)",
	"const wchar_t *": "&C.StringToUTF16(%v)[0]",
	"wchar_t *":       "&C.StringToUTF16(%v)[0]",
}

var ReFunc = regexp.MustCompile("(\\w+\\*?)?\\s+(\\*?\\w+)\\((.*?)\\)")
var globFuncs = template.FuncMap{
	"join": strings.Join,
	"raw":  raw,
}

var tagTmpl = template.Must(template.New("tag").Funcs(globFuncs).Parse(`
	{{define "list"}}{{join . ", " | raw}}{{end}}
	func {{.Name}}({{template "list" .ArgsDefine }}) {{.Ret.Define}} {
		return C.{{.Name}}({{template "list" .ArgsCall}}})
	}
`))

func raw(x string) interface{} { return template.HTML(x) }

func ParseTags(filename string) ([]Tag, error) {
	var (
		// buf []byte
		err error
	)
	file, err := os.Open(filename) // For read access.
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(file)
	for i := 0; i < SkipHeader; i += 1 {
		r.ReadLine()
	}

	var tags = make([]Tag, 0)
	for {
		raw, _, err := r.ReadLine()
		if err != nil {
			break
		}

		if tag, err := parseTag(string(raw)); err == nil {
			tags = append(tags, tag)
		} else {
			log.Printf("parseTag 失败: %s", err)
		}
	}
	return tags, nil
}

func parseTag(line string) (Tag, error) {
	line = strings.Replace(line, "MSPAPI", "", -1)
	sects := strings.Split(line, "\t")

	fun, re := sects[0], sects[2]
	mats := ReFunc.FindStringSubmatch(re)
	if len(mats) < 1 {
		return Tag{}, fmt.Errorf("不能匹配函数的格式: %s", re)
	}
	source := mats[0]
	args_s := mats[3]
	fun = mats[2]
	ret_s := mats[1]

	args, err := trans_args(args_s)
	if err != nil {
		return Tag{}, err
	}

	// log.Printf("ret %s", ret_s)

	ret, err := trans_typ(ret_s)
	if err != nil {
		return Tag{}, err
	}

	return Tag{
		Name:   fun,
		Args:   args,
		Ret:    ret,
		source: source,
	}, nil
}

func (t *Tag) Source() string {
	buf := &bytes.Buffer{}
	tagTmpl.Execute(buf, t)
	return buf.String()
}

func (t *Tag) ArgsDefine() []string {
	var args []string
	for _, arg := range t.Args {
		args = append(args, arg.Define())
	}

	return args
}

func (t *Tag) ArgsCall() []string {
	var args []string
	for _, arg := range t.Args {
		args = append(args, arg.Call())
	}

	return args
}

func (a *Arg) Define() string {
	return fmt.Sprintf("%s %s", a.ArgName, a.Type)
}

func (a *Arg) Call() string {
	if _, ok := TypeConverts[a.rawType]; ok {
		log.Printf("rawType: %s", a.rawType)
		if typ, ok := TypeConverts[a.rawType].(string); ok {
			return fmt.Sprintf(typ, a.ArgName)
		}
	}
	return fmt.Sprintf("%s", a.ArgName)
}

func (t *Type) Define() string {
	return t.Type
}
