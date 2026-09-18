package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MakMoinee/go-boiler-plate/internal/common"
	"github.com/MakMoinee/go-boiler-plate/internal/models"
	"github.com/MakMoinee/go-mith/pkg/mithclientrest"
)

type RestAPI interface {
	GetUsers(parentContext context.Context, schema any) ([]models.UserRequest, error)
}

type Restapi struct {
	Users  models.UserRequest
	Client *mithclientrest.Client
}

func (r *Restapi) GetUsers(parentContext context.Context, schema any) ([]models.UserRequest, error) {
	common.LOGGER.Info("inside restapi>GetUsers()")
	ctx, cancel := context.WithTimeout(parentContext, 10*time.Second)
	defer cancel()

	resp, err := r.Client.Do(ctx, schema)
	if err != nil {
		return []models.UserRequest{}, err
	}

	userList := []models.UserRequest{}
	if err := json.Unmarshal(resp, &userList); err != nil {
		common.LOGGER.ErrorContext(ctx, fmt.Sprintf("decode response: %v", err))
		return []models.UserRequest{}, err
	}

	return userList, nil
}

func NewRestApi(baseURL string, resolver mithclientrest.OperationResolver) RestAPI {
	restApi := Restapi{
		Client: mithclientrest.New(&mithclientrest.Config{
			BaseURL: baseURL,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			OperationResolver: resolver,
		}),
	}

	return &restApi
}
