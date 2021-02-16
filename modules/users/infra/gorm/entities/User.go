package entities

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	Id        string    `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Name      string    `json:"name" gorm:"notnull" valid:"notnull"`
	Email     string    `json:"email" gorm:"notnull;unique" valid:"email"`
	Password  string    `json:"-" gorm:"notnull" valid:"notnull"`
	Avatar    string    `json:"avatar" valid:"-"`
	CreatedAt time.Time `json:"create_at" gorm:"autoCreateTime" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" valid:"-"`
}

func (u *User) isValid() error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


type UserUnhide struct {
	Password  string    `json:"password" gorm:"notnull" validate:"required"`
	User
}

func NewUser(name string, email string, password string) (*User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		Id:       uuid.NewV4().String(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		CreatedAt: time.Now(),
	}

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
