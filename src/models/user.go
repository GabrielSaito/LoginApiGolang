package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	Id          uint64    `json:"id, omitempty"`
	Name        string    `json:"name, omitempty"`
	NickName    string    `json:"nickname, omitempty"`
	Email       string    `json:"email, omitempty"`
	Password    string    `json:"password, omitempty"`
	DateCreated time.Time `json:"password, omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validator(); err != nil {
		return err
	}
	if err := user.refactor(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validator() error {
	if user.Name == " " {
		return errors.New("")
	}
	if user.NickName == " " {
		return errors.New("")
	}
	if user.Email == " " {
		return errors.New("")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email invalid")
	}

	if user.Password == " " {
		return errors.New("")
	}
	return nil
}

func (user *User) refactor(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordWithHash, err := security.Hash(user.Password)

		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}
	return nil
}
