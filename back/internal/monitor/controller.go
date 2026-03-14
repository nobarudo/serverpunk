package monitor

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FaceStatus は外部に公開して良い「抽象化された」データ
type FaceStatus struct {
	Heat     int    `json:"heat"`     // CPUの代わり（0〜100の整数）
	LeftEye  string `json:"left_eye"` // メモリの代わり
	RightEye string `json:"right_eye"`
	Mouth    string `json:"mouth"`
	Noise    int    `json:"noise"` // ネットワークの代わり
}

func GetStatus(c *gin.Context) {
	state := getCurrentState() // 以前作った関数

	// 1. CPU: 綺麗なグラデーションを活かすため、小数点以下を四捨五入して「熱量(Heat)」とする
	heat := int(math.Round(state.CPU))

	// 2. メモリ: サーバー内で顔のパーツを決定し、生データはここで破棄する
	leftEye, rightEye, mouth := "・", "・", "_"
	switch {
	case state.Memory > 90:
		leftEye, rightEye, mouth = "×", "×", "д"
	case state.Memory > 80:
		leftEye, rightEye, mouth = ">", "<", "o"
	case state.Memory > 70:
		leftEye, rightEye, mouth = ">", "<", "_"
	}

	// 3. ネットワーク: 小数を丸めて「ノイズ強度」とする
	noise := int(math.Round(state.Network))

	// 抽象化したデータだけをフロントエンドに渡す
	c.JSON(http.StatusOK, FaceStatus{
		Heat:     heat,
		LeftEye:  leftEye,
		RightEye: rightEye,
		Mouth:    mouth,
		Noise:    noise,
	})
}
