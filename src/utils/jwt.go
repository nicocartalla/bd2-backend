package utils


import (
	"github.com/golang-jwt/jwt/v4"
	"bd2-backend/src/config"
	"time"
	"fmt"
	"context"
	"github.com/mitchellh/mapstructure"
)
var secret []byte

func init() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		ErrorLogger.Fatal("cannot load config:", err)
	}
	secret = []byte(cfg.JwtKey)
}

type JwtPayload struct {
	Email  string `json:"email" mapstructure:"email"`
	RoleID int    `json:"role_id" mapstructure:"role_id"`
}

type MyJWTClaims struct {
    *jwt.RegisteredClaims
    UserInfo interface{}
}

func CreateToken(sub string, userInfo interface{}) (string, time.Time, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(time.Hour * 1)
	token.Claims = &MyJWTClaims{
			&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   sub,
		},
		userInfo,
	  }
	val, err := token.SignedString(secret)
	if err != nil {
		return "",exp, err
	}
	return val,exp, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)   (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

type claimskey int
var claimsKey claimskey

func SetJWTClaimsContext(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
  }
  
func jWTClaimsFromContext(ctx context.Context) (jwt.MapClaims, bool) {
	claims, ok := ctx.Value(claimsKey).(jwt.MapClaims)
	return claims, ok
}

func GetJwtPayloadFromClaim(ctx context.Context) (JwtPayload, error) {
	claims, ok := jWTClaimsFromContext(ctx)
	if !ok {
		return *new(JwtPayload), fmt.Errorf("error getting claims from context")
	}

	userInfo := claims["UserInfo"].(map[string]interface{})
		var jwtPayload JwtPayload
		err := mapstructure.Decode(userInfo, &jwtPayload)
		if err != nil {
			return *new(JwtPayload), err
		}
		return jwtPayload, nil

} 
