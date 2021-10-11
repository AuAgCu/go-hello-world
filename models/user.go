package models

type User struct {
	ID         int    `json:"id"`
	FIRST_NAME string `json:"firstName"`
	LAST_NAME  string `json:"lastName"`
}

type UserCollection struct {
	Users []User `json:"items"`
}

func GetUser() (tc UserCollection) {
	tc = UserCollection{
		[]User{
			{1, "Hello", "World"},
		},
	}

	return
}
