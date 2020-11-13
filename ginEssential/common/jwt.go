package common

import (
	"ginEssential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtkey = []byte("a_secret_crect")   //构建一个jwt密钥.影响JWT协议第三部分字段的内容

type claims struct {
	UserID uint
	jwt.StandardClaims
}


func ReleaseToken(user model.User)(string,error){
	expirationTime := time.Now().Add(7*24*time.Hour)     //token的有效时间
	claims := &claims{
		UserID: user.ID,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "oceanlearn.tech",
			Subject: "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err := token.SignedString(jwtkey)
	if err != nil{
		return "", err
	}
	return tokenString,nil

}

func ParseToken(tokenString string)(*jwt.Token, *claims, error){        //从tokenstring中解析出claims并返回
	claims := &claims{}
	token , err := jwt.ParseWithClaims(tokenString,claims,func(token *jwt.Token)(i interface{},err error){
		return jwtkey,err
	})
	return token, claims,err
}