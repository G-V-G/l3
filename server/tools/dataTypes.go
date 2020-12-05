package tools

//User JSON struct
type User struct {
	Id        int      `json:"-"`
	Name string `json:"name"`
	Interests []string `json:"interests"`
}

//Forum JSON struct
type Forum struct {
	Id      int    `json:"-"`
	Name string `json:"name"`
	Topic string `json: "topic"`
	Users []string `json:"users"`
}

//Forums JSON struct
type Forums struct {
	Forums []*Forum `json:"forums"`
}

//ResponseName JSON struct
type ResponseName struct {
	Name string `json:"name"`
}

// Users array
type Users struct {
	UsersArr []*User `json:"users"`
}

type errorObject struct {
	Message string `json:"message"`
}
