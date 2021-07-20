package Country


type Country struct {
	Name  string `json:"name"`
	Extra string `json:"extra"`
}

type Countries []Country

type Lists struct {
	Red   Countries `json:"red"`
	Amber Countries `json:"amber"`
	Green Countries `json:"green"`
}