package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

var DB *gorm.DB

type Todo struct {
	ID   int       `json:"id"`
	Item int       `json:"item"`
	Time time.Time `json:"time"`
}

func initMysql() (err error) {
	DSN := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", DSN)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	// 连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}

	// 程序退出，关闭连接
	defer DB.Close()

	// 模型绑定
	DB.AutoMigrate(&Todo{}) //  表名为todos

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1Group := r.Group("/things")
	{
		// 添加数据
		v1Group.POST("/todo", func(c *gin.Context) {
			// 1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 2.存入数据库
			// 3.返回响应

			num := 1
			title := &Todo{
				Item: num + 1,
				Time: time.Now(),
			}
			if err := DB.Create(title).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				//c.JSON(http.StatusOK, todo)
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		// 查找所有的
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todoList,
				})
			}
		})

		// 修改某一个
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			var todo Todo
			if err = DB.Where("id = ? ", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}
			c.ShouldBind(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			if err = DB.Where("id = ? ", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					id:     "delete",
				})
			}
		})
	}
	// 查看某一个
	r.GET("/todoone/:id", func(c *gin.Context) {
		fmt.Println("123456")
		id := c.Param("id")
		var todo Todo
		if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "success",
				"data": todo,
			})
		}
	})

	r.Run("0.0.0.0:9090")

}
