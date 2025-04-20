package services

import (
	"snipz/internal/storage/repository"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

type TokenService struct {
	token    *paseto.Token
	key      *paseto.V4SymmetricKey
	parser   *paseto.Parser
	duration time.Duration
}

func NewTokenService() *TokenService {
	token := paseto.NewToken()
	key := paseto.NewV4SymmetricKey()
	parser := paseto.NewParser()
	duration := 1 * time.Hour

	return &TokenService{
		&token,
		&key,
		&parser,
		duration,
	}
}

type TokenPayload struct {
	ID     uuid.UUID
	UserID int64
}

func (ts *TokenService) CreateToken(user repository.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := TokenPayload{
		ID:     id,
		UserID: user.ID,
	}

	err = ts.token.Set("payload", payload)
	if err != nil {
		return "", err
	}

	issuedAt := time.Now()
	expireAt := issuedAt.Add(ts.duration)

	ts.token.SetIssuedAt(issuedAt)
	ts.token.SetExpiration(expireAt)

	token := ts.token.V4Encrypt(*ts.key, nil)

	return token, nil
}

func (ts *TokenService) VerifyToken(token string) (*TokenPayload, error) {
	var payload *TokenPayload

	parsedToken, err := ts.parser.ParseV4Local(*ts.key, token, nil)
	if err != nil {
		return nil, err
	}

	err = parsedToken.Get("payload", &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
