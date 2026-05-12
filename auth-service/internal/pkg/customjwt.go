package pkg

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	privateKey   *rsa.PrivateKey
	publicKey    *rsa.PublicKey
	refreshStore *refreshMemoryStore
}

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	jwt.RegisteredClaims
}

// простое in-memory хранилище refresh-токенов
type refreshMemoryStore struct {
	sync.Mutex
	tokens map[string]string
}

func NewJWTManager(privatePath, publicPath string) (*JWTManager, error) {
	privBytes, err := ioutil.ReadFile(privatePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read private key: %w", err)
	}
	pubBytes, err := ioutil.ReadFile(publicPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, fmt.Errorf("invalid public key: %w", err)
	}

	return &JWTManager{
		privateKey:   privateKey,
		publicKey:    publicKey,
		refreshStore: &refreshMemoryStore{tokens: make(map[string]string)},
	}, nil
}

func (m *JWTManager) GenerateAccessToken(username string, ttl time.Duration) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			Issuer:    "Pidor",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(m.privateKey)
}

func (m *JWTManager) GenerateRefreshToken(username string, ttl time.Duration) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			Issuer:    "Pidor",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(m.privateKey)
	if err != nil {
		return "", err
	}
	m.refreshStore.Lock()
	m.refreshStore.tokens[signed] = username
	m.refreshStore.Unlock()
	return signed, nil
}

func (m *JWTManager) VerifyAccessToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.publicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}
	return claims, nil
}

func (m *JWTManager) RefreshAccessToken(refreshToken string, ttl time.Duration) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return m.publicKey, nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}
	m.refreshStore.Lock()
	username, ok := m.refreshStore.tokens[refreshToken]
	m.refreshStore.Unlock()
	if !ok || username != claims.Username {
		return "", errors.New("refresh token not recognized")
	}
	return m.GenerateAccessToken(username, ttl)
}

func (m *JWTManager) RevokeRefreshToken(token string) {
	m.refreshStore.Lock()
	delete(m.refreshStore.tokens, token)
	m.refreshStore.Unlock()
}
