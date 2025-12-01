// middleware/oper_log.go
package middleware

import (
	"fmt"
	"ginManager/logger"
	"ginManager/models/entity"
	"ginManager/service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserLogMiddleware 记录操作日志（跳过静态资源 & 登录接口）
func UserLogMiddleware(svc *service.UserLogService) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 跳过登录 & 静态
		if path == "/api/auth/login" || path == "/favicon.ico" {
			c.Next()
			return
		}

		// 继续处理
		c.Next()

		// 事后记录
		duration := time.Since(start).Milliseconds()
		userID, _ := c.Get("userID") // JWT 中间件写入
		var uid *uint64
		if u, ok := userID.(uint64); ok && u > 0 {
			uid = &u
		}
		userName, _ := c.Get("userName")
		var username *string
		if un, ok := userName.(string); ok && un != "" {
			username = &un
		}
		durationMs := int(duration)
		module := extractModule(path)
		ip := c.ClientIP()
		userAgent := c.Request.UserAgent()
		errorMsg := ""
		fmt.Println(userID, userName, start, path, method, ip, userAgent, durationMs)
		log := entity.UserLog{
			UserID:     uid,
			Username:   username, // 可再查表补
			Module:     &module,  // 简单截取
			Action:     &method,
			Method:     &method,
			Path:       &path,
			IP:         &ip,
			UserAgent:  &userAgent,
			Status:     1,
			ErrorMsg:   &errorMsg,
			DurationMs: &durationMs,
			CreatedAt:  start,
		}
		// 状态码 >=400 认为是失败
		if c.Writer.Status() >= 400 {
			log.Status = 0
			errStr := c.Errors.ByType(gin.ErrorTypePrivate).String()
			log.ErrorMsg = &errStr
		}
		if err := svc.Create(c.Request.Context(), &log); err != nil {
			logger.L.Error("user log create fail", zap.Error(err))
		}
	}
}

// 简单截取模块名 /api/user/list -> user
func extractModule(path string) string {
	if len(path) < 5 {
		return "other"
	}
	segs := strings.Split(path[1:], "/")
	if len(segs) < 2 {
		return "other"
	}
	return segs[1]
}
