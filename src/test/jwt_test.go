package scheduler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

var secrets = "2345345"

func TestJWT(t *testing.T) {
	getToken("kkk")
	getUser("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVfdGltZSI6MTIzLCJ1c2VyX2lkIjoia2trIn0.hvUMa4RyjtocISv6V9eStSrlAMy-a80RDW-hmwUr7Qg")
}

func getToken(user string) {
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user,
		"create_time": 123,
	})

	fmt.Println(token2.SignedString([]byte(secrets)))
}

func getUser(tokenStr string) {
	// 检查token是否在redis中存在，不存在则拒绝
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte(secrets), nil
	})
	if err != nil {
		// token 竟然搞不出来，这应该是程序错误
		return
	}
	// 这里从jwt中反解出userId，从而可以进行下一步操作
	userId := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok = claims["user_id"].(string)
		if !ok {
			return
		}
	}
	fmt.Println(userId)
}
