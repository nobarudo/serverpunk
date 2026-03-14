//go:build !dev

package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed dist/*
var frontendAssets embed.FS

func setupFrontend(router *gin.Engine) {
	assets, err := fs.Sub(frontendAssets, "dist")
	if err != nil {
		log.Fatalf("フロントエンドのビルドファイルが見つかりません: %v", err)
	}

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"error": "API Endpoint not found"})
			return
		}

		localPath := strings.TrimPrefix(path, "/")
		if localPath == "" {
			localPath = "index.html"
		}

		// ファイルが存在するか確認
		file, err := assets.Open(localPath)
		if err != nil {
			// 存在しない場合は index.html を返す
			indexFile, _ := fs.ReadFile(assets, "index.html")
			c.Data(http.StatusOK, "text/html; charset=utf-8", indexFile)
			return
		}
		file.Close()

		// 静的ファイルの配信
		http.FileServer(http.FS(assets)).ServeHTTP(c.Writer, c.Request)
	})
}
