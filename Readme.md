main 中目前使用的结构方法来实现多个不同的接口的相同方法的调用

```bash
GET请求
    /:id/:name
代码中: 
    id := c.Param("id")
```

```bash
POST请求
    第一种post form:
    header(Content-TypeL: application/x-www-form-urlencoded)
    body(form) cid: 10; name: aa
代码中: cid := c.PostForm("cid")
    第二种post json(1、map方式 2、接收结构体):
    header(Content-TypeL: application/json)
    body({"cid": 10, "name": "abc"})
        获取提交参数的方式:
        1、map方式
            param := make(map[string]interface{})
            err := c.BindJSON(&param)
        2、结构体方式
        定义一个结构跟传入的json数据一致, search := &Search{}, err := c.BindJSON(&search)
```

```bash
GORM
  	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	dao    层用于数据库连接的基础配置如连接: Db, err = gorm.Open("mysql", config.MysqlDB)
	models 层用于数据的具体操作 insert delete update select
```