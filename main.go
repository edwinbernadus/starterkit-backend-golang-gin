package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"starterkit/models"
	"starterkit/module_db"
	"starterkit/module_socket"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//get
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	// /student/2
	//routing_id
	r.GET("/student/:id", func(c *gin.Context) {
		var id = c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "inquiry-" + id,
		})
	})

	//post_body_to_model
	r.POST("/student", func(c *gin.Context) {

		// {
		// 	"name" : "one",
		// 	"age" : 20
		// }
		var json models.Student
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		age := strconv.Itoa(json.Age)
		result := json.Name + " - " + age

		c.JSON(http.StatusOK, gin.H{"status": result})
	})

	//post
	r.POST("/student2", func(c *gin.Context) {

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		jsonStr := string(body)
		var student models.Student
		//json_convert
		err2 := json.Unmarshal([]byte(jsonStr), &student)
		if err2 != nil {
			fmt.Println(err2)
			// handle error
		}

		fmt.Println(string(body))

		c.JSON(http.StatusOK, gin.H{"status": student.Age})
	})

	//get_header
	r.GET("/test-header", func(c *gin.Context) {
		var d = c.Request.Header.Get("Authorization")
		c.String(http.StatusOK, d)
	})

	r.GET("/ws", func(c *gin.Context) {
		module_socket.RegisterWebSocket(c)
	})

	r.GET("/db/get_list", func(c *gin.Context) {
		db := module_db.GetDb()
		var students []*models.Student
		//sql_select_list
		db.Find(&students)

		str, err := json.Marshal(students)
		fmt.Println("get list")
		fmt.Println(string(str))
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": students})

	})

	r.GET("/db/get_one/:id", func(c *gin.Context) {
		var id = c.Param("id")

		var student *models.Student
		db := module_db.GetDb()
		//sql_select_get_one
		db.First(&student, id) // find product with integer primary key

		c.JSON(http.StatusOK, gin.H{
			"message": "inquiry-" + student.Name,
		})
	})

	r.POST("/db/create", func(c *gin.Context) {
		var db = module_db.GetDb()

		var student models.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//sql_create
		db.Create(&student)

		c.JSON(http.StatusOK, gin.H{
			"message": "create-item",
		})
	})

	r.POST("/db/update/:id", func(c *gin.Context) {
		var id = c.Param("id")
		var student *models.Student
		db := module_db.GetDb()

		db.First(&student, id) // find product with integer primary key

		var studentInput models.Student
		if err := c.ShouldBindJSON(&studentInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//sql_update
		student.Name = studentInput.Name
		student.Age = studentInput.Age

		db.Save(&student)

		c.JSON(http.StatusOK, gin.H{
			"message": "update",
		})
	})

	r.GET("/db/delete/:id", func(c *gin.Context) {
		var id = c.Param("id")
		var student *models.Student
		db := module_db.GetDb()

		db.First(&student, id) // find product with integer primary key
		//sql_delete
		db.Delete(&student)

		c.JSON(http.StatusOK, gin.H{
			"message": "delete",
		})
	})
	r.GET("/db/counter", func(c *gin.Context) {
		var db = module_db.GetDb()
		count := int64(0)
		//sql_count
		db.Model(&models.Student{}).Count(&count)
		c.JSON(http.StatusOK, gin.H{
			"total": count,
		})

	})

	//web_socket
	r.Run()
}
