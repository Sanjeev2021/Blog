package Controller

import (
	"encoding/json"
	"net/http"

	"blog/model"
	"blog/service"
)

// create blog handler
func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	var newBlog model.Blog
	err := json.NewDecoder(r.Body).Decode(&newBlog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//sevice call for blog
	blog, err := service.NewBlog(newBlog.ID, newBlog.AuthorID, newBlog.Title, newBlog.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(blog)
}

// update blog title
func UpdateBlogTitle(w http.ResponseWriter, r *http.Request) {
	var updateBlog model.Blog
	err := json.NewDecoder(r.Body).Decode(&updateBlog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateBlogTitle(&updateBlog, updateBlog.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Handler  to delete Blog
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	var deleteBlog model.Blog
	err := json.NewDecoder(r.Body).Decode(&deleteBlog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteBlog(&deleteBlog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GET BLOGS HANDLER
func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var getBlog model.Blog
	err := json.NewDecoder(r.Body).Decode(&getBlog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	getblogs, err := service.GetBlogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(getblogs)
}
