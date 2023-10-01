# task-5-pbi-btpns-DaveChristianThio

# Download important libraries
"gorm.io/driver/mysql"
"gorm.io/gorm"
"github.com/golang-jwt/jwt/v4"
"github.com/gin-gonic/gin"
"golang.org/x/crypto/bcrypt"


# This API uses MySQL
Create database [NAME]
Access database in phpmyadmin


# Document Structure
controllers: Contains database logic, namely models and queries
database: Contains database configuration and is used to run database connections and migrations
helpers : Contains functions that can be used in each place in this case jwt, bcrypt, headerValue
middlewares: Contains functions used for the JWT authentication process which is used for fire protection
models: Accommodates the creation of structs in this case using the user struct for data and authentication purposes, and contains models used for database relations
router: Contains the routing / endpoint configuration that will be used to access the api
go mod : Used for package / dependency management in the form of libraries
main.go : Initialize gin router and run gin server

# Register User (POST)
http://localhost:8080/users/register
Write in body (POSTMAN)
{
  "username": "[username]",
  "email": "[email]",
  "password": "[password]"
}

# Login User (POST) (Adds JWT Token (Cookies))
http://localhost:8080/users/login
Write in body (POSTMAN)
{
  "username": "[username]",
  "password": "[password]"
}

# Logout (GET) (Removes existing JWT Token (Cookies))
http://localhost:8080/users/login

# Update User (PUT)
http://localhost:8080/users/:id
Write in body (POSTMAN)
{
  "username": "[username]",
  "email": "[email]",
  "password": "[password]"
}

# Delete User (DELETE)
http://localhost:8080/users/:id


# Get Photo (GET) (Gets every photo uploaded by unique user ID)
http://localhost:8080/photos

# Create Photo (POST) (Adds new photo to the user ID)
http://localhost:8080/photos

# Update Photo (PUT)
http://localhost:8080/photos/:photoId
Write in body (POSTMAN)
{
  "title": "[title]",
  "caption": "[caption]",
  "photo_url": "[photoURL]"
}

# Delete Photo (DELETE)
http://localhost:8080/photos/:photoId
