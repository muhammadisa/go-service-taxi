package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Authenticated struct
type Authenticated struct {
	UserID       uuid.UUID `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	AccessToken  string    `json:"access_token"`
}

// HashPassword hashing password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compare hashed password with password string
func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateToken generate JWT token
func GenerateToken(id uuid.UUID) (string, string, error) {
	// AccessToken
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id.String()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", "", err
	}

	// RefreshToken
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["user_id"] = id.String()
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refresh, err := refreshToken.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", "", err
	}
	return token, refresh, nil
}

// ExtractToken extract token from Bearer
func ExtractToken(c echo.Context) string {

	bearerToken := c.Request().Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""

}

// ExtractClaimsFromToken extracting jwt claims
func ExtractClaimsFromToken(c echo.Context, targetClaim string) (interface{}, error) {
	token, err := jwt.Parse(ExtractToken(c), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if targetClaim == "" {
			return claims, nil
		}
		return claims[targetClaim], nil
	}
	return nil, nil
}

// JWTExtractData validate jwt token
func JWTExtractData(c echo.Context) (string, error) {
	claims, err := ExtractClaimsFromToken(c, "user_id")
	if err != nil {
		return "", err
	}
	if claims != nil {
		return fmt.Sprintf("%s", claims), nil
	}
	return "", fmt.Errorf("Claims key maybe not found")
}

// JWTTokenValidate validate jwt token
func JWTTokenValidate(c echo.Context) error {
	claims, err := ExtractClaimsFromToken(c, "")
	if err != nil {
		return err
	}
	if claims != nil {
		Pretty(claims)
	}
	return nil
}

// Pretty string prettier
func Pretty(data interface{}) {
	_, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return
	}
}
