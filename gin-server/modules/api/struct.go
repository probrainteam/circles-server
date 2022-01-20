package api

type ResgisterForm struct {
	Email  string `json:"email"`
	PW     string `json:"pw"`
	PUBKEY string `json:"pubkey"`
}
type LoginForm struct {
	ID  string `json:"Email"`
	PW  string `json:"pw"`
	ENC string `json:"enc"`
}
type ModifyForm struct {
	PW  string `json:"pw"`
	NEW string `json:"new"`
}
type Member struct {
	SID    string `json:"student_id"`
	MAJOR  string `json:"major"`
	NAME   string `json:"name"`
	YEAR   int    `json:"year"`
	EMAIL  string `json:"email"`
	PHONE  string `json:"phone"`
	PAID   bool   `json:"paid"`
	STATUS int    `json:"status"`
}
type ReplyJoinForm struct {
	SID int `json:"sid"`
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
