import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/dgrijalva/jwt-go"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/labstack/echo"
)

// JWTConfig is the struct for JWT configuration
type JWTConfig struct {
    SigningKey string
}

// User is the struct for user data
type User struct {
    gorm.Model
    Username string `json:"username" form:"username"`
    Password string `json:"password" form:"password"`
}

// CheckPassword checks if the password is correct
func (u *User) CheckPassword(password string) bool {
    return u.Password == password
}

// CreateToken creates a JWT token
func CreateToken(signingKey string, username string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    tokenString, err := token.SignedString([]byte(signingKey))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// VerifyToken verifies the JWT token
func VerifyToken(signingKey string, tokenString string) (jwt.Claims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(signingKey), nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
