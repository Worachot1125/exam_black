package login

import (
	"app/app/request"
	"app/app/response"
	"app/app/util/jwt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

const tokenTTL = 10 * time.Minute // อายุ token

// ----------- Controller -------------

// Login: ตรวจ user → ออก token → ส่งกลับ
func (ctl *Controller) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid input")
		return
	}

	user, err := ctl.Service.Login(ctx, req.User_Number, req.Password)
	if err != nil {
		response.Unauthorized(ctx, err.Error())
		return
	}

	exp := time.Now().Add(tokenTTL).Unix()
	claims := jwt5.MapClaims{
		"user_id":     user.ID,
		"user_number": user.User_Number,
		"first_name":  user.FirstName,
		"last_name":   user.LastName,
		"exp":         exp,
	}

	token, err := jwt.CreateToken(claims, viper.GetString("JWT_SECRET_USER"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	res := response.LoginResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		UserNumber: user.User_Number,
		Role_ID:    user.Role_ID,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  res,
		"token": token,
		"exp":   exp,
	})
}

// Debug: ถอดรหัส token + ตรวจ exp (ใช้ Header หรือ ?token=)
func (ctl *Controller) Debug(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		tokenStr = ctx.Query("token")
	}
	if tokenStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	claims, err := jwt.JwtDecode(tokenStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token: " + err.Error()})
		return
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(expFloat) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"claims": claims})
}

// VerifyToken: สำหรับ frontend เช็ก token ว่ายังใช้ได้
func (ctl *Controller) VerifyToken(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		tokenStr = ctx.Query("token")
	}
	if tokenStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	claims, err := jwt.JwtDecode(tokenStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
		return
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok || time.Now().Unix() > int64(expFloat) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Token valid",
		"claims":  claims,
	})
}

// ----------- Middleware -------------

// AuthMiddleware: เช็ก Header Authorization: Bearer <token> เท่านั้น
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Authorization header must be Bearer <token>"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok || userID == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "invalid token: user_id not found"})
			return
		}

		ctx.Set("user_id", userID)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
