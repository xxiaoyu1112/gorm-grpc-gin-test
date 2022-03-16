package Handler

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"time"
//)
//
//func CreateHandler(c *gin.Context) {
//	// 1.从请求中把数据拿出来
//	var todo Todo
//	c.BindJSON(&todo)
//	// 2.存入数据库
//	// 3.返回响应
//
//	num := 1
//	title := &Todo{
//		Item: num + 1,
//		Time: time.Now(),
//	}
//	if err := DB.Create(title).Error; err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"error": err.Error(),
//		})
//	} else {
//		//c.JSON(http.StatusOK, todo)
//		c.JSON(http.StatusOK, gin.H{
//			"code": 2000,
//			"msg":  "success",
//			"data": todo,
//		})
//	}
//}
//
//func FindHandler(c *gin.Context) {
//	var todoList []Todo
//	if err := DB.Find(&todoList).Error; err != nil {
//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 2000,
//			"msg":  "success",
//			"data": todoList,
//		})
//	}
//}
//
//func ReviseHandler(c *gin.Context) {
//	id, ok := c.Params.Get("id")
//	if !ok {
//		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
//		return
//	}
//	var todo Todo
//	if err := DB.Where("id = ? ", id).First(&todo).Error; err != nil {
//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//	}
//	c.BindJSON(&todo)
//	if err := DB.Updates(&todo).Error; err != nil {
//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 2000,
//			"msg":  "success",
//			"data": todo,
//		})
//	}
//}
//func DeleteHandler(c *gin.Context) {
//	id, ok := c.Params.Get("id")
//	if !ok {
//		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
//		return
//	}
//	if err := DB.Where("id = ? ", id).Delete(Todo{}).Error; err != nil {
//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 2000,
//			"msg":  "success",
//			id:     "delete",
//		})
//	}
//}
