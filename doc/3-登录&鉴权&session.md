## 登录

系统第一个问题就是登录鉴权的问题

使用cookie + session的方式进行鉴权

- session中间件: [gin-contrib/sessions: Gin middleware for session management (github.com)](https://github.com/gin-contrib/sessions)

### session data

首先指定session的data存储包括:

```go
// SessionData represents the session.
type SessionData struct {
	UID   uint64 // user ID
	UName string // username
}
```

### session save

```go
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
```

### session get

```go
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
		return ret
	}

	c.Set("session", ret)

	return ret
}
```

