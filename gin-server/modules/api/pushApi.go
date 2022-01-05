package api

import (
	"errors"

	ErrChecker "circlesServer/modules/errors"
	"circlesServer/modules/storage"

	"github.com/gin-gonic/gin"
)

func getAnnouncement(c *gin.Context) ([]Msg, error) {
	annoList := make([]Msg, 0)
	db := storage.DB()
	rows, err := db.Query(`your query or GORM`)
	if err != nil {
		return []Msg{}, errors.New("Nothing")
	}
	defer rows.Close()

	for rows.Next() {
		var title, content string
		if err := rows.Scan(&title, &content); err != nil {
			return []Msg{}, err
		}
		m := Msg{
			TYPE:    0,
			TITLE:   title,
			CONTENT: content,
			PID:     0,
			UID:     0,
		}
		annoList = append(annoList, m)
	}

	return annoList, nil
}
func Announcement(c *gin.Context) error {
	var reqBody Announce
	err := c.ShouldBindJSON(&reqBody)
	if err := ErrChecker.Check(err); err != nil {
		return err
	}
	db := storage.DB()
	_, err = db.Exec(`your query or GORM`)
	if err != nil {
		return err
	}
	return nil
}
