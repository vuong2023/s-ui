package api

import (
	"encoding/gob"
	"s-ui/database/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	loginUser = "LOGIN_USER"
)

func init() {
	gob.Register(model.User{})
}

func SetLoginUser(c *gin.Context, userName string) error {
	s := sessions.Default(c)
	s.Set(loginUser, userName)
	return s.Save()
}

func SetMaxAge(c *gin.Context, maxAge int) error {
	s := sessions.Default(c)
	s.Options(sessions.Options{
		Path:   "/",
		MaxAge: maxAge,
	})
	return s.Save()
}

func GetLoginUser(c *gin.Context) string {
	s := sessions.Default(c)
	obj := s.Get(loginUser)
	if obj == nil {
		return ""
	}
	objStr, ok := obj.(string)
	if !ok {
		return ""
	}
	return objStr
}

func IsLogin(c *gin.Context) bool {
	return GetLoginUser(c) != ""
}

func ClearSession(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	s.Save()
}
