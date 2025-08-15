package employees

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Employees struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

func CreateEmployeesTable(db *sqlx.DB) {
	table := `CREATE TABLE employees (
		id INTEGER PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		department VARCHAR(100),
		salary INTEGER
		)`
	db.Exec(table)
}

func CreateEmployeesData(db *sqlx.DB) {
	sqlStr := `INSERT INTO employees(name,department,salary) VALUES ('王五','技术部',10000)`
	db.Exec(sqlStr)
}

func QueryEmployees(db *sqlx.DB) {
	e := []Employees{}
	err := db.Select(&e, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		panic(err)
	}
	for _, v := range e {
		fmt.Printf("名字: %s, 部门: %s, 工资: %d\n", v.Name, v.Department, v.Salary)
	}
}

func QueryMaxSalary(db *sqlx.DB) {
	var e Employees
	err := db.Get(&e, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println("工资最高的:")
	fmt.Printf("\n\t名字: %s, 部门: %s, 工资: %d\n", e.Name, e.Department, e.Salary)
}
