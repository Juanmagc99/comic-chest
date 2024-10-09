package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"time"

	"juanmagc99.comic-chest/internal/validator"
)

const (
	ScopeActivation = "activation"
)

type Token struct {
	Plaintext string
	Hash      []byte
	UserID    int64
	Expiry    time.Time
	Scope     string
}

type TokenModel struct {
	DB *sql.DB
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}

func generateToken(userID int64, ttl time.Duration, scope string) (*Token, error) {

	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)

	//Populate randomBytes with random data
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	//5 bits are used for every byte in our randomBytes (16*8)/5 = 25.6 characters = 26 we use no padding
	//so the program doesnt add "=" at the end
	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	//Create hash from out string token
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	return token, nil
}

func (m TokenModel) New(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)

	if err != nil {
		return nil, err
	}

	err = m.Insert(token)
	return token, err
}

func (m TokenModel) Insert(token *Token) error {

	query := `
		INSERT INTO tokens (hash, user_id, expiry, scope)
		VALUES ($1, $2, $3, $4)
	`

	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)

	return err
}

func (m TokenModel) DeleteAllForUser(scope string, userID int64) error {
	query := `
		DELETE FROM tokens
		WHERE scope = $1 AND user_id = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, scope, userID)

	return err
}
