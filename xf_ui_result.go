package xf

type AIUIResult struct {
	Data []AIUIResultData `json: "data"`
}

type AIUIResultData struct {
	Params struct {
		Sub string `json: "sub"`
	} `json: "params"`
	Content []AIUIResultContent `json: "content"`
}

type AIUIResultContent struct {
	DTE    string `json: "dte"`
	DTF    string `json: "dtf"`
	CNT_ID string `json: "cnt_id"`
}
