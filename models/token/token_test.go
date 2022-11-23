package token_test

import (
	"testing"
	"token/models/token"
	"token/models/unittest"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	err := unittest.PrepareTestDB()
	assert.NoError(t, err)

	// test init db with empty token
	tk, err := token.GetToken()
	assert.NoError(t, err)
	assert.Equal(t, token.TokenID, tk.ID)
	assert.Equal(t, tk.Token, "")
}

func TestUpdateToken(t *testing.T) {
	err := unittest.PrepareTestDB()
	assert.NoError(t, err)

	// error when token is empty
	err = token.UpdateToken("")
	assert.Error(t, err)

	// update token
	err = token.UpdateToken("test")
	assert.NoError(t, err)

	// check token updated
	tk, err := token.GetToken()
	assert.NoError(t, err)
	assert.Equal(t, "test", tk.Token)
}

func TestDeleteToken(t *testing.T) {
	err := unittest.PrepareTestDB()
	assert.NoError(t, err)

	// delete token
	err = token.DeleteToken()
	assert.NoError(t, err)

	// check token is empty
	tk, err := token.GetToken()
	assert.NoError(t, err)
	assert.Equal(t, tk.Token, "")
}

func TestGenToken(t *testing.T) {
	// token random
	tk1 := token.GenToken()
	tk2 := token.GenToken()
	assert.NotEmpty(t, tk1)
	assert.NotEmpty(t, tk2)
	assert.NotEqual(t, tk1, tk2)
}

func TestCheckToken(t *testing.T) {
	err := unittest.PrepareTestDB()
	assert.NoError(t, err)

	// empty token is invalid
	valid, err := token.CheckToken("")
	assert.NoError(t, err)
	assert.False(t, valid)

	// token invalid
	valid, err = token.CheckToken("test")
	assert.NoError(t, err)
	assert.False(t, valid)

	// old token invalid, new token valid
	err = token.UpdateToken("test2")
	assert.NoError(t, err)
	valid, err = token.CheckToken("test")
	assert.NoError(t, err)
	assert.False(t, valid)
	valid, err = token.CheckToken("test2")
	assert.NoError(t, err)
	assert.True(t, valid)

}

func TestEncryptDecryptToken(t *testing.T) {
	tk := "test"

	// plain token encrypt
	encTk, err := token.EncryptToken(tk)
	assert.NotEmpty(t, encTk)
	assert.NotEqual(t, tk, encTk)
	assert.NoError(t, err)

	// token decrypt to plain
	decTk, err := token.DecryptToken(encTk)
	assert.NoError(t, err)
	assert.Equal(t, tk, decTk)
}
