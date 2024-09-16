package helper

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword"
	expectedHash := "88dde1469b43f8baf4f6340c0590f76e" // 预期的哈希值

	hashedPassword := HashPassword(password)

	if hashedPassword != expectedHash {
		t.Fatalf("哈希值不匹配: got %v, want %v", hashedPassword, expectedHash)
	}
}
