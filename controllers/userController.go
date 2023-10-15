package controllers

import (
	"Demo/database"
	token "Demo/middlewares"
	"Demo/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	if checkEmail.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exists",
		})
		return
	}

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

		token, err := token.GenerateToken(user.Id)

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
