package server

import (
	"fmt"
	"log"

	_ "circlesServer/modules/storage"
	"circlesServer/modules/token"

	"github.com/gin-gonic/gin"
)

const port = ":4000"

func Serve(mode int) { // local : 4000 호스팅 시작
	r := gin.Default()
	if err := token.RedisInit(); err != nil {
		panic(fmt.Errorf("fatal error : redis is off status"))
	}
	api := r.Group("/api")
	if mode == 1 {
		api.Use(dummy)
	}
	RegisterApiHandlers(api)
	r.Run(port)
}

func dummy(c *gin.Context) {
	// access token of request header check stage
	log.Println("Access Token Check Stage")
}
func RegisterApiHandlers(api *gin.RouterGroup) {
	/*  Reply			200 -> token , uid
	400 -> ID or PW incorrect
	*/
	api.POST("/account/login", getLogin)

	/*  Reply			200 -> null
	400 -> Modify fail
	*/
	api.PUT("/account/pw", postModifyPW)

	/*  Reply			200 -> null ( mail send )
	400 -> DB Conn or Query err
	*/
	api.POST("/account/pw", postFindPW)

	/*  Reply			200 -> token delete
	400 -> ID or PW incorrect
	*/
	api.POST("/account/logout", postLogout)

	/*  Reply			200 -> register success
	400 -> DB Conn or Query err
	*/
	api.POST("/account", postRegister)

	/*  Reply			200 -> id
	400 -> DB Conn or Query err
	*/
	api.POST("/account/id", postFindID)

	/*  Reply			200 -> List<post>
	400 -> DB Conn or Query err
	*/
	api.GET("/members", getMembers)

	/*  Reply			200 -> null
	400 -> DB Conn or Query err
	*/
	api.POST("/members/add", postAddMember)

	/*  Reply			200 -> null
	400 -> DB Conn or Query err
	*/
	api.POST("/members/permit", postPermitJoin)

	/*  Reply			200 -> null
	400 -> DB Conn or Query err
	*/
	api.POST("/members/deny", postDenyJoin)

	/*  Reply			200 -> null
	400 -> DB Conn or Query err
	*/
	api.GET("/members/join", getJoin)

}
