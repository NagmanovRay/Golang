package models

import(
	"github.com/jinzhu/gorm"
	"github.com/akhil/go-bookstore/pkg/config"
)

var db *form.DB
type Book struct {
	gorm.Book
	Name string `gorm:"json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(%Book{})
}

func (b *Book) CreateBook(){
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Book)
	return Book
}

func GetBookByID(Id init64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID init64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}