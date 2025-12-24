package handlers

import "github.com/gin-gonic/gin"


type UserHandler struct{}


func (u *UserHandler) GetUsers(c *gin.Context){   //get all users
	c.JSON(200 , gin.H{
			"users":"user list",
		})
}

func (u *UserHandler) GetUser(c *gin.Context){    // get single user
	c.JSON(200 , gin.H{
			"users":"user list",
		})
}

func (u *UserHandler) CreateUser(c *gin.Context){   // create user
	c.JSON(200 , gin.H{
			"users":"user list",
		})
}

func (u *UserHandler) DeleteUser(c *gin.Context){    // delete a user
	c.JSON(200 , gin.H{
			"users":"user list",
		})
}

func (u *UserHandler) UpdateUser(c *gin.Context){   // update user details
	c.JSON(200 , gin.H{
			"users":"user list",
		})
}

