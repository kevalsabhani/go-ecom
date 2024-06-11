package auth

import "testing"

func TestHashPasswor(t *testing.T) {
	hash, err := HashedPassword("password")
	if err != nil {
		t.Errorf("error hashing password:%v", err)
	}

	if hash == "" {
		t.Errorf("hash should not be empty")
	}

	if hash == "password" {
		t.Errorf("expected hash to be different from password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashedPassword("password")
	if err != nil {
		t.Errorf("error hashing password:%v", err)
	}

	if !ComparePasswords(hash, "password") {
		t.Errorf("expected password to match")
	}

	if ComparePasswords(hash, "wrong") {
		t.Errorf("expected password to not match")
	}
}
