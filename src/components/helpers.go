package helpers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

// ValidateRequest validates the request based on the provided schema.
func ValidateRequest(schema map[string]interface{}, req *http.Request) error {
	validate := validator.New()
	err := validate.Struct(req.URL.Query())
	if err != nil {
		return err
	}
	return nil
}

// GenerateAccessToken generates an access token with the provided payload.
func GenerateAccessToken(payload map[string]interface{}) (string, error) {
	atClaims := jwt.MapClaims{}
	for k, v := range payload {
		atClaims[k] = v
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString([]byte("secret-key"))
}

// GenerateRefreshToken generates a refresh token with the provided payload.
func GenerateRefreshToken(payload map[string]interface{}) (string, error) {
	rtClaims := jwt.MapClaims{}
	for k, v := range payload {
		rtClaims[k] = v
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	return rt.SignedString([]byte("secret-key"))
}

// ParseAccessToken parses the provided token and returns the payload.
func ParseAccessToken(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}
	return claims, nil
}

// ParseRefreshToken parses the provided token and returns the payload.
func ParseRefreshToken(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, err
	}
	return claims, nil
}

// GetTimeInHours returns the current time in hours.
func GetTimeInHours() int {
	now := time.Now()
	hours := now.Hour()
	return hours
}

// GetTimeInMinutes returns the current time in minutes.
func GetTimeInMinutes() int {
	now := time.Now()
	minutes := now.Minute()
	return minutes
}

// GetTimeInSeconds returns the current time in seconds.
func GetTimeInSeconds() int {
	now := time.Now()
	seconds := now.Second()
	return seconds
}

// GetDurationString returns the provided duration in human-readable format.
func GetDurationString(duration time.Duration) string {
	d := time.Duration(duration)
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
}

// GetRandomString returns a random string of the provided length.
func GetRandomString(length int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	return string(b)
}