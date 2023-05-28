package test

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

var priKeyStr = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDNRsMvfxHzqLeG
m0XAXc1R3Mz03UH7f5zmQssv5jkMpXxyuALWNYpD2jgQ2jgClJUWnR/AiGIHAarM
yAWyACQJagREEcIct0zVP67aF4dxdwULp633YzQW/P8zUXzIYnF2/LAZOXMTPNVu
OtDJJx6ylPotL1eHLtLlrm6Buauv73lCDWAPW7Ck1s5vzIDPR2bEIYVPSl2V0t5V
Vo2UzHt1FQ8Mq2m87C3hK0X9DzTyjxivEQvjz2RnmJL1B1aYc2uBCP+Po2BWk5st
mAhCVE6KgN+Kzzw7nq6HSAwzQMN2q6wWIX8dEEkysN+i3r1DWYVdmAHe8xXzoZC2
0ljAiTx9AgMBAAECggEAOjJZm9mWqVGn1k2nFMaDMzY6FSG1Vyvyu4UsPcQzYFZ6
Vbu9ciVzsZDoXD7Bqlmdeu/I1LjVsc7TMYNzuec4UELyOqoeMi31Zm/LjqElERIB
KDC9rWk+l1XatB3iFp1yNZ2l/0C+UzHoAHxEPQMOPOrnkm/djMHFoMA3LCgLeLEV
3vHKkeCt8de6F7+KGamso0v6Iadvn/lHI26IN+BJ0qRkDDI6HUoJqbnZJxgTluAO
dwAXy3jWimetuBeOIT4wNXUDiYdWqYdXs5FRPe8memnk7o+atDPfJ+twS2XkSyDQ
zLWr+v1pLj08gocLUutLfFxVwTKthYSmj24Y7ht6sQKBgQDnHDQ9uJ1lRAxsEkGf
KBvnz0hb8xhtXdF7QWIXcoIBuCHJ3c718/SPhhuw//YA6XlJ3piTBBrq2gsOJiDP
xvE+cekxEbd9u53hVJ2fmrN5bvkTTOVhvaWq4wNg11hFtvjcU+OzG5ehffo+dHzu
x/b0bkSPj52nP+SsgykPPCFl7wKBgQDjYk/TpFpSZdaCjwLygENiHnM85x0Ia+jR
n/lfEEMRJfcXUihBh1NXUXpZ8xsmffwbUv+EAz6BVZUfZ0U8PU4hGrx6r1+Qlc6G
Dmx7SxaMVjh/yzdFFVUsCtpaeBiN0+jg4L1t3RPZ3RP0BPF18YLiZYu+Le3kh/K8
Wp376k3QUwKBgQCk+dY1BReeVVBEyVd8xMX+2VI/CS6zy5ghU1AxirVRgt7j2mnF
2ysGVWZpGJ7EkeXaHINv1ytb4OCpbgBYMhy+RdSACbShlY+jbaLDb0yU7+nvpCHO
fvHHJhygQbkqsu29YkkV7ylzx5kegks4rRgV7q0UiiGxZYPYvhxOWs9AkwKBgQCn
LUjzmgquOiGUrADGunbQVQL07Bb0ciIivTNjKVml6fvZMZZXV192+3ixWYPEsSwC
CuvB64CxJnMVO6Azwf8HZ9jbesUQUJQfC8vGelaBp4Kysn5YVG7iirgve8zRudOm
QpYDiF9n9psM9fVxebd5LJ+pm6skMq/Mu/MbnBDJ8QKBgBxUZdW5Qxkav3YenhP/
kkL86UHDqlcxrtBsJGSgi5Yo3tbiGpYO2HBGF+6YZIrjoqkcqVEUMDsfwPivvhXz
yNSRK986CTegDJ4NZZ+YJdj3yuwXW/Uk4tKPU2A+UHMYKtGUuwDC+rq7Qb0qrjFc
jx4gSFcpFKVifGCd7iAv6/6+
-----END PRIVATE KEY-----`

var pubKeyStr = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzUbDL38R86i3hptFwF3N
UdzM9N1B+3+c5kLLL+Y5DKV8crgC1jWKQ9o4ENo4ApSVFp0fwIhiBwGqzMgFsgAk
CWoERBHCHLdM1T+u2heHcXcFC6et92M0Fvz/M1F8yGJxdvywGTlzEzzVbjrQySce
spT6LS9Xhy7S5a5ugbmrr+95Qg1gD1uwpNbOb8yAz0dmxCGFT0pdldLeVVaNlMx7
dRUPDKtpvOwt4StF/Q808o8YrxEL489kZ5iS9QdWmHNrgQj/j6NgVpObLZgIQlRO
ioDfis88O56uh0gMM0DDdqusFiF/HRBJMrDfot69Q1mFXZgB3vMV86GQttJYwIk8
fQIDAQAB
-----END PUBLIC KEY-----`

func TestJwt(t *testing.T) {
	priKeyBase64 := base64.StdEncoding.EncodeToString([]byte(pubKeyStr))
	fmt.Println(priKeyBase64)
	pk, err := base64.StdEncoding.DecodeString(priKeyBase64)
	assert.NoError(t, err)
	fmt.Println(string(pk))

	priKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(priKeyStr))
	assert.NoError(t, err)
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKeyStr))
	assert.NoError(t, err)

	sig, err := generateToken(priKey, 1001)
	assert.NoError(t, err)

	valid, err := validateToken(pubKey, sig)
	assert.NoError(t, err)
	assert.True(t, valid)
}

// 生成Token
func generateToken(priKey *rsa.PrivateKey, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"uid": userId,
	})
	return token.SignedString(priKey)
}

// 验证Token是否有效
func validateToken(pubKey *rsa.PublicKey, sigString string) (bool, error) {
	token, err := jwt.ParseWithClaims(sigString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确认签名方法符合要求
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok {
		fmt.Println(claims)
	}

	return true, nil
}
