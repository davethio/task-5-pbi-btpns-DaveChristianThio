# task-5-pbi-btpns-DaveChristianThio

# Download important libraries
# "gorm.io/driver/mysql"
#	"gorm.io/gorm"
# "github.com/golang-jwt/jwt/v4"
# "github.com/gin-gonic/gin"
# "golang.org/x/crypto/bcrypt"


# This API uses MySQL
# Create database [NAME]
# Access database in phpmyadmin


# Document Structure
# controllers : Berisi antara logic database yaitu models dan query
# database: Berisi konfigurasi database serta digunakan untuk menjalankan koneksi database dan migration
# helpers : Berisi fungsi-fungsi yang dapat digunakan di setiap tempat dalam hal ini jwt, bcrypt, headerValue
# middlewares :Berisi fungsi yang digunakan untuk proses otentikasi jwt yang digunakan untuk proteksi api
# models : Menampung pembuatan struct dalam kasus ini menggunakan struct user untuk keperluan data dan authentication, dan berisi models yang digunakan untuk relasi database 
# router : Berisi konfigurasi routing / endpoint yang akan digunakan untuk mengakses api
# go mod : Yang digunakan untuk manajemen package / dependency berupa library

# Register User (POST)
# http://localhost:8080/users/register
# Write in body (POSTMAN)
# {
#  "username": "[username]",
#  "email": "[email]",
#  "password": "[password]"
# }

# Login User (POST) (Adds JWT Token (Cookies))
# http://localhost:8080/users/login
# Write in body (POSTMAN)
# {
#  "username": "[username]",
#  "password": "[password]"
# }

# Logout (GET) (Removes existing JWT Token (Cookies))
# http://localhost:8080/users/login

# Update User (PUT)
# http://localhost:8080/users/:id
# Write in body (POSTMAN)
# {
#  "username": "[username]",
#  "email": "[email]",
#  "password": "[password]"
# }

# Delete User (DELETE)
# http://localhost:8080/users/:id


# Get Photo (GET) (Gets every photo uploaded by unique user ID)
# http://localhost:8080/photos

# Create Photo (POST) (Adds new photo to the user ID)
# http://localhost:8080/photos

# Update Photo (PUT)
# http://localhost:8080/photos/:photoId
# Write in body (POSTMAN)
# {
#  "title": "[title]",
#  "caption": "[caption]",
#  "photo_url": "[photoURL]"
# }

# Delete Photo (DELETE)
# http://localhost:8080/photos/:photoId
