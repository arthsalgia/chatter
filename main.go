package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/arthsalgia/messages-api/controllers"
	"github.com/arthsalgia/messages-api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed chatter-app/dist/*
var frontendFiles embed.FS

func main() {
	homeDir, err := os.UserHomeDir()
	defaultIMessagePath := ""
	if err == nil {
		defaultIMessagePath = filepath.Join(homeDir, "Library", "Messages", "chat.db")
	}

	dbPathPtr := flag.String("db", "./chat.db", "path to the iMessage SQLite chat file")
	flag.Parse()

	dbPath := *dbPathPtr

	if *dbPathPtr == "./chat.db" {
		if _, err := os.Stat("./chat.db"); err == nil {
			dbPath = "./chat.db"
			fmt.Println("Found local copy of chat.db in current directory.")
		} else if defaultIMessagePath != "" {
			if _, err := os.Stat(defaultIMessagePath); err == nil {
				testFile, readErr := os.Open(defaultIMessagePath)
				if readErr == nil {
					testFile.Close()
					dbPath = defaultIMessagePath
					fmt.Println("Access granted to system iMessage database.")
				}
			}
		}
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Println("\nError: iMessage database could not be found or accessed!")
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("To use Chatter, please choose ONE of the following options:")
		fmt.Println("")
		fmt.Println("  Option 1 (Easiest):")
		fmt.Println("  Copy your 'chat.db' file into the folder you are currently running 'chatter' from.")
		fmt.Println("  (You can find it on your Mac at: ~/Library/Messages/chat.db)")
		fmt.Println("")
		fmt.Println("  Option 2:")
		fmt.Println("  Grant your Terminal (or iTerm) 'Full Disk Access' via:")
		fmt.Println("  System Settings -> Privacy & Security -> Full Disk Access")
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("Stopping application...")
		os.Exit(1)
	}

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

	api := router.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello from Chatter!"})
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
		api.GET("/sentiment-analysis", controllers.SentimentAnalysis)
		api.GET("/get-all-chats", controllers.GetAllChats)
	}

	distFS, err := fs.Sub(frontendFiles, "chatter-app/dist")
	if err != nil {
		panic(err)
	}

	assetServer := http.FileServer(http.FS(distFS))

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		if len(path) >= 4 && path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{"error": "API route not found"})
			return
		}

		filePath := path
		if filePath == "/" {
			filePath = "index.html"
		} else {
			filePath = path[1:]
		}

		f, err := distFS.Open(filePath)
		if err == nil {
			f.Close()
			assetServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		indexFile, err := distFS.Open("index.html")
		if err != nil {
			c.String(http.StatusNotFound, "React build index.html not found")
			return
		}
		defer indexFile.Close()

		stat, _ := indexFile.Stat()
		c.DataFromReader(http.StatusOK, stat.Size(), "text/html; charset=utf-8", indexFile, map[string]string{})
	})

	fmt.Println("Server running on http://localhost:8000")
	router.Run(":8000")
}
