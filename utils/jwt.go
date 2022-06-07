package utils

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type JwtClaims struct {
	*jwt.StandardClaims
	Prv 		string `json:"prv"`
	UserId      string `json:"user_id"`
	Authorities string `json:"authorities"`
}

var (
	key    []byte = []byte("7b21ff514564ea1be967ff7433485b00")
	issuer string = "ipcloud-api"
)

//生成token令牌
func CreateToken(userId, apiId string, expireTime int64) (string, error) {
	fmt.Println(time.Now().Unix())
	claims := JwtClaims{
		&jwt.StandardClaims{
			IssuedAt:  	time.Now().Unix(),
			Subject:	userId,
			ExpiresAt: 	time.Now().Unix() + expireTime,
			Issuer:    	issuer,
			Id:        	apiId,
			NotBefore: 	time.Now().Unix(),
		},
		apiId,
		userId,
		"ROLE_ADMIN,AUTH_WRITE,ACTUATOR",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		log.Printf("parse date err： %v", err)
		return "", errors.New("生成Token异常，请检查参数")
	}
	return ss, nil
}

//销毁token
func DestoryToken() string {

	claims := JwtClaims{
		&jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() - 99999),
			Issuer:    issuer,
			Id:        "exit",
		},
		issuer,
		"exit",
		"ROLE_ADMIN,AUTH_WRITE,ACTUATOR",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		// log.Error(err)
		log.Printf("parse date err： %v", err)
		return ""
	}
	return ss
}

//检查token
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}

//解析jwt
func ParseJwt(token string) (*JwtClaims, error) {
	var jclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(token, jclaim, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err, token)
		return nil, errors.New("parase with claims failed.")
	}
	return jclaim, nil
}
