package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333

	t.Setenv("PORT", fmt.Sprintf("%d", wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("failed to create config: %+v", err)
	}

	if got.Port != wantPort {
		t.Errorf("want %d, but got %d", wantPort, got.Port)
	}

	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %q, but got %q", wantEnv, got.Env)
	}
}
