package appstoreconnectjwt

import (
	"testing"
	"time"
)

const (
	// privateKeyValidFake satisfies the private key requirement
	// for generating JWT token.
	//
	// Follwing openssl command was used to generate this private key.
	// openssl ecparam -genkey -name prime256v1 -noout | \
	//      openssl pkcs8 -topk8 -nocrypt -in ec_private.pem
	//
	privateKeyValidFake = `
-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgLr73kREgj9lV5HHg
dFDBpMfzJ/Y/hdielkVexW5ML9OhRANCAASxG1wW2Wlth3tE3fxjAAf0bd83M1p8
m4B8ipZ3jY5hvTb9zbM8GbhExZotyZW/B27acARhMToQcFIpO3GurIrd
-----END PRIVATE KEY-----
`

	privateKeyInValidPEM = `
-----BEGIN-----
-----END-----`

	privateKeyNotPKCS8 = `
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIOMiG3gDoYSl4W7cajCoS3Fz6ZRefNLgvcBVGC+/+3FdoAoGCCqGSM49
AwEHoUQDQgAEmslhCNl+oO1R0ribBDsdROmEdXF2aJ4gDRxiLF626qbvUAu/SvBV
goIxuiHe8UW5+HzuZ6FLxvcRDtTDTx8mlg==
-----END EC PRIVATE KEY-----
`

	privateKeyNotECSDA = `
-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAKNwapOQ6rQJHetP
HRlJBIh1OsOsUBiXb3rXXE3xpWAxAha0MH+UPRblOko+5T2JqIb+xKf9Vi3oTM3t
KvffaOPtzKXZauscjq6NGzA3LgeiMy6q19pvkUUOlGYK6+Xfl+B7Xw6+hBMkQuGE
nUS8nkpR5mK4ne7djIyfHFfMu4ptAgMBAAECgYA+s0PPtMq1osG9oi4xoxeAGikf
JB3eMUptP+2DYW7mRibc+ueYKhB9lhcUoKhlQUhL8bUUFVZYakP8xD21thmQqnC4
f63asad0ycteJMLb3r+z26LHuCyOdPg1pyLk3oQ32lVQHBCYathRMcVznxOG16VK
I8BFfstJTaJu0lK/wQJBANYFGusBiZsJQ3utrQMVPpKmloO2++4q1v6ZR4puDQHx
TjLjAIgrkYfwTJBLBRZxec0E7TmuVQ9uJ+wMu/+7zaUCQQDDf2xMnQqYknJoKGq+
oAnyC66UqWC5xAnQS32mlnJ632JXA0pf9pb1SXAYExB1p9Dfqd3VAwQDwBsDDgP6
HD8pAkEA0lscNQZC2TaGtKZk2hXkdcH1SKru/g3vWTkRHxfCAznJUaza1fx0wzdG
GcES1Bdez0tbW4llI5By/skZc2eE3QJAFl6fOskBbGHde3Oce0F+wdZ6XIJhEgCP
iukIcKZoZQzoiMJUoVRrA5gqnmaYDI5uRRl/y57zt6YksR3KcLUIuQJAd242M/WF
6YAZat3q/wEeETeQq1wrooew+8lHl05/Nt0cCpV48RGEhJ83pzBm3mnwHf8lTBJH
x6XroMXsmbnsEw==
-----END PRIVATE KEY-----
`
)

var (
	unexpectedResult = "unexpected result: (%v) wanted (%v)"
)

func TestNew(t *testing.T) {
	cfg := &Config{
		ISS:       "ISS",
		KID:       "KID",
		ExpiresIn: 1 * time.Minute,
		AUD:       "AUD",
		PK:        privateKeyValidFake,
	}

	tcs := []struct {
		name     string
		config   *Config
		expected error
	}{
		{"Empty", nil, ErrConfigIsNil},
		{"NonEmpty", cfg, nil},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			_, got := New(tt.config)
			if got != tt.expected {
				t.Errorf(unexpectedResult, got, tt.config)
			}
		})
	}
}

func TestBearerToken(t *testing.T) {
	var tc1, _ = New(&Config{
		ISS:       "ISS",
		KID:       "KID",
		ExpiresIn: 1 * time.Minute,
		AUD:       "AUD",
		PK:        privateKeyValidFake,
	})
	var tc2, _ = New(&Config{
		ISS:       "ISS",
		KID:       "KID",
		ExpiresIn: 1 * time.Minute,
		AUD:       "AUD",
		PK:        privateKeyInValidPEM,
	})
	var tc3, _ = New(&Config{
		ISS:       "ISS",
		KID:       "KID",
		ExpiresIn: 1 * time.Minute,
		AUD:       "AUD",
		PK:        privateKeyNotPKCS8,
	})
	var tc4, _ = New(&Config{
		ISS:       "ISS",
		KID:       "KID",
		ExpiresIn: 1 * time.Minute,
		AUD:       "AUD",
		PK:        privateKeyNotECSDA,
	})

	tests := []struct {
		name   string
		cli    *Client
		expErr error
	}{
		{"Valid", tc1, nil},
		{"PrivateKeyNotValidPEM", tc2, ErrPrivateKeyNotValidPEM},
		{"PrivateKeyNotValidPKCS8", tc3, ErrPrivateKeyNotValidPKCS8},
		{"PrivateKeyNotECDSA", tc4, ErrPrivateKeyNotECDSA},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.cli.BearerToken()
			if err != tt.expErr {
				t.Errorf(unexpectedResult, err, tt.expErr)
				return
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	t.Log("Todo")
}

func TestTokenReuse(t *testing.T) {
	t.Log("Todo")
}
