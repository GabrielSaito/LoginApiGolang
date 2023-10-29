package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) SearchId(Id uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nickname, email, dateCreated from users where id = ?",
		Id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()
	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.Id,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.DateCreated,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repository Users) SearchForEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	user := models.User{}

	if line.Next() {
		if err = line.Scan(&user.Id, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repository Users) Search(nameOrNick string) ([]models.User, error) {

	nameOrNick = fmt.Sprintf("%%s%%", nameOrNick)

	lines, err := repository.db.Query(

		"SELECT id, name, nickname, email, dateCreated FROM users WHERE name LIKE ? OR nickname LIKE ?;",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.Id,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.DateCreated,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository Users) Create(User models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nickname, email, password) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(User.Name, User.NickName, User.Email, User.Password)
	if err != nil {
		return 0, err
	}

	lastIdEntry, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIdEntry), nil
}

func (repository Users) Delete(Id uint64) error {
	statement, err := repository.db.Prepare("delete from users where Id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(Id); err != nil {
		return err
	}
	return nil
}

func (repository Users) SearchPassword(userId uint64) (string, error) {
	line, err := repository.db.Query("select password from users where id = ?", userId)
	return "", err

	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}
	return user.Password, nil
}
func (repository Users) NewPassword(userId uint64, password string) error {
	statement, err := repository.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userId); err != nil {
		return err
	}
	return nil
}
