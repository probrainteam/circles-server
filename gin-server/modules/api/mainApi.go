package api

import (
	ErrChecker "circlesServer/modules/errors"
	. "circlesServer/modules/storage"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAddr() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
func GetMemberList(c *gin.Context) ([]Member, error) {
	db := DB()
	var length int
	_ = db.QueryRow(`select count(*) from probrain`).Scan(&length)
	if length == 0 {
		return []Member{}, errors.New("Nothing to show")
	}
	rows, err := db.Query(`select * from probrain`)
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
	return Members, nil
}
func AddMember(c *gin.Context) (int, error) {
	var reqBody Member

	err := c.ShouldBind(&reqBody)
	if err != nil {
		return 1, err
	}

	file, _, err := c.Request.FormFile("content")

	var pid int
	db := DB()
	db.QueryRow(`your query or GORM`)
	basePath := "http://" + getAddr()
	localPath := "/Users/macbook/Sites" // custom
	dirPath := "/Member_img/"           // custom
	path := basePath + dirPath + strconv.Itoa(pid+1) + `.png`

	if err != nil {
		return -1, err
	}
	_, err = db.Exec(`your query or GORM`)
	if err := ErrChecker.Check(err); err != nil {
		return -1, err
	}

	path = localPath + dirPath + strconv.Itoa(pid+1) + `.png`
	dst, err := os.Create(path)
	if err != nil {
		return -1, err
	}
	defer dst.Close()

	if err != nil {
		return -1, err
	}
	if _, err := io.Copy(dst, file); err != nil {
		return -1, err
	}
	_, err = db.Exec(`your query or GORM`)
	if err := ErrChecker.Check(err); err != nil {
		return -1, err
	}
	return pid, nil
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
