package sooc

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	CompanyId int    `json:"comp_id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}
