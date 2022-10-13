package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) GetUserByEmail(db *gorm.DB, email string) (*User, error) {
    err := db.Debug().Model(&User{}).Where("email = ?", email).Take(&u).Error
    if err != nil {
        return &User{}, err
    }

    if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}

    return u, err
}

func (u *User) GetUserByEmailFake(db *gorm.DB, email string) (*User, error) {
    err := db.Debug().Raw("SELECT * FROM users WHERE email = '" + email + "'").Row().Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt)
    if err != nil {
        return &User{}, err
    }

    if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}

    return u, err
}
