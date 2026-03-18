APP_NAME := serverpunk
IMAGE_NAME := serverpunk-builder
CONTAINER_NAME := temp-builder

# デフォルトの挙動を指定
.PHONY: all lint build run clean

all: lint build

lint:
	semgrep scan --config p/golang \
	             --config p/security-audit \
	             --config p/javascript \
	             --config p/typescript \
	             --config p/secrets .

# 1. コンテナでビルドしてバイナリを取り出す一連のコマンド
build:
	@echo "🚀 Podmanでビルドを開始します..."
	podman build -t $(IMAGE_NAME) .
	@echo "📦 バイナリを抽出しています..."
	-podman rm -f $(CONTAINER_NAME) 2>/dev/null || true
	podman create --name $(CONTAINER_NAME) $(IMAGE_NAME)
	podman cp $(CONTAINER_NAME):/app/serverpunk ./$(APP_NAME)
	podman rm $(CONTAINER_NAME)
	@echo "✨ ビルド完了！ $(APP_NAME) が生成されました！"

run:
	@echo "🚀 サーバーちゃんを起動します..."
	podman run --rm -p 8080:8080 --env-file back/configs/server.env $(IMAGE_NAME)

# 2. お掃除コマンド（やり直したい時用）
clean:
	@echo "🧹 不要なファイルを削除します..."
	podman rmi $(IMAGE_NAME) || true
	@echo "✨ クリーンアップ完了！"
