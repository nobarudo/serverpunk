package monitor

import (
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// ServerState はさーばーちゃんの機嫌を決める0〜100のスコア
type ServerState struct {
	CPU     float64 `json:"cpu"`
	Memory  float64 `json:"memory"`
	Network float64 `json:"network"` // MaxMbpsを100%とした割合
}

var (
	mu           sync.RWMutex
	currentState ServerState
)

// Start はバックグラウンドで監視を開始します（main関数で1度だけ呼ぶ）
// maxMbpsには回線の上限（例: さくらのVPSなら 100.0）を指定します
func Start(interval time.Duration, maxMbps float64) {
	go func() {
		// 初回のネットワーク通信量を取得
		prevNet, _ := net.IOCounters(false)

		for {
			time.Sleep(interval) // 例: 1秒待機

			// 1. CPUとメモリの取得
			c, _ := cpu.Percent(0, false)
			v, _ := mem.VirtualMemory()

			// 2. ネットワークの通信量（差分）を取得
			currNet, _ := net.IOCounters(false)
			var mbps float64
			if len(currNet) > 0 && len(prevNet) > 0 {
				bytesSent := currNet[0].BytesSent - prevNet[0].BytesSent
				bytesRecv := currNet[0].BytesRecv - prevNet[0].BytesRecv

				// Byte/sec を Mbps に変換
				totalBytes := float64(bytesSent + bytesRecv)
				mbps = (totalBytes * 8) / 1000000 / interval.Seconds()
			}
			prevNet = currNet

			// ネットワーク使用率を 0〜100 のパーセンテージに変換
			netPercent := (mbps / maxMbps) * 100
			if netPercent > 100 {
				netPercent = 100
			}

			// 3. 最新のステータスを安全に書き込み
			mu.Lock()
			if len(c) > 0 {
				currentState.CPU = c[0]
			}
			currentState.Memory = v.UsedPercent
			currentState.Network = netPercent
			mu.Unlock()
		}
	}()
}

// GetCurrentState は最新のシステム負荷を即座に返します（APIハンドラ内で呼ぶ）
func getCurrentState() ServerState {
	mu.RLock()
	defer mu.RUnlock()
	return currentState
}
