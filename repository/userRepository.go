package repository

import (
	"app/models"
	"database/sql"
)

type UserRepository interface {
	CreateUser(user *models.User)
	GetUserById(userId int) models.User
}

type UserRepositoryImpl struct {
	db sql.DB
}

func NewUserReporitory(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{*db}
}

func (ur *UserRepositoryImpl) CreateUser(user *models.User) {
	_, err := ur.db.Exec("INSERT INTO t_user (first_name, last_name) VALUES (?, ?)", user.FIRST_NAME, user.LAST_NAME)
	println(err.Error())
}

func (ur *UserRepositoryImpl) GetUserById(id int) models.User {
	user := models.User{}
	ur.db.QueryRow("SELECT * FROM t_user where user_id = ?", id).Scan(&user.ID, &user.LAST_NAME, &user.FIRST_NAME)
	return user
}
