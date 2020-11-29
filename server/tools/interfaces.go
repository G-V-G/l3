package tools

type Username struct {
	Username string `json:"username"`
}

type FullUser struct {
	Username  string   `json:"username"`
	Interests []string `json:"interests"`
}

type Interests struct {
	Interests []string `json:"interests"`
}

type ID struct {
	Id int `json:"id"`
}
