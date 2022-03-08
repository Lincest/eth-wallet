## 登录

系统第一个问题就是登录鉴权的问题

使用cookie + session的方式进行鉴权

- session中间件: [gin-contrib/sessions: Gin middleware for session management (github.com)](https://github.com/gin-contrib/sessions)

### session init

```go
// session
// store := cookie.NewStore([]byte(utils.Rand.String(16))) // use 16 random string as secret of session, 这样会导致每次服务器重启之前的用户session失效
store := cookie.NewStore([]byte("skqiswkdjcaqwedj")) // 可以手动指定一串secret防止重启服务器cookie失效, 但是这种secret不应该出现在源代码中, 这里为了简化就直接把密钥写死
router.Use(sessions.Sessions("eth-wallet-session", store))
```

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
		return ret // 错误处理也是直接返回, 但是可以通过UID == 0判断session不存在
	}

	c.Set("session", ret)

	return ret
}
```

## 鉴权

### login时设置session

- 首先用户访问`api/v1/login`进行登录, 登录时以用户id和username设置session的data

  ```go
  func LoginAction(c *gin.Context) {
  	resp := utils.NewBasicResp()
  	defer c.JSON(http.StatusOK, resp)
  	name := c.PostForm("name")
  	password := c.PostForm("password")
  	// ...省略数据库逻辑
      // session核心
  	session := &utils.SessionData{UID: user.ID, UName: user.Name}
  	if err := session.Save(c); nil != err {
  		c.Status(http.StatusInternalServerError)
  		return
  	}
  	resp.Msg = "登陆成功"
  }
  ```

- 此时可以通过postman查看服务器的返回值就包含了cookie: 

  ![](https://youpai.roccoshi.top/img/202203081938483.png)

### LoginCheck中间件

用于检查当前cookie是否正确, 如果当前session不存在会导致uid为默认值(`0`)

```go
// 中间件 authmiddleware.go
// LoginCheck 检查用户登录状态
func LoginCheck(c *gin.Context) {
	session := utils.GetSession(c)
	if 0 == session.UID {
		result := utils.NewBasicResp()
		result.Code = model.CodeErr
		result.Msg = "unauthenticated request"
		c.AbortWithStatusJSON(http.StatusOK, result)
		return
	}
	c.Next()
}
```

同时在`router`中包含这个check中间件并且建立一个新的auth用于授权的api, 所以目前的api group情况如下: 

```go
// 不经过鉴权可以访问的api, 如登录, 注册等
v1 := router.Group("api/v1")
{
    v1.POST("/register", ctl.RegisterAction)
    v1.POST("/login", ctl.LoginAction)
    // ...
}
// 经过鉴权的api
authGroup := v1.Group("auth")
authGroup.Use(ctl.LoginCheck) // auth中间件检查用户是否登录
{
    authGroup.GET("/hello-world/:user", ctl.HelloWorldAction)
    // ...
}
```

