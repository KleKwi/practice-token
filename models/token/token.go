package token

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"token/models/db"

	"token/util"

	gouuid "github.com/google/uuid"
)

const TokenID uint64 = 1

type Token struct {
	ID    uint64 `xorm:"pk autoincr"`
	Token string `xorm:"unique"`
}

func init() {
	db.RegisterModel(new(Token), func() (err error) {
		_, err = db.FirstOrCreate(&Token{ID: TokenID})
		if err != nil {
			return
		}

		return
	})
}

func GetToken() (token Token, err error) {
	tk := new(Token)
	_, err = db.GetEngine().ID(TokenID).Get(tk)
	if err != nil {
		return
	}

	tk.Token, err = DecryptToken(tk.Token)
	if err != nil {
		return
	}

	token = *tk

	return
}

func UpdateToken(tk string) (err error) {
	if tk == "" {
		err = errors.New("token is empty")
		return
	}

	encToken, err := EncryptToken(tk)
	if err != nil {
		return
	}

	token := Token{Token: encToken}
	_, err = db.GetEngine().ID(TokenID).AllCols().Update(&token)
	if err != nil {
		return
	}

	return err
}

func DeleteToken() (err error) {
	token := Token{Token: ""}
	_, err = db.GetEngine().ID(TokenID).AllCols().Update(&token)
	if err != nil {
		return
	}

	return
}

func CheckToken(tk string) (valid bool, err error) {
	if tk == "" {
		return
	}

	token, err := GetToken()
	if err != nil {
		return
	}

	if token.Token == tk {
		valid = true
		return
	}

	return
}

func GenToken() (token string) {
	token = encodeSha1(gouuid.New().String())
	return
}

func EncryptToken(tk string) (encToken string, err error) {
	rc4Tk, err := util.Rc4([]byte(tk))
	if err != nil {
		return
	}

	encToken = util.Base64Encode(rc4Tk)

	return
}

func DecryptToken(tk string) (decToken string, err error) {
	decTk, err := util.Base64Decode(tk)
	if err != nil {
		return
	}

	rc4Tk, err := util.Rc4(decTk)
	if err != nil {
		return
	}

	decToken = string(rc4Tk)

	return
}

// encodeSha1 string to sha1 hex value.
func encodeSha1(str string) string {
	h := sha1.New()
	_, _ = h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
