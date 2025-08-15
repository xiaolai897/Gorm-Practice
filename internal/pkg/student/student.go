package student

import "gorm.io/gorm"

type Student struct {
	ID    int64  `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(64)"`
	Age   int
	Grade string `gorm:"type:varchar(256)"`
}

func CreateStudentTable(db *gorm.DB) {
	db.AutoMigrate(&Student{})
}

func CreateStudentData(db *gorm.DB) {
	s := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	db.Create(&s)
}

func UpdateStudentData(db *gorm.DB) {
	var s Student
	db.Where("name = ?", "张三").Find(&s)
	s.Grade = "四年级"
	db.Save(&s)
}

func DeleteStudentData(db *gorm.DB) {
	var s Student
	db.Where("age > ?", 15).Delete(&s)
}

func QueryStudentData(db *gorm.DB) []Student {
	var s []Student
	db.Where("age > ?", 18).Find(&s)
	return s
}
