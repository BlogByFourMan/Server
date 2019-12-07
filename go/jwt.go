package swagger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// SecretKey : secret key for jwt
const SecretKey = "123qwe"

func ValidateToken(w http.ResponseWriter, r *http.Request) bool {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
		return false
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "token is invalid")
		return false
	}

	return true
}

func SignToken(userName string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["name"] = userName
	token.Claims = claims
	return token.SignedString([]byte(SecretKey))
}
