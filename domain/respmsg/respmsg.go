package respmsg

type Resp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
