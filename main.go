package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Render!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	// Renderから提供されるPORT環境変数を取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // ローカル開発用のデフォルトポート
	}

	log.Printf("Server starting on port %s", port)
	// "localhost:8080" ではなく ":8080" のようにコロンから始める
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
