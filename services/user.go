package services

import (
	"database/sql"
	"fmt"
	"learninggo/models"
)

type UserServices struct {
	connection *sql.DB
}

func NewUseService(connection *sql.DB) UserServices {
	return UserServices{
		connection: connection,
	}
}

func (pr *UserServices) CreateUser(user models.User) (int, error) {

	var id int
	query, err := pr.connection.Prepare("insert into users " +
		"(name, age)" +
		" values($1, $2) returning id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Name, user.Age).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *UserServices) GetUsers() ([]models.User, error) {
	query := "select * from users"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.User{}, err
	}

	var userList []models.User
	var userObj models.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.Id,
			&userObj.Name,
			&userObj.Age,
		)
		if err != nil {
			fmt.Println(err)
			return []models.User{}, err
		}

		userList = append(userList, userObj)

	}

	rows.Close()
	return userList, nil
}
