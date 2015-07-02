package application

import (
	"net/http"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/skibish/paffol/generator"
	"github.com/skibish/paffol/helpers"
)

// Run function starts the application and setup's everything
func Run() {
	router := gin.New()

	// Routes to public assets
	router.Static("images", "public/images")
	router.Static("css", "public/css")
	router.Static("js", "public/js")

	// Set renderer
	router.HTMLRender = renderer()

	// == Router configuration ==

	// Show index page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{})
	})

	// Show results page with generated password
	router.POST("/magic", func(c *gin.Context) {
		original := c.PostForm("phrase")
		g := generator.Generator{Input: original}
		password := g.Generate()

		c.HTML(http.StatusOK, "result", gin.H{
			"original": original,
			"password": password,
		})
	})

	// If /magic is refreshed, redirect to root
	router.GET("/magic", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Show about page
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about", gin.H{"activeAbout": true})
	})

	// Show tips page
	router.GET("/tips", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tips", gin.H{"activeTips": true})
	})

	// Start router
	router.Run(helpers.GetPort())
}

// renderer is a function, where all needed templates are configured
func renderer() multitemplate.Render {
	var layoutsPath = "views/layouts/"
	var pagesPath = "views/pages/"

	r := multitemplate.New()
	// Index page template
	r.AddFromFiles("index", layoutsPath+"application.tmpl", layoutsPath+"header.tmpl", layoutsPath+"footer.tmpl", layoutsPath+"share.tmpl", layoutsPath+"magicbox.tmpl", pagesPath+"index.tmpl")

	// Result page template
	r.AddFromFiles("result", layoutsPath+"application.tmpl", layoutsPath+"header.tmpl", layoutsPath+"footer.tmpl", layoutsPath+"share.tmpl", layoutsPath+"magicbox.tmpl", pagesPath+"result.tmpl")

	// About page template
	r.AddFromFiles("about", layoutsPath+"application.tmpl", layoutsPath+"header.tmpl", layoutsPath+"footer.tmpl", layoutsPath+"share.tmpl", pagesPath+"about.tmpl")

	// Tips page template
	r.AddFromFiles("tips", layoutsPath+"application.tmpl", layoutsPath+"header.tmpl", layoutsPath+"footer.tmpl", layoutsPath+"share.tmpl", pagesPath+"tips.tmpl")

	return r
}
