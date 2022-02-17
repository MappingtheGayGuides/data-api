package main

import (
	"github.com/regan008/PersonWeb/models"
	//"fmt"
	"log"
	"net/http"
	//"strconv"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()
	r.Use(CORSMiddleware())
	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("person", getPersons)
		v1.GET("person/:id", getPersonById)
		// v1.POST("person", addPerson)
		// v1.PUT("person/:id", updatePerson)
		// v1.DELETE("person/:id", deletePerson)
		// v1.OPTIONS("person", options)
	}
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(":1313")
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "*")


        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func getPersons(c *gin.Context) {

	persons, err := models.GetPersons(100)

	checkErr(err)

	if persons == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": persons})
	}
}
//call get person by id function from person.go
func getPersonById(c *gin.Context) {

	// grab the Id of the record want to retrieve
	id := c.Param("id")

	person, err := models.GetPersonById(id)

	checkErr(err)
	// if the name is blank we can assume nothing is found
	if person.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": person})
	}
}
//
// func addPerson(c *gin.Context) {
//
// 	var json models.Person
//
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	success, err := models.AddPerson(json)
//
// 	if success {
// 		c.JSON(http.StatusOK, gin.H{"message": "Success"})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 	}
// }
//
// func updatePerson(c *gin.Context) {
//
// 	var json models.Person
//
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	personId, err := strconv.Atoi(c.Param("id"))
//
// 	fmt.Printf("Updating id %d", personId)
//
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 	}
//
// 	success, err := models.UpdatePerson(json, personId)
//
// 	if success {
// 		c.JSON(http.StatusOK, gin.H{"message": "Success"})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 	}
// }
//
// func deletePerson(c *gin.Context) {
//
// 	personId, err := strconv.Atoi(c.Param("id"))
//
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 	}
//
// 	success, err := models.DeletePerson(personId)
//
// 	if success {
// 		c.JSON(http.StatusOK, gin.H{"message": "Success"})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 	}
// }

// func options(c *gin.Context) {
//
// 	ourOptions := "HTTP/1.1 200 OK\n" +
// 		"Allow: GET,POST,PUT,DELETE,OPTIONS\n" +
// 		"Access-Control-Allow-Origin: http://locahost:8080\n" +
// 		"Access-Control-Allow-Methods: GET,POST,PUT,DELETE,OPTIONS\n" +
// 		"Access-Control-Allow-Headers: Content-Type\n"
//
// 	c.String(200, ourOptions)
// }

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
