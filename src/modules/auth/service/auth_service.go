package service

import (
	authCus "gorm/src/modules/auth/models"
	"gorm/src/modules/user/models"
	"gorm/src/modules/user/service"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	userService service.UserService
}

func NewAuthService(userService service.UserService) AuthService {
	return &authService{userService: userService}
}

func (s *authService) Login(username, password string) (string, error) {
	user, _ := s.userService.FindUserByUsername(username)

	// Generate JWT
	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authService) ValidateToken(tokenString string) (bool, error) {
	token, err := parseJWT(tokenString)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func generateJWT(user models.User) (string, error) {
	// Define the token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &authCus.CustomClaims{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Token válido por 24 horas
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func parseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// funcion para obtener el usuario del token
func GetUserFromToken(tokenString string) (models.User, error) {
	token, err := parseJWT(tokenString)
	if err != nil {
		return models.User{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return models.User{}, err
	}

	user := models.User{
		Username: claims["sub"].(string),
		// Agrega más campos según sea necesario
	}
	return user, nil
}
