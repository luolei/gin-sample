package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {

    fmt.Println("get tags")
}

func AddTag(c *gin.Context) {

    fmt.Println("add tag")
}

func EditTag(c *gin.Context) {

    fmt.Println("edit tag")
}

func DeleteTag(c *gin.Context) {

    fmt.Println("delete tag")
}
