import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [sveltekit()],
  server: {
    proxy: {
      // '/api' から始まるリクエストをすべてGo(8080ポート)へ転送する
      "/serverpunk/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/serverpunk/, ""),
      },
    },
  },
});
