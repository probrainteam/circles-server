package server

import (
	. "circlesServer/modules/reader"
	. "circlesServer/modules/storage"
	"circlesServer/modules/token"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var port string = `:` + GetConfig(`server.PORT`)

func Serve(mode string) { // local : 4000 호스팅 시작
	r := gin.Default()
	publicAPI := r.Group("/api") // no need auth
	authAPI := r.Group("/api")   // need auth
	if mode == `deploy` {
		authAPI.Use(dummy)
	} else if mode == `dev` { // use mock data
	} else if mode == `debug` { // log everything
		r.Use(logAll)
	} else {
		panic(fmt.Errorf(`Unknown command : ` + mode)) // exception
	}
	_, err := Redis()
	if err != nil {
		panic(fmt.Errorf("redis is off status"))
	}
	if DB().Ping() != nil {
		panic(fmt.Errorf("mysql is off status"))
	}
	RegisterApiHandlers(publicAPI, authAPI)
	r.Run(port)
}
func logAll(c *gin.Context) {
	// request, time, user log
	log.Println("LOG")

}
func dummy(c *gin.Context) {
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
}
func RegisterApiHandlers(api *gin.RouterGroup, auth *gin.RouterGroup) {
	/*  Reply			200 -> token , uid
	400 -> ID or PW incorrect
	*/
	api.POST("/user/login", login)

	/*  Reply			200 -> token delete
	400 -> ID or PW incorrect
	*/
	auth.POST("/user/logout", logout)

	/*  Reply			200 -> null
	400 -> Modify fail
	*/
	auth.PUT("/user/pubkey", modifyPubKey)

	/*  Reply			200 -> null
	400 -> Modify fail
	*/
	auth.PUT("/user/pw", modifyPW)

	/*  Reply			200 -> null ( mail send )
	400 -> DB Conn or Query err
	*/
	api.POST("/user/pw", resetPW)

	/*  Reply			200 -> register success
	400 -> DB Conn or Query err
	*/
	api.POST("/user", register)

	/*  Reply			200 -> id
	400 -> DB Conn or Query err
	*/
	api.POST("/user/id", findID)

	/*  Reply			200 -> List<member>
	400 -> DB Conn or Query err
	*/
	auth.GET("/members", getMembers)

	/*  Reply			200 -> null
	400 -> DB Conn or Query err
	*/
	auth.POST("/member", addMember)

	/*  Reply			200 -> null
	400 -> DB Conn or Query err
	*/
	auth.POST("/members/permit", permitJoin)

	/*  Reply			200 -> null
	400 -> DB Conn or Query err
	*/
	auth.POST("/members/deny", denyJoin)

	/*  Reply			200 -> List<member>
	400 -> DB Conn or Query err
	*/
	auth.GET("/members/join", getJoin)

	/*  Reply			200 -> access Token
	400 -> refresh
	*/
	api.POST("/token", reissueAccess)

}
