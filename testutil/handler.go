package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg any
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("failed to unmarshal want: %v", err)
	}

	if err := json.Unmarshal(got, &jg); err != nil {
		t.Fatalf("failed to unmarshal got: %v", err)
	}

	if diff := cmp.Diff(jw, jg); diff != "" {
		t.Fatalf("got != want:\n%s", diff)
	}
}

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })

	body, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	if got.StatusCode != status {
		t.Fatalf("status code mismatch: got %d, want %d", got.StatusCode, status)
	}

	if len(gb) == 0 && len(body) == 0 {
		return
	}

	AssertJSON(t, body, gb)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	return bt
}
