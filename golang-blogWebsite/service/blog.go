package service

import (
	"errors"
	"fmt"
	"time"

	"blog/database"
	"blog/model"
)

var allBlogs []*model.Blog

func NewBlog(id uint, userID uint, title string, content string) (*model.Blog, error) {

	user, err := GetUserById(userID)
	if err != nil {
		return nil, errors.New("cant find user")
	}

	if user == nil {
		return nil, errors.New("user does not exist")
	}

	if id <= 0 {
		return nil, errors.New("The id cant be less than 0 or equal to 0.")
	}

	// if authorid <= 0 {
	// 	return nil, errors.New("The author id cant be less than 0 or equal to 0.")
	// }

	if title == "" {
		return nil, errors.New("The title cant be empty.")
	}

	if content == "" {
		return nil, errors.New("the content cant be empty.")
	}

	createdAt := time.Now().Truncate(24 * time.Hour)
	updatedAt := time.Now().Truncate(24 * time.Hour)

	var tempblog = &model.Blog{
		ID:        id,
		UserID:    userID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Title:     title,
		Content:   content,
	}

	db := database.Get_GormDB()
	fmt.Println("Reached this line>>>>>")
	if err := db.Create(tempblog).Error; err != nil {
		return nil, err
	}
	fmt.Println("end of create db in NewBlog  service>>>>>>>>>>")

	allBlogs = append(allBlogs, tempblog)
	return tempblog, nil
}

func UpdateBlogTitle(b *model.Blog, title string) error {

	if b.Title == "" {
		return errors.New("the title cant be empty.")
	}

	if title == "" {
		return errors.New("The title cant be empty.")
	}

	b.Title = title
	return nil
}

func DeleteBlog(b *model.Blog) error {
	if b.Title == "" {
		return errors.New("The blog does not exist.")
	}

	b.ID = 0
	b.Title = ""

	return nil
}

func GetBlogs() ([]*model.Blog, error) {
	var blogs []*model.Blog

	db := database.Get_GormDB()
	if err := db.Find(&blogs).Error; err != nil {
		return nil, errors.New("No blog found.")
	}

	return blogs, nil
}

func GetBlogbyId(blogID uint) ([]*model.Blog, error) {
	var blog []*model.Blog

	db := database.Get_GormDB()
	if err := db.Find(&blog, blogID).Error; err != nil {
		return nil, errors.New("Blog not found.")
	}

	return blog, nil
}
