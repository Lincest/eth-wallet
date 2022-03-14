package utils

/**
    utils
    @author: roccoshi
    @desc: session的处理逻辑
**/

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SessionData represents the session.
type SessionData struct {
	UID       uint   // user ID
	UName     string // username
	UPassword string // user password
	NetworkID uint   // current net-work id
}

// Save saves the current session of the specified context.
// 根据当前的context保存相关的session
func (sd *SessionData) Save(c *gin.Context) error {
	session := sessions.Default(c)
	sessionDataBytes, err := json.Marshal(sd)
	if nil != err {
		return err
	}
	session.Set("data", string(sessionDataBytes))

	return session.Save()
}

// GetSession returns session of the specified context.
// 返回当前context对应的session
func GetSession(c *gin.Context) *SessionData {
	ret := &SessionData{}

	session := sessions.Default(c)
	sessionDataStr := session.Get("data")
	if nil == sessionDataStr {
		return ret
	}

	err := json.Unmarshal([]byte(sessionDataStr.(string)), ret)
	if nil != err {
		return ret // 此时ret为空的sessionData, 可以使用UID == 0鉴别
	}

	c.Set("session", ret)

	return ret
}
