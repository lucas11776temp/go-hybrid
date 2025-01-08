package user

type Course struct {
	Id        int32  `json:"id"`
	CreatedAt string `json:"created_at"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
