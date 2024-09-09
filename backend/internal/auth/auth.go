package auth

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"automated-seo-analyzer/backend/internal/database"
	"automated-seo-analyzer/backend/internal/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(username, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}

	log.Printf("Attempting to insert user: %s", username)
	_, err = database.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	if err != nil {
		log.Printf("Error inserting user into database: %v", err)
		return err
	}

	log.Printf("Successfully created user: %s", username)
	return nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := database.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserIDFromToken(r *http.Request) (int, error) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	user, err := GetUserByUsername(claims.Username)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func SetUserInContext(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, "username", username)
}

func GetUserFromContext(ctx context.Context) string {
	if username, ok := ctx.Value("username").(string); ok {
		return username
	}
	return ""
}
