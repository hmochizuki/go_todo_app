package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"` // 値が空の場合フィールドを省略する
}

func RespondJson(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json")
	bodyBytes, err := json.Marshal(body) // marshal: メモリに適した構造体を、通信の意適したキシキ(Json)に変換する
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rsp := ErrResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			fmt.Printf("failed to encode err response: %v", err)
		}
		return
	}

	w.WriteHeader(status)
	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		fmt.Printf("failed to write body: %v", err)
	}
}
