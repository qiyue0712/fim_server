package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	hash := HashPwd("1234567")
	fmt.Println(hash)
}

func TestCheckPwd(t *testing.T) {
	//$2a$04$ajIdj44I7aErml8Kwnk7uOo0ZZHuOHahD9WRZaAqwjO642Rtrf8eS
	//$2a$04$6wmykKpp.mGEBi6.fc6CrOc1u64B8MiD3Q7Dp9TcPqfvNgH7SD2X.
	ok := CheckPwd("$2a$04$ajIdj44I7aErml8Kwnk7uOo0ZZHuOHahD9WRZaAqwjO642Rtrf8eS", "123456")
	fmt.Println(ok)
}
