package users

import (
	"github.com/Emanuel9/bookstore_users-api/domain/users"
	"github.com/Emanuel9/bookstore_users-api/services"
	"github.com/Emanuel9/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	return
	//}
	//
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	//c.String(http.StatusNotImplemented, "implement me!")
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implement me!")
//}
