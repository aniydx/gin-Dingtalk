main 中目前使用的结构方法来实现多个不同的接口的相同方法的调用

```azure
GET请求
    /:id/:name
代码中: 
    id := c.Param("id")
```

```azure
POST请求
    第一种post form:
    header(Content-TypeL: application/x-www-form-urlencoded)
    body(form) cid: 10; name: aa
代码中: cid := c.PostForm("cid")
    第二种post json(1、map方式 2、接收结构体):
    header(Content-TypeL: application/json)
    body({"cid": 10, "name": "abc"})
        1、map方式
            param := make(map[string]interface{})
            err := c.BindJSON(&param)
        2、结构体方式
        定义一个结构跟传入的json数据一致, search := &Search{}, err := c.BindJSON(&search)
```
