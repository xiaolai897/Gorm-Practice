package books

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Price  int
}

func CreateBooksTable(db *sqlx.DB) {
	table := `CREATE TABLE books (
		id INTEGER PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		author VARCHAR(100),
		price INTEGER
		)`
	db.Exec(table)
}

func CreateBookData(db *sqlx.DB) {
	sqlStr := `INSERT INTO books(title,author,price) 
		VALUES ('红楼梦','曹雪芹',60), ('活着','余华',28), ('三体','刘慈欣',168), ('三国演义','罗贯中', 40)`
	db.Exec(sqlStr)
}

func QueryBooks(db *sqlx.DB) {
	b := []Book{}
	err := db.Select(&b, "SELECT * FROM books WHERE price > ?", 50)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n大于50元的书籍:")
	for _, v := range b {
		fmt.Printf("\n\t名称: %s, 作者: %s, 价格: %d", v.Title, v.Title, v.Price)
	}
}
