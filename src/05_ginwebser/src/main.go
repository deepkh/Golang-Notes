package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
  _"os"
  _"io/ioutil"
  _"errors"
  "io"
  "github.com/tpkeeper/gin-dump"
  "time"
)

func UserNameGet(c *gin.Context) {
  name := c.Param("name")
  action := c.Param("action")
  message := name + " is " + action
  if c.FullPath() == "/user/:name/*action" {
    message += "  yes"
  }
  c.String(http.StatusOK, message)
}

func QueryGet(c *gin.Context) {
  firstname := c.DefaultQuery("firstname", "Guest")
  lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
  c.String(http.StatusOK, "firstname=%s lastname=%s", firstname, lastname)
}

/*
  POST /form_urlencoded?pa=&pb= HTTP/1.1
  Content-Type: application/x-www-form-urlencoded

  name=manu&message=this_is_great
*/
func FormUrlEncodedPost(c *gin.Context) {
  name := c.DefaultPostForm("name", "anonymous")
  message := c.PostForm("message")
  pa := c.Query("pa")
  pb := c.DefaultQuery("pb", "default_pb")

  c.JSON(200, gin.H{
    "status":  "posted",
    "message": message,
    "name":    name,
    "pa": pa,
    "pb": pb,
  })
}

func JsonPost(c *gin.Context) {
  // Binding from JSON
  type Login struct {
    User     string `form:"user" json:"user" xml:"user"  binding:"required"`
    Password string `form:"password" json:"password" xml:"password" binding:"required"`
  }

  // Example for binding JSON ({"user": "manu", "password": "123"})
  var json Login
  if err := c.ShouldBindJSON(&json); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
/*
  if json.User != "manu" || json.Password != "123" {
    c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
    return
  }
*/
  c.JSON(http.StatusOK, gin.H{"status": "json.User="+json.User+" json.Password="+json.Password})
}

type CustomReader struct {
  remaining_size int64
  total_size int64
}

func NewCustomReader(total_size int64) *CustomReader {
  return &CustomReader{
    remaining_size: total_size,
    total_size: total_size,
  }
}

func (p *CustomReader) Read(buf []byte) (n int, err error) {
  n = 0
  if p.remaining_size > 0 {
    if p.remaining_size < 4096 {
      n = int(p.remaining_size)
    } else {
      n = 4096
    }

    for i:=0; i<n; i++ {
      buf[i] = byte('A' + i%3)
    }

    p.remaining_size -= int64(n)
    fmt.Printf("read %v/%v/%v/%v\n", n, len(buf), p.remaining_size, p.total_size)
    return n, nil
  } else {
    return n, io.EOF
  }
}

func ReaderGet(c *gin.Context) {
  var content_length int64 = 1000000
  reader := NewCustomReader(content_length)
  /*extraHeaders := map[string]string{
    "Content-Disposition": `attachment; filename="1.htm"`,
  }*/
  c.DataFromReader(http.StatusOK, content_length, "text/html", reader, nil/*extraHeaders*/)
}

func ServerRender(c *gin.Context) {
  c.HTML(http.StatusOK, "server_render.tmpl", gin.H{
    "title": "Hello World!",
  })
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
  session_key, err := c.Cookie("session_key")
  if err != nil {
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized: session_key empty"})
    return
  }

  if session_key != "123456" {
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized: invalid session_key"})
    return
  }

  // Continue down the chain to handler etc
  c.Next()
}

func BindQueryString(c *gin.Context) {
  //look like only "string" can binding currently
  //c.Bind for query string and post data
  //test1: Post by using 'content-type: application/x-www-form-urlencoded'
  //        curl -X POST -d "v1=vvvv1&v2=vvv2&v3=1234" http://127.0.0.1:8080/bind_query_string -k -v
  //test2: Get by using
  //        curl -X GET "http://127.0.0.1:8080/bind_query_string?v1=vvvv1&v2=vvv2&v3=1234" -k -v
  //test3: Post by using  'content-type: application/json'
  //        curl -X POST -d  '{"v1":"vvvvv1","v2":"vvvvv2","v3":1234}' -H 'content-type: application/json' http://127.0.0.1:8080/bind_query_string -k -v
  type query struct {
    V1 string   `form:"v1"`
    V2 string   `form:"v2"`
    V3 int   `form:"v3"`
  }

  var q query

  // If `GET`, only `Form` binding engine (`query`) used.
  // If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
  // See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
  if c.ShouldBind(&q) == nil {
    c.String(200, "bind ok v1:%s v2:%s v3:%d\n", q.V1, q.V2, q.V3)
  } else {
    c.String(200, "bind not ok v1:%s v2:%s v3:%d\n", q.V1, q.V2, q.V3)
  }
}


func main() {
  var router *gin.Engine = gin.Default()

  //middleware: dump request, response 
  router.Use(gindump.Dump())

  //middleware: custom log format 
  router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
    // your custom format
    return fmt.Sprintf("%s xx - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
        param.ClientIP,
        param.TimeStamp.Format(time.RFC1123),
        param.Method,
        param.Path,
        param.Request.Proto,
        param.StatusCode,
        param.Latency,
        param.Request.UserAgent(),
        param.ErrorMessage,
    )
  }))

  router.GET("/user/:name/*action", UserNameGet)
  router.GET("/query", QueryGet)
  router.POST("/post_form_urlencoded", FormUrlEncodedPost)
  router.POST("/post_json", JsonPost)

  //StaticFS provide ability to list file and list direcoty and also file download ability
  router.StaticFS("/test_file", http.Dir("test_file"))

  // write data from custom reader
  router.GET("/reader", ReaderGet)

  // html render by using template
  router.LoadHTMLGlob("src/05_ginwebser/src/templates/*")  
  router.GET("/render", ServerRender)

  // group: /v1/ping 
  v1 := router.Group("/v1")
  v1.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
  });

  // login: set session_key Cookie's value to 123456
  router.GET("/login", func(c *gin.Context) {
    c.SetCookie("session_key", "123456", 3600, "/", "192.168.4.22", false, true)
    c.String(200, "login ok\n")
  });

  // logout: clear the session_key Cookie's value
  router.GET("/logout", func(c *gin.Context) {
    c.SetCookie("session_key", "", 3600, "/", "192.168.4.22", false, true)
    c.String(200, "logout ok\n")
  });

  // middleware: authed private path
  private := router.Group("/")
  private.Use(AuthRequired)
  {
    //curl -k -v --cookie "session_key=123456" http://127.0.0.1:8080/me
    private.GET("/me", func(c *gin.Context) {
      session_key, _ := c.Cookie("session_key")
      c.String(200, "My Authed Session Key is: " +session_key+"\n")
    });
  }

  // bind query string to struct
  router.POST("/bind_query_string", BindQueryString);
  router.GET("/bind_query_string", BindQueryString);

  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
