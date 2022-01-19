package api

import (
	. "circlesServer/modules/reader"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func getCircle(num uint64) string {
	return []string{"probrain", "grow", "argos", "adm2n", "ana", "motion", "spg", "pai"}[num]
}

func getAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func getCircleNum(c *gin.Context) (uint64, error) {
	access := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
	fmt.Println(access)
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(access, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetConfig(`token.ACCESS_SECRET`)), nil
	})

	fmt.Println(token)
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	num, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return 10, err
	}
	return num, nil
}
