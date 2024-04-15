package routes

import (
	"net/http"

	controller "blog/Controller"
)

func Routes() {
	http.HandleFunc("/create-blog", controller.CreateBlogHandler)
	http.HandleFunc("/update-blog", controller.UpdateBlogTitle)
	http.HandleFunc("/delete-blog", controller.DeleteBlog)
	http.HandleFunc("/get-blog", controller.GetBlogs)

	http.HandleFunc("/create-user", controller.CreateUserHandler)
	http.HandleFunc("/update-username", controller.UpdateUsernameHandler)
	http.HandleFunc("/delete-user", controller.DeleteUser)
	http.HandleFunc("/get-users", controller.GetUser)

}
