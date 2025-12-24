package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// 1. テスト用のリクエストを作成
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2. レスポンスを記録するためのRecorderを作成
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	// 3. ハンドラーを直接呼び出し、リクエストとレコーダーを渡す
	handler.ServeHTTP(rr, req)

	// 4. レスポンスのステータスコードを検証
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 5. レスポンスのボディを検証
	expected := "Hello from Render!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
