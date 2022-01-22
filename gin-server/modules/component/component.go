package component

import (
	. "circlesServer/const"
	. "circlesServer/modules/reader"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func GetCircle(num uint64) string {
	return []string{"probrain", "grow", "argos", "adm2n", "ana", "motion", "spg", "pai"}[num]
}

func GetAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func GetCircleNum(r *http.Request, access bool) (uint64, error) {
	encToken := strings.Split(r.Header.Get("Authorization"), " ")[1]
	claims := ReadToken(encToken, access)
	num, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return 10, err
	}
	if num > 7 {
		return 10, errors.New("cicle is not matched yet")
	}
	return num, nil
}

func GetUuid(r *http.Request, access bool) (string, error) {
	encToken := strings.Split(r.Header.Get("Authorization"), " ")[1]
	claims := ReadToken(encToken, access)
	var uuid string
	if access {
		uuid = claims[ACCESSUUID].(string)
	} else {
		uuid = claims[REFRESHUUID].(string)
	}
	return uuid, nil
}
func ReadToken(token string, access bool) jwt.MapClaims {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if access {
			return []byte(GetConfig(`token.ACCESS_SECRET`)), nil
		} else {
			return []byte(GetConfig(`token.REFRESH_SECRET`)), nil
		}
	})
	return jwt.MapClaims{}
}
