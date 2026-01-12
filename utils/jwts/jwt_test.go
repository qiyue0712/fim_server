package jwts

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "wjy",
	}, "12345", 8)
	fmt.Println(token, err)
}

func TestParseToken(t *testing.T) {
	payload, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IndqeSIsInJvbGUiOjEsImV4cCI6MTc2ODIzODM2MH0.FlBHYcb4XUPiDR8bcX2Q21Q7O6mPMqg00CUqYqx0e-4", "12345")
	fmt.Println(payload, err)
}
