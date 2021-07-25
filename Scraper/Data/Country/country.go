package Country

type Country struct {
	Name  string `json:"name"`
	Extra string `json:"extra"`
}

type Lists struct {
	Red   []Country `json:"red"`
	Amber []Country `json:"amber"`
	Green []Country `json:"green"`
}
