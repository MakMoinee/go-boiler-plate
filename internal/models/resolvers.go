package models

import (
	"fmt"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/mithclientrest"
)

type UserResolver struct{}

func (UserResolver) OperationFor(payload any) (mithclientrest.Operation, error) {
	switch payload.(type) {
	case UserRequest:
		return mithclientrest.Operation{
			Method: http.MethodGet,
			Path:   "/users",
		}, nil
	default:
		return mithclientrest.Operation{}, fmt.Errorf("unsupported payload %T", payload)
	}
}
