package tools

//User JSON struct
type User struct {
	Name string `json:"name"`
	Interests []string `json:"interests"`
}

//Forum JSON struct
type Forum struct {
	Name string `json:"name"`
	Topic string `json: "topic"`
	Users []string `json:"users"`
}

//Forums JSON struct
type Forums struct {
	Forums []Forum `json:"forums"`
}

//ResponseName JSON struct
type ResponseName struct {
	Name string `json:"name"`
}