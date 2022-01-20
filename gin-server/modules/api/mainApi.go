package api

import (
	. "circlesServer/modules/component"
	ErrChecker "circlesServer/modules/errors"
	. "circlesServer/modules/storage"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMemberList(c *gin.Context) ([]Member, error) {
	num, err := GetCircleNumAcc(strings.Split(c.Request.Header.Get("Authorization"), " ")[1])
	if err != nil {
		return []Member{}, err
	}
	circle := GetCircle(num)
	db := DB()
	rows, err := db.Query(`select * from ` + circle)
	if err := ErrChecker.Check(err); err != nil {
		return []Member{}, err
	}
	defer rows.Close()
	Members := make([]Member, 0)
	var mem Member
	for rows.Next() {
		err := rows.Scan(&mem.SID, &mem.MAJOR, &mem.NAME, &mem.YEAR,
			&mem.EMAIL, &mem.PHONE, &mem.PAID, &mem.STATUS)
		if err := ErrChecker.Check(err); err != nil {
			return []Member{}, err
		}
		Members = append(Members, mem)
	}
	if len(Members) == 0 {
		return []Member{}, errors.New("nothing to show")
	}
	return Members, nil
}
func AddMember(c *gin.Context) error {
	var reqBody Member
	err := c.ShouldBind(&reqBody)
	if err != nil {
		return err
	}
	db := DB()
	num, err := GetCircleNumAcc(strings.Split(c.Request.Header.Get("Authorization"), " ")[1])
	circle := GetCircle(num)
	_, err = db.Exec(`insert into `+circle+` (student_id, major, name, year, email, phone, paid, status) values (?,?,?,?,?,?,?,?)`, reqBody.SID, reqBody.MAJOR, reqBody.NAME, reqBody.YEAR, reqBody.EMAIL, reqBody.PHONE, reqBody.PAID, reqBody.STATUS)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}
func Deny(c *gin.Context) error {
	var reqBody ReplyJoinForm
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	db := DB()
	var count int
	_ = db.QueryRow(`your query or GORM`).Scan(&count)
	if count == 0 {
		return errors.New("Nothing")
	}
	_, err = db.Exec(`your query or GORM`)

	if err != nil {
		return err
	}
	return nil
}
func Permit(c *gin.Context) error {
	var reqBody ReplyJoinForm
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	db := DB()
	var count int
	_ = db.QueryRow(`your query or GORM`)
	if count == 0 {
		return errors.New("Nothing")
	}
	_, err = db.Exec(`your query or GORM`)

	if err != nil {
		return err
	}
	return nil
}
func Join(c *gin.Context) ([]Member, error) {
	var reqBody Member
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return []Member{}, err
	}
	db := DB()
	_, err = db.Exec(`your query or GORM`)
	if err != nil {
		return []Member{}, err
	}
	list := make([]Member, 0)
	return list, nil
}
func GetNumMember(c *gin.Context) (int, error) {
	var reqBody Member
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return -1, err
	}
	db := DB()
	_, err = db.Exec(`your query or GORM`)
	if err != nil {
		return -1, err
	}
	return -1, nil
}
