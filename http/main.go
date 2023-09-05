package main

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)

func init() {
	_, err := sql.Open("postgres", "postgresql://cvzn:cvzninfra557@db-auth.cumrcppqpeht.ap-northeast-1.rds.amazonaws.com/cvzn_users")
	if err != nil {
		panic(err)
	}


}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run(":8080")
}
