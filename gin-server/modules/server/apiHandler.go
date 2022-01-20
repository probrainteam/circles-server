package server

import (
	. "circlesServer/modules/api"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	token, err := LoginUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "token": token})
	}
}
func logout(c *gin.Context) {
	err := LogoutUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}
func modifyPubKey(c *gin.Context) {
	err := ModifyPubKey(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}
func modifyPW(c *gin.Context) {
	err := ModifyPW(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}
func resetPW(c *gin.Context) {
	err := FindUserPW(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}
func register(c *gin.Context) {
	err := RegisterUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}
func findID(c *gin.Context) {
	id, err := FindUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "id": id})
	}
}
func getMembers(c *gin.Context) {
	posts, err := GetMemberList(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "members": posts})
	}
}
func addMember(c *gin.Context) {
	pid, err := AddMember(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "pid": pid})
	}
}
func permitJoin(c *gin.Context) {
	err := Permit(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}
func denyJoin(c *gin.Context) {
	err := Deny(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil})
	}
}
func getJoin(c *gin.Context) {
	list, err := Join(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "join": list})
	}

}
func reissueAccess(c *gin.Context) {
	token, err := ReissueAccess(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"error": nil, "token": token})
	}

}
