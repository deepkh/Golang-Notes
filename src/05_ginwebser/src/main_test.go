package main

import (
  "fmt"
  "testing"
  "github.com/go-resty/resty/v2"
  "encoding/json"
)

const (
  default_url = "http://127.0.0.1:8080"
)

func Dump(resp *resty.Response, err error) {
  // Explore response object

  if err != nil {
    fmt.Printf("  Error      :%v\n", err)
    return
  }
  fmt.Printf("%-9v %-3s %-80s %v \"%v\"\n", resp.Time(), resp.Request.Method, resp.Request.URL, resp.StatusCode(), resp)
}

func Get(uri string) {
  resp, err := resty.New().R().Get(default_url + uri)
  Dump(resp, err)
}

func Gets(uris []string) {
  for _, e := range uris {
    Get(e)
  }
}

func TestUser(t *testing.T) {
  uris := []string{
      "/user/abc",
      "/user/defg/",
      "/user/hijkl/abc_action",
      "/user/mnopqr/abcd_action/",
  }

  Gets(uris)
}

func TestQuery(t *testing.T) {
  uris := []string{
      "/query?firstname=Carlos&lastname=Yang",
      "/query/?firstname=&lastname=BbBBb",
      "/query/?lastname=CCCCCCCC",
  }

  Gets(uris)
}

// Post by use -> "Content-Type":  "application/x-www-form-urlencoded"
//func Post(uri string, headers *map[string]string, form_data *map[string]string, body interface{}) {
func Post(uri string, headers interface{}, form_data interface{}, body interface{}) {
  cli := resty.New().R();

  if headers != nil {
    cli.SetHeaders(headers.(map[string]string))
  }

  if form_data != nil {
    cli.SetFormData(form_data.(map[string]string))
  }

  if body != nil {
    cli.SetBody(body)
  }

  resp, err := cli.Post(default_url+uri)
  Dump(resp, err)
}

// Post by use -> "Content-Type":  "application/x-www-form-urlencoded"
func TestFormPost(t *testing.T) {
  Post("/post_form_urlencoded?pa=123&pb=456",
    nil,
    map[string]string{
      "name": "Rick",
      "message": "Martin",
    },
    nil,
  );

  Post("/post_form_urlencoded",
    nil,
    map[string]string{
    },
    nil,
  );
}

func StructToString(obj interface{}) string {
    b, err := json.Marshal(obj)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return "failed to convert StructToString";
    }
    return string(b)
}

// Post by use -> "Content-Type":  "application/json"
func TestPostJson(t *testing.T) {
  // Binding from JSON
  type Login struct {
    User     string `form:"user" json:"user" xml:"user"  binding:"required"`
    Password string `form:"password" json:"password" xml:"password" binding:"required"`
  }

  Post("/post_json",
    map[string]string{
      "Content-Type": "application/json",
    },
    nil,
    StructToString(&Login{User:"XXXXXXXX", Password: "GGGGGGGG"}),
  );
}
