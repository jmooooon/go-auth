package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmooooon/go-auth/pkg/api"
)

// GetToken ...
func GetToken(id int, userType int) (string, error) {
	var secret = []byte("Lakad-Matatataaaaaaaaag")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
		
	claims["id"] = id
	claims["ut"] = userType
	claims["time"] = time.Now().Unix()
	
	res, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return res, err
}

// TokenDecrypt ...
func TokenDecrypt(req string) (*api.TokenResponce, error) {
	claims:=jwt.MapClaims{}
	var res api.TokenResponce

	_, err := jwt.ParseWithClaims(req,claims,func(token *jwt.Token)(interface{},error){
		return []byte("Lakad-Matatataaaaaaaaag"),nil
	})

	if err != nil {
		return nil, err
	}

	for k,v := range claims{
		
		if k=="id"{
			res.Id=int32(v.(float64))
		}

		if k=="ut"{
			res.UserType=int32(v.(float64))
		}

	}
	return &res, nil
}
