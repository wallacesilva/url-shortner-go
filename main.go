package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacesilva/url-shortnet-go/core"
)

type FormShortUrl struct {
	Url string `form:"url"`
}

type UrlShorten struct {
	Code string `uri:"code" binding:"required"`
}

func main() {
	projectTitle := "URL Shortner using GO"

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": projectTitle,
		})
	})

	r.POST("/short", func(c *gin.Context) {
		var form FormShortUrl
		c.Bind(&form)
		c.HTML(http.StatusOK, "short.tmpl", gin.H{
			"title":        projectTitle,
			"url_short":    core.GetBaseUrl("/xpto"),
			"url_original": form.Url,
		})
	})

	r.GET("/:code", func(c *gin.Context) {
		var urlShorten UrlShorten
		if err := c.ShouldBindUri(&urlShorten); err != nil {
			c.JSON(400, gin.H{"message_error": err})
			return
		}

		// TODO: redirect to URL unshorten
		c.JSON(200, gin.H{"code": core.GetUrlByCode(urlShorten.Code)})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
