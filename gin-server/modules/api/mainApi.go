package api

import (
	ErrChecker "circlesServer/modules/errors"
	. "circlesServer/modules/storage"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetMemberList(c *gin.Context) ([]Member, error) {
	num, err := getCircleNum(c)
	if err != nil {
		return []Member{}, err
	}
	circle := getCircle(num)
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
func AddMember(c *gin.Context) (int, error) {
	var reqBody Member

	err := c.ShouldBind(&reqBody)
	if err != nil {
		return 1, err
	}

	// file, _, err := c.Request.FormFile("content")

	// var pid int
	db := DB()
	// basePath := "http://" + getAddr()
	// localPath := "/Users/macbook/Sites" // custom
	// dirPath := "/Member_img/"           // custom
	// path := basePath + dirPath + strconv.Itoa(pid+1) + `.png`

	_, err = db.Exec(`insert into `)
	if err := ErrChecker.Check(err); err != nil {
		return -1, err
	}

	// path = localPath + dirPath + strconv.Itoa(pid+1) + `.png`
	// dst, err := os.Create(path)
	// if err != nil {
	// 	return -1, err
	// }
	// defer dst.Close()

	if err != nil {
		return -1, err
	}
	// if _, err := io.Copy(dst, file); err != nil {
	// 	return -1, err
	// }
	_, err = db.Exec(`your query or GORM`)
	if err := ErrChecker.Check(err); err != nil {
		return -1, err
	}
	return 0, nil
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
