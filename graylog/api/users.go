package api

type UserApi struct {
	Id string
	Username string
}

type UsersApi struct {
	Users *[]UserApi
}
