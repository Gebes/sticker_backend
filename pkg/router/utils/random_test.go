package utils

import (
	"log"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {

	length := 128
	s, _ := GenerateRandomString(length)
	log.Println(s)
	if len(s) != length {
		t.Errorf("GenerateRandomString() length got = %v, want %v", len(s), length)
	}
}
