package main

import (
	"testing"
)

func TestAuthToken_IsValid(t *testing.T) {
	tests := []struct {
		name string
		t    AuthToken
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.IsValid(); got != tt.want {
				t.Errorf("AuthToken.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthToken_IsSet(t *testing.T) {
	tests := []struct {
		name string
		t    AuthToken
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.IsSet(); got != tt.want {
				t.Errorf("AuthToken.IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
