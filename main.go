package main

import (
	"fmt"
	"grappa"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsData(t time.Time) string {
	year, month, date := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, date)
}

func main() {
	//r := grappa.New()
	//r.Use(grappa.Logger())
	//r.SetFuncMap(template.FuncMap{
	//	"FormatAsDate": FormatAsData,
	//})
	//r.LoadHTMLGlob("templates/*")
	//r.Static("/assets", "./static")
	//
	//stu1 := &student{Name: "Kiracoon", Age: 20}
	//stu2 := &student{Name: "Kira", Age: 22}
	////r.GET("/index", func(c *grappa.Context) {
	////	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	////})
	//r.GET("/", func(c *grappa.Context) {
	//	c.HTML(http.StatusOK, "css.tmpl", nil)
	//})
	//r.GET("/students", func(c *grappa.Context) {
	//	c.HTML(http.StatusOK, "arr.tmpl", grappa.H{
	//		"title":  "Grappa",
	//		"stuArr": [2]*student{stu1, stu2},
	//	})
	//})
	//
	//r.GET("/date", func(c *grappa.Context) {
	//	c.HTML(http.StatusOK, "custom_func.tmpl", grappa.H{
	//		"title": "Grappa",
	//		"now":   time.Date(2024, 8, 23, 4, 29, 0, 0, time.UTC),
	//	})
	//})
	////
	////v1 := r.Group("/v1")
	////{
	////	v1.GET("/", func(c *grappa.Context) {
	////		c.HTML(http.StatusOK, "<h1>Welcome to Grappa!<h1>")
	////	})
	////	v1.GET("/ping", func(c *grappa.Context) {
	////		c.String(http.StatusOK, "%s Pong! Hello %s!\n", c.Path, c.Query("name"))
	////	})
	////}
	////
	////v2 := r.Group("/v2")
	////{
	////	v2.GET("/ping/:name", func(c *grappa.Context) {
	////		c.String(http.StatusOK, "%s Pong! Hello %s!\n", c.Path, c.Params["name"])
	////	})
	////
	////	v2.GET("/assets/*filepath", func(c *grappa.Context) {
	////		c.JSON(http.StatusOK, grappa.H{"filepath": c.Param("filepath")})
	////	})
	////
	////	v2.POST("/login", func(c *grappa.Context) {
	////		c.JSON(http.StatusOK, grappa.H{
	////			"username": c.PostForm("username"),
	////			"password": c.PostForm("password"),
	////		})
	////	})
	////}

	r := grappa.Default()
	r.GET("/", func(c *grappa.Context) {
		c.String(http.StatusOK, "Hello Kiracoon!")
	})

	r.GET("/panic", func(c *grappa.Context) {
		names := []string{"Kiracoon"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":7777")
}
