package jwt

import "github.com/golang-jwt/jwt/v5"

// JwtDecode ใช้ jwt.ParseUnverified เพื่อดึง claims (ไม่ verify signature)
func JwtDecode(raw string) (jwt.MapClaims, error) {
	// ParseUnverified(tokenString, destinationClaims)
	claims := jwt.MapClaims{}
	_, _, err := new(jwt.Parser).ParseUnverified(raw, claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
