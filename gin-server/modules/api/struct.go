package api

type ResgisterForm struct {
	Email string `json:"email"`
	PW    string `json:"pw"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
type LoginForm struct {
	ID string `json:"Email"`
	PW string `json:"pw"`
}
type ModifyForm struct {
	UID int    `json:"uid"`
	PW  string `json:"pw"`
	NEW string `json:"new"`
}
type Project struct {
	PID         int    `json:"pid"`
	UID         int    `json:"uid"`
	TITLE       string `json:"title"`
	DESCRIPTION string `json:"desc"`
	TOTAL       int    `json:"total"`
	TERM        int    `json:"term"`
	DUE         string `json:"due"`
	PATH        string `json:"path"`
	FE          int    `json:"fe"`
	BE          int    `json:"be"`
	AOS         int    `json:"aos"`
	IOS         int    `json:"ios"`
	PM          int    `json:"pm"`
	DESIGNER    int    `json:"designer"`
	DEVOPS      int    `json:"devops"`
	ETC         int    `json:"etc"`
}
type AddMemberForm struct {
	TITLE string `json:"title" form:"title"`
	DESC  string `json:"desc" form:"desc"`
	TOTAL int    `json:"total" form:"total"`
	TERM  int    `json:"term" form:"term"`
	DUE   string `json:"due" form:"due"`
	PATH  string `json:"path" form:"path"`
	FE    int    `json:"fe" form:"fe"`
	BE    int    `json:"be" form:"be" `
	AOS   int    `json:"aos" form:"aos"`
	IOS   int    `json:"ios" form:"ios"`
}
type JoinForm struct {
	PID      int    `json:"pid"`
	UID      int    `json:"uid"`
	CATEGORY string `json:"category"`
}
type ReplyJoinForm struct {
	PID int `json:"pid"`
	UID int `json:"uid"`
}
type Announce struct {
	TITLE   string `json:"title"`
	CONTENT string `json:"content"`
}
type Msg struct {
	TYPE    int    `json:"type"`
	TITLE   string `json:"title"`
	CONTENT string `json:"content"`
	PID     int    `json:"pid"`
	UID     int    `json:"uid"`
}
type member struct {
}
