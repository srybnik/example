package models

type Response struct {
	Data   any    `json:"data,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Error  string `json:"error,omitempty"`
	Result int    `json:"result"`
	Reason uint32 `json:"reason,omitempty"`
	State  int8   `json:"state"`
}
