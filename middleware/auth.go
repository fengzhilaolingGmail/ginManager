package middleware

import (
	"ginManager/logger"
	"ginManager/service"
	"ginManager/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewAuthMiddleware 返回 JWT+权限校验中间件
// perm 为空时只做登录态校验；否则校验具体权限
func NewAuthMiddleware(perm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 提取 Bearer
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "未提供token"})
			c.Abort()
			return
		}
		tokenStr := auth[7:]

		// 2. 解析 JWT
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "token无效或已过期"})
			c.Abort()
			return
		}

		// 3. 写入userID到上下文
		c.Set("userID", claims.UserID)

		// 4. 只需登录态 -> 放行
		if perm == "" {
			c.Next()
			return
		}

		// 5. 鉴权限
		svc := service.NewAuthService()
		codes, err := svc.GetPermissions(c.Request.Context(), claims.UserID)
		if err != nil {
			logger.L.Error("get permissions fail", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "系统异常"})
			c.Abort()
			return
		}
		// 支持 * 通配
		ok := false
		for _, v := range codes {
			if v == "*" || v == perm {
				ok = true
				break
			}
		}
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"code": 1, "msg": "无接口权限"})
			c.Abort()
			return
		}

		c.Next()
	}
}
