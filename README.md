
# 🖥️ Serverpunk

> サーバーの鼓動を、視覚と直感で監視する。
> 超軽量・環境非依存のシングルバイナリ・モニタリングダッシュボード。

**Serverpunk** は、サーバーのリソース状態（CPU、メモリ、負荷など）をリアルタイムで監視するWebダッシュボードです。最大の特長は、サーバーの負荷状態に応じて画面中央の表情やノイズが変化する、サイバーパンク/レトロモニター調のUIです。

## ✨ Features

- **👀 直感的な視覚モニタリング**: 数字だけでなく、顔文字（`・_・`）の表情や画面の揺れ（ストレス）、グリッチノイズでサーバーの悲鳴を直感的に察知できます。
- **📱 レスポンシブ対応**: PCのワイド画面でも、スマートフォンの縦長画面でも最適化されたUIで表示されます。

## 🛠️ Technology Stack

- **Backend**: Go (Gin)
- **Frontend**: SvelteKit 5 (adapter-static / SPA mode)
- **Styling**: Vanilla CSS (Flexbox / CSS Animations)
- **Build**: Podman / Docker (Multi-stage build)

## 🚀 Getting Started

### Prerequisites
ビルドには `podman` (または `docker`) と `make` が必要です。ホスト環境にGoやNode.jsを用意する必要はありません。

### Build

```bash
git clone https://github.com/nobarudo/serverpunk.git
cd serverpunk
make build
```

ビルドが完了すると、ディレクトリ内に単一の実行ファイル（`serverpunk`）が生成されます。

### Run

```bash
./serverpunk
```

デフォルトでは `http://localhost:8080/serverpunk/` でダッシュボードが起動します。

## 📜 License

MIT License
