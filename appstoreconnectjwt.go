package appstoreconnectjwt

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"strings"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrPrivateKeyNotValidPEM   = errors.New("pk is not a valid PEM type")
	ErrPrivateKeyNotValidPKCS8 = errors.New("pk must be a encoded PKCS#8 type")
	ErrPrivateKeyNotECDSA      = errors.New("pk must be of ECDSA type")
	ErrConfigIsNil             = errors.New("config is nil")
)

type Config struct {
	ISS string
	KID string
	// ExpiresIn defines when the token should expire.
	// It is AppStore's responsibility to enforce token expiration check.
	// This field is only used to set `exp` field in Claims
	// defined at https://tools.ietf.org/html/rfc7519#section-4.1.4.
	//
	// This field should not be set more than 20 minutes in the future
	// for AppStore to consider it valid.
	ExpiresIn time.Duration
	AUD       string
	PK        string
}

type Client struct {
	bearer string
	token  *jwt.Token
	claims *jwt.StandardClaims
	cfg    *Config
	lock   sync.Mutex
}

// New is a constructor that creates new client with valid jwt token.
func New(cfg *Config) (*Client, error) {
	if cfg == nil {
		return nil, ErrConfigIsNil
	}
	return &Client{
		cfg: cfg,
	}, nil
}

// BearerToken returns a signed JWT token for AppStoreConnect.
// This method handles token reuse.
func (c *Client) BearerToken() (string, error) {
	if err := c.newIfExpired(); err != nil {
		return "", err
	}
	return c.bearer, nil
}

func (c *Client) expireDuration() time.Duration {
	return c.cfg.ExpiresIn
}

// newIfExpired generates a new bearer token only if previously issued token is expired.
// This method does not handle clock skew.
func (c *Client) newIfExpired() error {
	defer c.lock.Unlock()
	c.lock.Lock()

	if c.bearer == "" {
		if err := c.newBearerTokenString(); err != nil {
			return err
		}
		return nil
	}

	t := time.Unix(c.claims.IssuedAt, 0)
	delta := time.Since(t)
	if delta >= (c.expireDuration()) {
		// token is expired, create new token
		return c.newBearerTokenString()
	}

	return nil
}

func (c *Client) newBearerTokenString() error {
	reader := strings.NewReader(c.cfg.PK)
	pk, err := privateKeyFromReader(reader)
	if err != nil {
		return err
	}

	c.claims = &jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Add(c.expireDuration()).Unix(),
		Issuer:    c.cfg.ISS,
		Audience:  c.cfg.AUD,
		IssuedAt:  time.Now().UTC().Unix(),
	}

	c.token = jwt.NewWithClaims(jwt.SigningMethodES256, c.claims)
	c.token.Header["kid"] = c.cfg.KID

	c.bearer, err = c.token.SignedString(pk)
	if err != nil {
		return err
	}
	return nil
}

type ReaderLength interface {
	Read(b []byte) (n int, err error)
	Len() int
}

func privateKeyFromReader(rl ReaderLength) (*ecdsa.PrivateKey, error) {
	b := make([]byte, rl.Len())
	for {
		_, err := rl.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	by, _ := pem.Decode(b)
	if by == nil {
		return nil, ErrPrivateKeyNotValidPEM
	}
	key, err := x509.ParsePKCS8PrivateKey(by.Bytes)
	if err != nil {
		return nil, ErrPrivateKeyNotValidPKCS8
	}
	switch pk := key.(type) {
	case *ecdsa.PrivateKey:
		return pk, nil
	default:
		return nil, ErrPrivateKeyNotECDSA
	}
}
