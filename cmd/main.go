package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

const Port string = ":8080"

func main()  {
	r := gin.Default()
	r.GET("/" , func (c *gin.Context){
		c.JSON(200, gin.H{
			"message": "HelloWorld",
		})
	})

	log.Printf("Server Is Running On Port %v" , Port)
	if err := r.Run(Port); err != nil {
		log.Fatalf("Failed to run server %v", err)
	}
}