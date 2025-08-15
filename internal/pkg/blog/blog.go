package blog

import (
	"errors"
	"fmt"
	"gorm-practice/internal/pkg/idgen"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Base struct {
	ID        int64 `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Createor  int64
	Updater   int64
	DeletedAt soft_delete.DeletedAt
}

type Comment struct {
	Base
	PostID  int64
	Post    Post   `gorm:"foreignKey:PostID"`
	Content string `gorm:"not null"`
}

type Post struct {
	Base
	Title         string `gorm:"type:varchar(256);not null"`
	Content       string `gorm:"not null"`
	CommentNumber int64
	CommentState  string `gorm:"type:varchar(256)"`
	CommentID     int64
	Comments      []Comment `gorm:"foreignKey:PostID"`
	UserID        int64
	User          User `gorm:"foreignKey:UserID"`
}

type User struct {
	Base
	Name       string `gorm:"type:varchar(256);index:,unique,where:deleted_at = 0;not null"` // A regular string field
	Email      string `gorm:"type:varchar(256);index:,unique,where:deleted_at = 0;not null"`
	Username   string `gorm:"type:varchar(256);index:,unique,where:deleted_at = 0;not null"`
	Password   string `gorm:"type:varchar(256);index:,unique,where:deleted_at = 0;not null"`
	PostNumber int64
	Posts      []Post `gorm:"foreignKey:UserID"`
}

// 新建表
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Comment{}, &Post{}, &User{})
}

// 新建用户
func CreateUserData(db *gorm.DB) {
	u := User{
		Name:     "随便吧3",
		Email:    "113@qq.com",
		Username: "suibian3",
		Password: "mmqwew3",
	}
	db.Create(&u)
}

// 新建文章
func CreatePost(db *gorm.DB, id int64) {
	p := Post{
		Title:   "建个标题",
		Content: "建个标题文章内容",
		UserID:  id,
	}
	db.Create(&p)
}

// 新建评论
func CreateComment(db *gorm.DB, id int64) {
	c := Comment{
		PostID:  id,
		Content: "不错不错，第三篇",
	}
	db.Create(&c)
}

// 删除评论
func DeleteComment(db *gorm.DB, id int64) {
	var c Comment
	db.Where("id = ?", id).Find(&c)
	db.Delete(&c)
}

// 查询用户所有信息
func QueryUserAllInfo(db *gorm.DB, id int64) {
	var p []Post
	db.Preload("Comments").Where("user_id = ?", id).Find(&p)
	for _, v := range p {
		fmt.Printf("\n文章标题:%s", v.Title)
		fmt.Printf("\n评论数量:%d", v.CommentNumber)
		for _, comment := range v.Comments {
			fmt.Printf("\n\t评论内容:%s", comment.Content)
		}
	}
}

// 查询评论数量最多的文章
func QueryMaxComment(db *gorm.DB) {
	var p Post
	db.Preload("User").Order("comment_number desc").Limit(1).Find(&p)
	fmt.Printf("文章所属用户:%s\n", p.User.Name)
	fmt.Printf("文章标题:%s\n", p.Title)
	fmt.Printf("文章内容:%s\n", p.Content)
}

func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
	snowID := idgen.GenerateID()
	u.ID = snowID
	return
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if p.UserID == 0 {
		return errors.New("文章创建失败，没有关联的用户")
	}
	if err := tx.Model(&User{}).Select("id").Where("id = ?", p.UserID).Take(&User{}).Error; err != nil {
		return err
	}
	p.ID = idgen.GenerateID()
	p.Createor = p.UserID
	p.Updater = p.UserID
	return
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.PostID == 0 {
		return errors.New("评论失败，没有关联的文章")
	}
	if err := tx.Model(&Post{}).Select("id").Where("id = ?", c.PostID).Take(&Post{}).Error; err != nil {
		return err
	}
	c.ID = idgen.GenerateID()
	return
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_number", gorm.Expr("post_number + 1")).Error
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_number", gorm.Expr("comment_number + 1")).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var p Post
	if err := tx.Where("id = ?", c.PostID).Find(&p).Error; err != nil {
		return err
	}
	p.CommentNumber -= 1
	fmt.Println(p.CommentNumber)
	if p.CommentNumber == 0 {
		p.CommentState = "无评论"
		tx.Save(&p)
	}
	return
}
