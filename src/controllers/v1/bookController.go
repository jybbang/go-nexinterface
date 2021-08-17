package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jybbang/go-core-architecture/core"
	"github.com/jybbang/nexinterface/src/books/commands"
	"github.com/jybbang/nexinterface/src/books/queries"
)

func AddBookController(r *gin.RouterGroup) {
	books := r.Group("/books")
	{
		books.GET("/", getBooks)
		books.GET("/:id", getBook)
		books.POST("/", createBook)
		books.DELETE("/:id", deleteBook)
	}
}

func getBooks(c *gin.Context) {
	cmd := &queries.GetBooksQuery{}

	result := core.GetMediator().Send(c, cmd)

	c.JSON(result.ToHttpStatus(), result.V)
}

func getBook(c *gin.Context) {
	cmd := &queries.GetBookQuery{
		Id: c.Param("id"),
	}

	result := core.GetMediator().Send(c, cmd)

	c.JSON(result.ToHttpStatus(), result.V)
}

func createBook(c *gin.Context) {
	cmd := new(commands.CreateBookCommand)
	c.BindJSON(cmd)

	result := core.GetMediator().Send(c, cmd)

	c.JSON(result.ToHttpStatus(), result.V)
}

func deleteBook(c *gin.Context) {
	cmd := &commands.DeleteBookCommand{
		Id: c.Param("id"),
	}

	result := core.GetMediator().Send(c, cmd)

	c.JSON(result.ToHttpStatus(), result.V)
}
