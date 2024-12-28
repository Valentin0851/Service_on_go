package user

type User struct {
	ID           string `json: "id" bson:"_is, omitempty"`
	Username     string `json: "username" bson:"username"`
	PasswordHash string `json: "-" bson:"password"`
	Email        string `json: "email" bson:"email"`
}

type CreateUserDTO struct {
	Username     string `json: "username"`
	PasswordHash string `json: "-"`
	Email        string `json: "email"`
}
