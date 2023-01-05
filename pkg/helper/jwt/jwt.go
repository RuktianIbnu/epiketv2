package jwt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenMetadata ...
type TokenMetadata struct {
	Id      int64
	Nip     string
	Nama    string
	Id_role int64
}

// CreateToken ...
func CreateToken(idUser int64, nip, nama string, id_role int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = idUser
	claims["nip"] = nip
	claims["Nama"] = nama
	claims["id_role"] = id_role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyToken(headerToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid ...
func TokenValid(headerToken string) error {
	token, err := verifyToken(headerToken)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// ExtractTokenMetadata ...
func ExtractTokenMetadata(headerToken string) (*TokenMetadata, error) {
	token, err := verifyToken(headerToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["id"]), 10, 64)
		if err != nil {
			// fmt.Println(err.Error())
			return nil, err
		}

		nip, ok := claims["nip"].(string)
		if !ok {
			return nil, err
		}

		nama, ok := claims["nama"].(string)
		if !ok {
			return nil, err
		}

		id_role, ok := claims["id_role"].(int64)
		if !ok {
			return nil, err
		}

		return &TokenMetadata{
			Id:      id,
			Nip:     nip,
			Nama:    nama,
			Id_role: id_role,
		}, nil
	}
	return nil, err
}
