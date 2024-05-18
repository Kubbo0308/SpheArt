package utils

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dyatlov/go-opengraph/opengraph"
)

// URLからOGPのインスタンスを取得する
func GetOGPImageFromURL(url string) (string, error) {
	// http.Client の設定
	client := &http.Client{
		// 10秒のタイムアウト
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// ライブラリのOGPインスタンス
	og := opengraph.NewOpenGraph()
	// HTMLの解析
	err = og.ProcessHTML(strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}

	// OGPの画像URLを取得
	if len(og.Images) > 0 {
		return og.Images[0].URL, nil
	}
	// 画像が見つからない場合は空文字を返す
	return "", nil
}
