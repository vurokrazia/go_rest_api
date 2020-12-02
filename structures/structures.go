package structures
type User struct {
	Id int 				`json:"id"`
	Username string 	`json:"username"`
	First_name string `json:"first_name"`
	Last_name string 	`json:"last_name"`
}
type Response struct {
	Status string 	`json:"status"`
	Data User 		`json:"data"`
	Message string 	`json:"message"`
}