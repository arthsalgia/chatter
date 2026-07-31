package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/arthsalgia/messages-api/controllers"
	"github.com/arthsalgia/messages-api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed chatter-app/dist/*
var frontendFiles embed.FS

func main() {
	dbPathPtr := flag.String("db", "./chat.db", "path to the iMessage SQLite chat file")
	flag.Parse()

	fmt.Println("Starting server...")
	fmt.Println("Reading through", *dbPathPtr, "...")

	services.ConnectDb(*dbPathPtr)
	start := controllers.LoadAllMessages()
	if !start {
		fmt.Println("Empty or invalid database")
		fmt.Println("Stopping server...")
		return
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost8000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 1. Register API Routes first
	api := router.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello from the Go API!"})
		})
		api.GET("/get-all", controllers.GetAll)
		api.GET("/best-friend", controllers.BestFriend)
		api.GET("/biggest-fan", controllers.BiggestFan)
		api.GET("/celebrity", controllers.Celebrity)
		api.GET("/nth-common", controllers.NthCommon)
		api.GET("/most-common-word", controllers.MostCommonWord)
		api.GET("/most-texted-date", controllers.MostTextedDate)
		api.GET("/search", controllers.Search)
		api.GET("/meta-data", controllers.MetaData)
	}

	// 2. Embedded React Frontend Setup
	distFS, err := fs.Sub(frontendFiles, "chatter-app/dist")
	if err != nil {
		panic(err)
	}

	assetServer := http.FileServer(http.FS(distFS))

	// 3. Serve static files and handle SPA routes safely via NoRoute (avoids router prefix crashes)
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// Guard: Never catch API routes here
		if len(path) >= 4 && path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{"error": "API route not found"})
			return
		}

		// Check if the requested static file exists in the embedded dist folder (e.g., JS, CSS, favicon)
		filePath := path
		if filePath == "/" {
			filePath = "index.html"
		} else {
			// strip leading slash for fs.Open
			filePath = path[1:]
		}

		f, err := distFS.Open(filePath)
		if err == nil {
			f.Close()
			assetServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		// Fallback to index.html for React Router client-side routes
		indexFile, err := distFS.Open("index.html")
		if err != nil {
			c.String(http.StatusNotFound, "React build index.html not found")
			return
		}
		defer indexFile.Close()

		stat, _ := indexFile.Stat()
		c.DataFromReader(http.StatusOK, stat.Size(), "text/html; charset=utf-8", indexFile, map[string]string{})
	})

	fmt.Println("🚀 Server running on http://localhost:8000")
	router.Run(":8000")
}
