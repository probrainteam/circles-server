package middleware

import (
	. "circlesServer/modules/component"
	"circlesServer/modules/token"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccessCheck(c *gin.Context) {
	// access token of request header check stage
	log.Println("Access Token Check Stage")
	accessValid, err := token.CheckTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if !accessValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "AT 만료.. RT 주세요"})
		c.Abort()
		return
	}
	num, err := GetCircleNum(c.Request, true)
	c.Set("circle", GetCircle(num))
	fmt.Println("c set circle : ", GetCircle(num), c.Params.ByName("circle"), c.Param("circle"))
	c.Next()
}
