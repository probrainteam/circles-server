package api

import (
	"circlesServer/modules/component"
	ErrChecker "circlesServer/modules/errors"
	. "circlesServer/modules/reader"
	"circlesServer/modules/storage"
	"circlesServer/modules/token"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
)

var gmail string
var gmailPW string

func init() {
	gmail = GetConfig(`gmail.ID`)
	gmailPW = GetConfig(`gmail.PW`)
}
func TestGetAllTable(c *gin.Context, table string) error {
	db := storage.DB()
	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
	var first string
	var second string
	var t string
	var q string
	for rows.Next() {
		err := rows.Scan(&first, &second, &t, &q)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(first, second, t, q)
	}
	return nil
}
func RegisterUser(c *gin.Context) error {
	var reqBody ResgisterForm
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	db := storage.DB()
	var count int
	_ = db.QueryRow(`select count(*) from manager where email = "` + reqBody.Email + `"`).Scan(&count)

	if count > 0 {
		return errors.New("ID Duplicate")
	}
	_, err = db.Exec(`insert into manager (email, pw, pubkey) values (?,?,?)`, reqBody.Email, reqBody.PW, reqBody.PUBKEY)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}

func LoginUser(c *gin.Context) (string, string, error) {
	var reqBody LoginForm
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return "", "", err
	}
	db := storage.DB()
	var pw string
	var count int
	var circle uint64
	row := db.QueryRow(`select count(*), pw, circle from manager where email = '` + reqBody.ID + `'`)
	err = row.Scan(&count, &pw, &circle)
	if err := ErrChecker.Check(err); err != nil {
		return "", "", errors.New("ID")
	}
	if reqBody.PW != pw { // PW 가 다르면 PW 가 다르다는 오류 반환
		return "", "", errors.New("PW")
	}
	ts, err := token.CreateToken(circle)
	if err := ErrChecker.Check(err); err != nil {
		return "", "", err
	}
	err = token.CreateAuth(circle, ts) // Redis 토큰 메타데이터 저장
	if err := ErrChecker.Check(err); err != nil {
		return "", "", err
	}
	c.SetCookie("refreshToken", ts.RefreshToken, 60*60*24*7, "/", "", true, true)
	return ts.AccessToken, component.GetCircle(circle), nil
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
		EMAIL string `json:"email"`
	}
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	var email string
	var count int
	db := storage.DB()
	err = db.QueryRow(`select count(*),email from manager where email = "`+reqBody.EMAIL+`"`).Scan(&count, &email)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	if count == 0 {
		return errors.New("invalid email")
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
	_, err = db.Exec(`update manager set pw ="` + pw + `" where email = "` + reqBody.EMAIL + `"`)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	auth := smtp.PlainAuth("", gmail, gmailPW, "smtp.gmail.com")
	from := gmail
	to := []string{reqBody.EMAIL}
	headerSubject := "Subject: 동아리 관리자 임시 PW 발급\r\n"
	headerBlank := "\r\n"
	body :=
		`안녕하세요 
	
프로브레인 개발팀입니다.

동아리 회원관리 시스템을 이용해주셔서 감사합니다.

` + email + `님의 임시 PW입니다.

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
	circle, _ := c.Keys["circle"].(string)
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	db := storage.DB()
	var count int
	_ = db.QueryRow(`select count(*) from manager where circle =` + circle + `and pw = "` + reqBody.PW + `"`).Scan(&count)
	if count == 0 {
		return errors.New("Invalid pw")
	}
	_, err = db.Exec(`update manager set pw = "` + reqBody.NEW + `" where circle = ` + circle)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}
func ModifyPubKey(c *gin.Context) error {
	return nil
}
func ReissueAccess(c *gin.Context) (string, error) {
	token, err := token.ReissueAccessToken(c.Request)
	if err != nil {
		return "", err
	}
	return token, nil
}
