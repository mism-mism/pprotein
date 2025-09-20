package main

import (
	"flag"
	"os"

	"github.com/kaz/pprotein/integration/standalone"
)

func main() {
	var (
		port    = flag.String("port", "19000", "Port to listen on")
		gitDir  = flag.String("git-dir", "", "Path to git repository directory (default: current directory)")
		help    = flag.Bool("help", false, "Show help message")
	)
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	// 環境変数からポート番号を取得（下位互換性のため）
	if envPort := os.Getenv("PORT"); envPort != "" {
		*port = envPort
	}

	// ポート番号にコロンプレフィックスを追加
	addr := ":" + *port

	standalone.Integrate(addr, *gitDir)
}
