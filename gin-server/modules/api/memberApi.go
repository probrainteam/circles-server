package api

import (
	. "circlesServer/modules/component"
	ErrChecker "circlesServer/modules/errors"
	. "circlesServer/modules/storage"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetMemberList(c *gin.Context) ([]Member, error) {
	circle, _ := c.Keys["circle"].(string)
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
	circle, _ := c.Keys["circle"].(string)
	_, err = db.Exec(`insert into `+circle+` (student_id, major, name, year, email, phone, paid, status) values (?,?,?,?,?,?,?,?)`, reqBody.SID, reqBody.MAJOR, reqBody.NAME, reqBody.YEAR, reqBody.EMAIL, reqBody.PHONE, reqBody.PAID, reqBody.STATUS)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}
func DeleteMember(c *gin.Context) error {
	sid := c.Params.ByName("sid")
	circle, _ := c.Keys["circle"].(string)
	db := DB()
	_, err := db.Exec(`delete ` + circle + ` where student_id = "` + sid + `"`)
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
	circle, _ := c.Keys["circle"].(string)
	db := DB()
	_, err = db.Exec(`delete from ` + circle + ` where student_id = "` + reqBody.SID + `"`)
	if err := ErrChecker.Check(err); err != nil {
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
	circle, _ := c.Keys["circle"].(string)
	db := DB()
	_, err = db.Exec(`update ` + circle + ` set status = 1 where student_id = "` + reqBody.SID + `"`)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	return nil
}
func Join(c *gin.Context) ([]Member, error) {
	circle, _ := c.Keys["circle"].(string)
	db := DB()
	rows, err := db.Query(`select * from ` + circle + `where status = 0`)
	if err := ErrChecker.Check(err); err != nil {
		return []Member{}, err
	}
	defer rows.Close()
	Joins := make([]Member, 0)
	var mem Member
	for rows.Next() {
		err := rows.Scan(&mem.SID, &mem.MAJOR, &mem.NAME, &mem.YEAR,
			&mem.EMAIL, &mem.PHONE, &mem.PAID, &mem.STATUS)
		if err := ErrChecker.Check(err); err != nil {
			return []Member{}, err
		}
		Joins = append(Joins, mem)
	}
	if len(Joins) == 0 {
		return []Member{}, errors.New("nothing to show")
	}
	return Joins, nil
}
func JoinApply(c *gin.Context) error {
	var reqBody JoinForm
	err := c.ShouldBind(&reqBody)
	if err != nil {
		return err
	}
	circle := GetCircle(uint64(reqBody.CIRCLE))
	db := DB()
	_, err = db.Exec(`insert into `+circle+` (student_id, major, name, year, email, phone, paid, status) values (?,?,?,?,?,?,?,?)`, reqBody.SID, reqBody.MAJOR, reqBody.NAME, reqBody.YEAR, reqBody.EMAIL, reqBody.PHONE, 0, 0)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}

	return nil
}
