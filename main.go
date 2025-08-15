package main

import (
	"fmt"
	"gorm-practice/internal/database"
	"gorm-practice/internal/pkg/account"
	"gorm-practice/internal/pkg/blog"
	"gorm-practice/internal/pkg/books"
	"gorm-practice/internal/pkg/employees"
	"gorm-practice/internal/pkg/student"
)

func main() {
	db := database.GetDB()
	sqldb := database.GetSqlxDB()
	s := student.QueryStudentData(db)
	fmt.Println(s)
	a := account.QueryAccountData(db)
	fmt.Println(a)
	t := account.QueryTransactions(db)
	fmt.Println(t)
	// blog.CreateTable(db)
	// blog.CreateUserData(db)
	// blog.CreatePost(db, 1955890128991817728)
	// blog.CreateComment(db, 1955891617487392768)
	// blog.DeleteComment(db, 1955893484200136704)
	// blog.QueryUserAllInfo(db, 1955889854453649408)
	blog.QueryMaxComment(db)
	employees.QueryEmployees(sqldb)
	employees.QueryMaxSalary(sqldb)
	books.QueryBooks(sqldb)
}
