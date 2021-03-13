 create table users(
     id  int  primary key,
     username varchar(10),
     department varchar(50),
     password varchar(20),
    repassword varchar(20),package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

//Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Static bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}

	return DB.DB().Ping()
}

func main() {
	//创建数据库
	//sql:CREAT DATABASE BUBBLES ;
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close() //完整的关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&Todo{}) //todos

	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")

	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)

	})
	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写一个待办事项 点击提交  发请求到这里
			//1从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			//2存入数据库
			//err = DB.Create(&todo).Error
			//if err !=nil{
			//
			// }
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
				//c.JSON(http.StatusOK,gin.H{
				//	"code":2000,
				//	"msg":"success",
				//	"data":todo,
				//})
			}
			//3返回响应
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
				//...
			}
		})
		//查看所有的待办事项，
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询todo这个表里的所有数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

            id ,ok:=c.Params.Get("id")
            if !ok{
				c.JSON(http.StatusOK,gin.H{"error":"id不存在"})
				return
            }
            var todo Todo
            if err = DB.Where("id=?",id).First(&todo).Error;err !=nil{
            	c.JSON(http.StatusOK,gin.H{"error":err.Error()})
				return
			}
            c.BindJSON(&todo)
            if err = DB.Save(&todo).Error;err!=nil{
            	c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else{
				c.JSON(http.StatusOK,todo)
			}
		})
		//删除某一个代办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id ,ok:=c.Params.Get("id")
			if !ok{
				c.JSON(http.StatusOK,gin.H{"error":"id不存在"})
				return
			}
			if err = DB.Where("id=?",id).Delete(Todo{}).Error;err !=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else{
				c.JSON(http.StatusOK,gin.H{
					id :"deleted",
				})
			}
		})

	}

	r.Run()

}

    phone varchar(26),
    email varchar(255),
    created_at  varchar(255),
    update_at  varchar(255)
);
