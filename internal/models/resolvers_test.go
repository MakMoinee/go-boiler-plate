package models

import (
	"net/http"
	"strings"
	"testing"

	"github.com/MakMoinee/go-mith/pkg/mithclientrest"
)

func TestUserResolver_OperationFor(t *testing.T) {
	resolver := UserResolver{}

	tests := []struct {
		name        string
		payload     any
		want        mithclientrest.Operation
		wantErr     bool
		errContains string
	}{
		{
			name:    "UserRequest",
			payload: UserRequest{},
			want: mithclientrest.Operation{
				Method: http.MethodGet,
				Path:   "/users",
			},
		},
		{
			name:        "unsupported payload",
			payload:     "invalid payload",
			wantErr:     true,
			errContains: "unsupported payload string",
		},
		{
			name:        "nil payload",
			payload:     nil,
			wantErr:     true,
			errContains: "unsupported payload <nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.OperationFor(tt.payload)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}

				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf(
						"error = %q, want error containing %q",
						err.Error(),
						tt.errContains,
					)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.Method != tt.want.Method {
				t.Errorf("OperationFor() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
