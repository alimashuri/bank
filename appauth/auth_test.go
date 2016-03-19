package appauth

import (
	"crypto/sha512"
	"encoding/hex"
	"testing"

	"github.com/ksred/bank/configuration"
)

func AccountsSetConfig(t *testing.T) {
	// Load app config
	Config, err := configuration.LoadConfig()
	if err != nil {
		t.Errorf("TestAccounts.SetConfig: %v", err)
	}
	// Set config in packages
	SetConfig(&Config)
}

func TestProcessAppAuthParams(t *testing.T) {
	data := []string{}
	_, err := ProcessAppAuth(data)
	if err == nil {
		t.Errorf("ProcessAppAuthParams does not pass. Looking for %v, got %v", "Too few parameters", nil)
	}

	data = []string{"", "1"}
	_, err = ProcessAppAuth(data)
	if err == nil {
		t.Errorf("ProcessAppAuthParams does not pass. Looking for %v, got %v", "Too few parameters", nil)
	}

	data = []string{"", "", "2", ""}
	_, err = ProcessAppAuth(data)
	if err == nil {
		t.Errorf("ProcessAppAuthParams does not pass. Looking for %v, got %v", "Too few parameters", nil)
	}

	data = []string{"", "", "3", ""}
	_, err = ProcessAppAuth(data)
	if err == nil {
		t.Errorf("ProcessAppAuthParams does not pass. Looking for %v, got %v", "Too few parameters", nil)
	}
}

func TestCreateRemoveUserPassword(t *testing.T) {
	AccountsSetConfig(t)

	user := "1234-1234-1234-1234"
	password := "test-password"

	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	_, err := CreateUserPassword(user, password)
	if err != nil {
		t.Errorf("CreateRemoveUserPassword Create does not pass. Looking for %v, got %v", nil, err)
	}

	_, err = RemoveUserPassword(user, hashedPassword)
	if err != nil {
		t.Errorf("CreateRemoveUserPassword Remove does not pass. Looking for %v, got %v", nil, err)
	}
}

func TestCreateRemoveCheckToken(t *testing.T) {
	AccountsSetConfig(t)

	user := "1234-1234-1234-1234"
	password := "test-password"

	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	_, err := CreateUserPassword(user, password)
	if err != nil {
		t.Errorf("CreateRemoveCheckToken Create does not pass. Looking for %v, got %v", nil, err)
	}

	token, err := CreateToken(user, password)
	if err != nil {
		t.Errorf("CreateRemoveCheckToken Create does not pass. Looking for %v, got %v", nil, err)
	}

	_, err = CheckToken(token)
	if err != nil {
		t.Errorf("CreateRemoveCheckToken Check does not pass. Looking for %v, got %v", nil, err)
	}

	_, err = RemoveToken(token)
	if err != nil {
		t.Errorf("CreateRemoveCheckToken Delete does not pass. Looking for %v, got %v", nil, err)
	}

	_, err = RemoveUserPassword(user, hashedPassword)
	if err != nil {
		t.Errorf("CreateRemoveCheckToken Remove does not pass. Looking for %v, got %v", nil, err)
	}

}

func TestGetUserFromToken(t *testing.T) {
	AccountsSetConfig(t)

	user := "1234-1234-1234-1234"
	password := "test-password"

	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	_, err := CreateUserPassword(user, password)
	if err != nil {
		t.Errorf("GetUserFromToken Create does not pass. Looking for %v, got %v", nil, err)
	}

	token, err := CreateToken(user, password)
	if err != nil {
		t.Errorf("GetUserFromToken Create does not pass. Looking for %v, got %v", nil, err)
	}

	userFromToken, err := GetUserFromToken(token)
	if err != nil {
		t.Errorf("GetUserFromToken Check does not pass. Looking for %v, got %v", nil, err)
	}

	if userFromToken != user {
		t.Errorf("GetUserFromToken GetFromToken does not pass. Looking for %v, got %v", user, userFromToken)
	}

	_, err = RemoveToken(token)
	if err != nil {
		t.Errorf("GetUserFromToken Delete does not pass. Looking for %v, got %v", nil, err)
	}

	_, err = RemoveUserPassword(user, hashedPassword)
	if err != nil {
		t.Errorf("GetUserFromToken Remove does not pass. Looking for %v, got %v", nil, err)
	}

}
