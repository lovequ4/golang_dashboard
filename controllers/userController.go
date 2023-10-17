package controllers

import (
	"Demo/database"
	token "Demo/middlewares"
	"Demo/models"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//為了讓swagger 可以@注釋使用，把struct移到func外
type SignUpForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// @Tags   UserAPI
// @Router /users/signup [post]
// @Param  signUpFormData body SignUpForm true "User registration data"
func SignUp(c *gin.Context) {
	user := models.User{}

	var signUpFormData SignUpForm

	//bind 綁定數據
	err := c.Bind(&signUpFormData)

	// c.json 回傳 http status code
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found SignUpForm Data",
		})
		return
	}

	checkEmail := database.DB.Where("email=?", signUpFormData.Email).First(&user)
	if checkEmail.Error != nil {
		if errors.Is(checkEmail.Error, gorm.ErrRecordNotFound) {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpFormData.Password), bcrypt.DefaultCost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error occurred encrypting password",
				})
				return
			}

			user = models.User{
				Name:     signUpFormData.Name,
				Password: string(hashedPassword),
				Email:    signUpFormData.Email,
				Date:     time.Now(),
			}

			result := database.DB.Save(&user)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Database operation failed",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "Saved data successfully",
				})
			}

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "An error occurred while checking email",
			})
			return
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exists",
		})
		return
	}
}

type SignInForm struct {
	NameOrEmail string `json:"nameoremail"`
	Password    string `json:"password"`
}

// @Tags   UserAPI
// @Router /users/signin [post]
// @Param SignInformData body SignInForm true "User SignUp data"
func SignIn(c *gin.Context) {
	user := models.User{}

	var signInFormData SignInForm

	err := c.BindJSON(&signInFormData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found SignInForm Data",
		})
		return
	}

	result := database.DB.Where("name=? OR email=?", signInFormData.NameOrEmail, signInFormData.NameOrEmail).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return

	} else {

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInFormData.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Incorrect password",
			})
		}

		token, err := token.GenerateToken(user.Id, user.Name, user.Role)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to generate token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

// @Tags   UserAPI
// @Router /users [get]
// @Param Authorization header string true "JWT token" default(Bearer)
func AllUsers(c *gin.Context) {
	var users []models.User
	database.DB = database.DB.Debug()
	if err := database.DB.Select("id,name,email,role,date").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to find user",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

type UserPutForm struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// @Tags   UserAPI
// @Router /users/{id} [put]
// @Param  id path int true "User ID"
// @Param Authorization header string true "JWT token" default(Bearer)
// @Param Request body UserPutForm true "User Put data"
func UserPut(c *gin.Context) {
	userId := c.Param("id")

	user := models.User{}

	var PutForm UserPutForm

	if err := c.ShouldBindJSON(&PutForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not Found PutForm data",
		})
		return
	}

	if err := database.DB.Where("id=?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Database operation failed",
			})
		}
		return
	}

	updates := map[string]interface{}{
		"Name":  PutForm.Name,
		"Email": PutForm.Email,
		"Role":  PutForm.Role,
	}

	result := database.DB.Model(&user).Updates(updates)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database operation failed",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
