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

//go:embed all:build
var frontendAssets embed.FS

func setupFrontend(router *gin.Engine) {
	assets, err := fs.Sub(frontendAssets, "build")
	if err != nil {
		log.Fatalf("フロントエンドのビルドファイルが見つかりません: %v", err)
	}

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/serverpunk/")
	})

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 1. API宛の通信は絶対に404を返す
		if strings.HasPrefix(path, "/serverpunk/api/") || strings.HasPrefix(path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"error": "API Endpoint not found"})
			return
		}

		searchPath := strings.TrimPrefix(path, "/serverpunk")
		searchPath = strings.TrimPrefix(searchPath, "/")

		log.Printf("[SPA Debug] Req: '%s' -> Search: '%s'", path, searchPath)

		// 2. ファイル探索（JSやCSS用）
		f, err := assets.Open(searchPath)
		if err == nil {
			stat, _ := f.Stat()
			if stat != nil && !stat.IsDir() {
				f.Close()
				c.FileFromFS("/"+searchPath, http.FS(assets))
				return
			}
			f.Close()
		}

		log.Printf("[SPA Debug] Serving index.html directly for: '%s'", path)

		// 🌟 修正箇所：Goの FileFromFS を使わず、メモリから直接読み込んで返す！
		// これにより、Goのお節介な301リダイレクトを完全に無効化します
		indexData, err := fs.ReadFile(assets, "index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "index.html not found in binary")
			return
		}

		// 純粋なHTMLデータとしてステータス200で強制的に送り返す
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexData)
	})
}
