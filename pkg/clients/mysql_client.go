package clients

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/aaryanraj/ebook-store-api/pkg/config"
	"github.com/aaryanraj/ebook-store-api/pkg/models"
)

var db *gorm.DB

type Book models.Book
type User models.User

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.Debug().AutoMigrate(&Book{}, &User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", ID).Find(&book)
	return &book, db
}

func DeleteBook(ID int64) (Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", ID).Delete(book)
	return book, db
}

// 05/02/2023 code for user login

func (u *User) SaveUser() (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers() (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(uid uint32) (*User, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"full_name": u.FullName,
			"user_name": u.UserName,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
