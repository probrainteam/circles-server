package api

import (
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	ErrChecker "circlesServer/modules/errors"
	"circlesServer/modules/storage"
	"circlesServer/modules/token"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var gmail string
var gmailPW string

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	gmail = viper.GetString(`gmail.ID`)
	gmailPW = viper.GetString(`gmail.PW`)
}
func RegisterUser(c *gin.Context) error {
	var reqBody ResgisterForm
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	db := storage.DB()
	var count int
	_ = db.QueryRow(`your query or GORM`)

	if count > 0 {
		return errors.New("ID Duplicate")
	}
	_, err = db.Exec(`your query or GORM`)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}

func LoginUser(c *gin.Context) (uint64, map[string]string, error) {
	var reqBody LoginForm
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return 0, map[string]string{}, err
	}
	db := storage.DB()
	var pw string
	row := db.QueryRow(`your query or GORM`)
	var uid uint64
	err = row.Scan(&uid, &pw)
	if err := ErrChecker.Check(err); err != nil {
		return 0, map[string]string{}, errors.New("ID")
	}
	if reqBody.PW != pw { // PW 가 다르면 PW 가 다르다는 오류 반환
		return 0, map[string]string{}, errors.New("PW")
	}
	ts, err := token.CreateToken(uid)
	if err := ErrChecker.Check(err); err != nil {
		return 0, map[string]string{}, err
	}
	err = token.CreateAuth(uid, ts) // Redis 토큰 메타데이터 저장
	if err := ErrChecker.Check(err); err != nil {
		return 0, map[string]string{}, err
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	return uid, tokens, nil
}
func LogoutUser(c *gin.Context) error {
	// request header 에 담긴 access & refresh token을 검증 후 redis 에서 삭제
	au, ru, err := token.ExtractBothTokenMetadata(c.Request)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	deleted, err := token.DeleteAuth(au.AccessUuid, ru.RefreshUuid)
	if err := ErrChecker.Check(err); err != nil || deleted == 0 {
		return err
	}
	return nil
}
func FindUserPW(c *gin.Context) error {
	var reqBody struct {
		ID string `json:"email"`
	}
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	var email string
	var name string
	var count int
	db := storage.DB()
	row := db.QueryRow(`your query or GORM`)
	err = row.Scan(&count, &email, &name)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Invalid id")
	}
	pwByte := []byte{}
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		if a := rand.Intn(5); a < 4 {
			pwByte = append(pwByte, byte(rand.Intn(25)+97))
		} else {
			pwByte = append(pwByte, byte(rand.Intn(10)+48))
		}
	}
	pw := string(pwByte)

	_, err = db.Exec(`your query or GORM`)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", gmail, gmailPW, "smtp.gmail.com")
	from := gmail
	to := []string{reqBody.ID}
	headerSubject := "Subject: 같이할래 임시 PW 발급\r\n"
	headerBlank := "\r\n"

	body :=
		`안녕하세요 
	
프로브레인 개발팀입니다.

동아리 회원관리 시스템을 이용해주셔서 감사합니다.

` + name + `님의 임시 PW입니다.

PW:` + pw
	msg := []byte(headerSubject + headerBlank + body)
	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
	return nil
}
func FindUserId(c *gin.Context) (string, error) {
	var reqBody struct {
		PUBKEY string `json:"pubkey"`
	}
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return "", err
	}
	db := storage.DB()
	row := db.QueryRow(`your query or GORM`)
	var email string
	err = row.Scan(&email)
	if err := ErrChecker.Check(err); err != nil {
		return "", err
	}
	return email, nil
}
func ModifyPW(c *gin.Context) error {
	var reqBody ModifyForm
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	db := storage.DB()
	var count int
	uid := strconv.Itoa(reqBody.UID)
	_ = db.QueryRow(`your query or GORM` + uid).Scan(&count)
	if count == 0 {
		return errors.New("Invalid pw")
	}
	_, err = db.Exec(`your query or GORM`)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}
func ModifyProfile(c *gin.Context) error {

	return nil
}