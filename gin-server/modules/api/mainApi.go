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
func GetMemberList(c *gin.Context) ([]Project, error) {
	db := DB()
	var length int
	_ = db.QueryRow(`your query or GORM`).Scan(&length)
	if length == 0 {
		return []Project{}, errors.New("Nothing to show")
	}
	rows, err := db.Query(`your query or GORM`)
	if err := ErrChecker.Check(err); err != nil {
		return []Project{}, err
	}
	defer rows.Close()
	projects := make([]Project, 0)
	var pos Project
	for rows.Next() {
		err := rows.Scan(&pos.PID, &pos.UID, &pos.TITLE,
			&pos.TOTAL, &pos.DESCRIPTION, &pos.DUE, &pos.TERM, &pos.PATH)
		if err := ErrChecker.Check(err); err != nil {
			return []Project{}, err
		}
		projects = append(projects, pos)
	}
	return projects, nil
}
func AddMember(c *gin.Context) (int, error) {
	var reqBody AddMemberForm

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
	dirPath := "/project_img/"          // custom
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
func Join(c *gin.Context) ([]member, error) {
	var reqBody JoinForm
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return []member{}, err
	}
	db := DB()
	_, err = db.Exec(`your query or GORM`)
	if err != nil {
		return []member{}, err
	}
	list := make([]member, 0)
	return list, nil
}
func GetNumProject(c *gin.Context) (int, error) {
	var reqBody JoinForm
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
