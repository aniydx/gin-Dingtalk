main 中目前使用的结构方法来实现多个不同的接口的相同方法的调用

GET请求
    /:id/:name
代码中: 
    id := c.Param("id")

POST请求
    第一种post form:
    header(Content-TypeL: application/x-www-form-urlencoded)
    body(form) cid: 10; name: aa
代码中: cid := c.PostForm("cid")
    第二种post json(1、map方式 2、接收结构体):
    header(Content-TypeL: application/json)
