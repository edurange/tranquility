package main

import (
	"os"
	"fmt"

	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis"

	"github.com/google/uuid"
)

var (
	queue = openQueue()
	realUUID = os.Args[1]
)

type MESSAGE struct{
	USER string `json:"user" binding:"required"`
	TIME float64 `json:"time" binding:"required"`
	COMMAND string `json:"command" binding:"required"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(simpleAuth())

	router.POST("/logger", sentMessage)

	router.GET("/results/:student", fetchLogs)

	router.Run(":8080")
}

func simpleAuth() gin.HandlerFunc {
	return func(c *gin.Context){
		clientID := c.Request.Header.Get("uuid")

		if(clientID == ""){
			authError(c, "missing secret")
			return
		}

		if(realUUID != clientID){
			authError(c, "incorrect secret")
		}

		c.Next()
	}
}

func authError(c *gin.Context, message interface{}){
	c.AbortWithStatusJSON(401, gin.H {"message": message})
}

func sentMessage(c *gin.Context){
	var message MESSAGE

	jErr := c.BindJSON(&message)
	if (jErr != nil){
		panic(jErr)
		c.JSON(http.StatusBadRequest, gin.H {})
		return
	}

	id, err := uuid.NewUUID() //unique woo
    if err !=nil {
        c.JSON(http.StatusBadRequest, gin.H {})
		return
    }

    fmt.Println(message)

	_, err2 := queue.ZAdd(message.USER, &redis.Z{message.TIME, message.COMMAND + " " + id.String()}).Result()
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H {})
		return
	}

	c.JSON(http.StatusOK, gin.H { //need to return OK status code and nothing else here....
		"message": "ok",
	})
}

func fetchLogs(c *gin.Context){
	student := c.Param("student")

	//get stuff from redis here 
	vals, err := queue.ZRangeWithScores(student, 0, -1).Result()
	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H {})
		return
	}

	c.JSON(http.StatusOK, gin.H { //do we need to searialize?
		student: vals,
	})
}

//opens redis connection 
func openQueue()(*redis.Client){					//just spitballing here...
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}