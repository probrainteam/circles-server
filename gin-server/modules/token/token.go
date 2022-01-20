package token

import (
	. "circlesServer/modules/component"
	. "circlesServer/modules/reader"
	. "circlesServer/modules/storage"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}
type RefreshDetails struct {
	RefreshUuid string
	UserId      uint64
}

var client *redis.Client

var ACCESS_SECRET string
var REFRESH_SECRET string

func init() {
	ACCESS_SECRET = GetConfig(`token.ACCESS_SECRET`)
	REFRESH_SECRET = GetConfig(`token.REFRESH_SECRET`)
}
func ExtractToken(r *http.Request) []string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 3 { // refresh token require
		return strArr[1:3]
	} else if len(strArr) == 2 { // else auth api
		return strArr[1:2]
	}
	return nil
}
func VerifyToken(r *http.Request) (accessToken *jwt.Token, refreshToken *jwt.Token, err error) {
	tokenString := ExtractToken(r)
	accessToken, err = jwt.Parse(tokenString[0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ACCESS_SECRET), nil
	})
	if err != nil {
		return nil, nil, err
	}
	if len(tokenString) == 2 { // refresh token require
		refreshToken, err = jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(REFRESH_SECRET), nil
		})
		if err != nil {
			return nil, nil, err
		}
		return accessToken, refreshToken, nil
	}
	return accessToken, nil, nil
}
func ExtractBothTokenMetadata(r *http.Request) (*AccessDetails, *RefreshDetails, error) {
	accessToken, refreshToken, err := VerifyToken(r)
	if err != nil {
		return nil, nil, err
	}
	accClaims, ok := accessToken.Claims.(jwt.MapClaims)

	if refreshToken != nil {
		refClaims, ok_ := refreshToken.Claims.(jwt.MapClaims)
		if ok && ok_ && accessToken.Valid && refreshToken.Valid {
			accessUuid, ok := accClaims["access_uuid"].(string)
			if !ok {
				return nil, nil, err
			}
			userId, err := strconv.ParseUint(fmt.Sprintf("%.f", accClaims["user_id"]), 10, 64)
			if err != nil {
				return nil, nil, err
			}
			refreshUuid, ok := refClaims["refresh_uuid"].(string)
			if !ok {
				return nil, nil, err
			}
			return &AccessDetails{
					AccessUuid: accessUuid,
					UserId:     userId,
				}, &RefreshDetails{
					RefreshUuid: refreshUuid,
					UserId:      userId,
				}, nil
		}
	}
	if ok && accessToken.Valid {
		accessUuid, ok := accClaims["access_uuid"].(string)
		if !ok {
			return nil, nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", accClaims["user_id"]), 10, 64)
		if err != nil {
			return nil, nil, err
		}
		return &AccessDetails{
				AccessUuid: accessUuid,
				UserId:     userId,
			}, nil,
			nil
	}
	return nil, nil, err
}

func FetchAuth(at *AccessDetails) (uint64, error) {
	client, err := Redis()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	userid, err := client.Get(at.AccessUuid).Result()
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userid, 10, 64)
	if at.UserId != userID {
		return 0, errors.New("unauthorized")
	}
	return userID, nil
}

func CreateToken(userid uint64) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(REFRESH_SECRET))
	if err != nil {
		return nil, err
	}
	return td, nil
}
func CreateAuth(userid uint64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	client, err := Redis()
	if err != nil {
		return err
	}
	defer client.Close()
	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
func DeleteAuth(accessUuid string, refreshUuid string) (int64, error) {
	client, err := Redis()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	deleted, err := client.Del(accessUuid, refreshUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func CheckTokenAuth(r *http.Request) (bool, error) { // access token의 만료 여부 체크
	au, _, err := ExtractBothTokenMetadata(r)
	if err != nil {
		return false, err
	}
	// request 의 Access 토큰을 추출
	valid, err := checkTokenAlive(au.AccessUuid)
	if err != nil {
		return false, err
	}
	// CheckAccessToken() 호출 : 추출한 AccessToken의 만료 여부를 검사
	return valid, nil
}

// 후 CheckRefreshToken() 호출 : RefreshToken의 만료 여부를 검사
// Refresh 만료   -> 강제 로그아웃 후 재로그인 요청
// Refresh 만료 X -> ReissueAccessToken() : AccessToken 재발급
func checkTokenAlive(uuid string) (bool, error) {
	client, err := Redis()
	if err != nil {
		return false, err
	}
	defer client.Close()
	_, err = client.Get(uuid).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
func ReissueAccessToken(r *http.Request) (string, error) {
	client, err := Redis()
	if err != nil {
		return "", err
	}
	defer client.Close()
	// request token 파싱하여 userid get
	userid, err := GetCircleNumRef(strings.Split(r.Header.Get("Authorization"), " ")[1])
	if err != nil {
		return "", err
	}
	td, err := CreateToken(userid)
	if err != nil {
		return "", err
	}
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC
	now := time.Now()
	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return "", errAccess
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	att := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = att.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return "", err
	}

	return td.AccessToken, nil
}
