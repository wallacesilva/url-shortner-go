package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FormShortUrl struct {
	Url string `form:"url"`
}

type UrlShorten struct {
	Code string `uri:"code" binding:"required"`
}

// TODO: fix me please
// var baseUrl string = os.Getenv("BASE_URL")

func getBaseUrl(uri string) string {

	var url string
	var baseUrl string

	baseUrl = os.Getenv("BASE_URL")

	if baseUrl != "" {
		url = baseUrl + uri
	} else {
		url = "http://0.0.0.0:8080" + uri
	}
	return url
}

// TODO: Add tests
func getUrlByCode(code string) string {
	return getBaseUrl("/" + code)
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
			"url_short":    getBaseUrl("/xpto"),
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
		c.JSON(200, gin.H{"code": getUrlByCode(urlShorten.Code)})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
