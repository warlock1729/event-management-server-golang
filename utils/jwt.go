package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "dekdnedned"

func GenerateToken(email string, user_id int64) (string, error) {
	// this is default map too
	fmt.Println(user_id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": user_id,
		// exp is used by jwt internally too
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(token string) (*jwt.Token,error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {

		// HS256 is a type under HMAC
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil,errors.New("Could not parse token")
	}

	if !parsedToken.Valid {
		return nil,errors.New("Invalid token !")
	}

	return parsedToken,nil
}

func ExtractJWTClaims(parsedToken *jwt.Token) (int64, error) {
	// we have used specific type of claims called map claims. Hence checking if the type is same or not
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0,errors.New("Invalid token claims")
	}

	// email := claims["email"].(string)
	user_id := int64(claims["user_id"].(float64))
	return user_id,nil
}


