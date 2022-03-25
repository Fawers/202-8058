package routes

type RequestKeyVal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RequestKey struct {
	Key string `json:"key"`
}

type ResponseMsg struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg,omitempty"`
}

type ResponseGet struct {
	Ok    bool   `json:"ok"`
	Value string `json:"value,omitempty"`
}

type ResponseDel struct {
	Ok  bool   `json:"ok"`
	Key string `json:"key"`
}

type ResponseGetAll struct {
	Ok    bool   `json:"ok"`
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
