package users

type User struct{
	Id 			int64   `json:"id"`
	FirstName   string  `json:"firstName"`
	MiddleName  string  `json:"middleName"`
	LastName    string  `json:"lastName"`
	UserName    string  `json:"userName"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Type        string  `json:"type"`
	Images      string  `json:"images"`
}
