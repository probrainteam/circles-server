package component

import (
	. "circlesServer/modules/reader"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"

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

func GetCircleNumAcc(access string) (uint64, error) {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(access, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetConfig(`token.ACCESS_SECRET`)), nil
	})
	num, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return 10, err
	}
	if num > 7 {
		return 10, errors.New("cicle is not matched yet")
	}
	return num, nil
}
func GetCircleNumRef(refresh string) (uint64, error) {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(refresh, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetConfig(`token.REFRESH_SECRET`)), nil
	})
	num, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return 10, err
	}
	if num > 7 {
		return 10, errors.New("cicle is not matched yet")
	}
	return num, nil
}
