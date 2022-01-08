package middlewares

import (
	"account-go/common"
	"account-go/model"
	"account-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 方行登录接口
		if c.Request.URL.Path == "/api/login" {
			c.Next()
		} else {
			// 获取 Authorization header
			tokenString := c.GetHeader("Authorization")

			// 验证格式
			if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
				util.Response(c, http.StatusUnprocessableEntity, 401, nil, "权限不足")
				// Abort 防止挂起的处理程序被调用。请注意，这不会停止当前处理程序。
				// 假设您有一个授权中间件来验证当前请求是否已获得授权。如果授权失败（例如：密码不匹配），
				//则调用 Abort 以确保不会调用此请求的其余处理程序。
				c.Abort()
				return
			}

			tokenString = tokenString[7:]

			token, claims, err := common.ParseToken(tokenString)

			if err != nil || !token.Valid {
				util.Response(c, http.StatusUnprocessableEntity, 401, nil, "权限不足")
				c.Abort()
				return
			}

			// 通过验证后获取 token 中的 userID
			userId := claims.UserId
			db := util.DB
			var user model.User
			err = db.Model(model.User{}).First(&user, userId).Error
			if err != nil {
				util.Response(c, http.StatusUnprocessableEntity, 401, nil, err.Error())
				c.Abort()
				return
			}

			if user.Id == 0 {
				util.Response(c, http.StatusUnprocessableEntity, 401, nil, "权限不足")
				c.Abort()
				return
			}

			// 用户信息写入上下文
			c.Set("user", user)
			c.Next()
		}
	}
}
