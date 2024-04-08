package service

import "github.com/jinzhu/gorm"

var (
	DB              *gorm.DB
	CategoriesTable = "public.categories"
	ImagesTable     = "public.images"
	PostsTable      = "public.posts"
	UsersTable      = "public.users"
	RolesTable      = "public.roles"
)

func SetDB(db *gorm.DB) {
	DB = db
}
