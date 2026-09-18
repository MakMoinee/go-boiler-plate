package routes

import (
	"fmt"
	"net/http"

	"github.com/MakMoinee/go-boiler-plate/internal/common"
	"github.com/MakMoinee/go-boiler-plate/internal/models"
	"github.com/MakMoinee/go-boiler-plate/internal/restapi"
	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/MakMoinee/go-mith/pkg/response"
	"github.com/go-chi/cors"
)

type routesInfo struct {
	RouteName  string
	RouteValue string
}

type handler struct {
	UserRestAPI restapi.RestAPI
}

type Routes interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

func newRoutes() Routes {
	handler := handler{}
	handler.UserRestAPI = restapi.NewRestApi(common.USER_API, models.UserResolver{})
	return &handler
}

// Set Routes
func Set(httpService *goserve.Service) {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders:   []string{"Link", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	newRoutesHandler := newRoutes()
	httpService.Router.Use(cors.Handler)
	initiateRoutes(httpService, newRoutesHandler)
}

func initiateRoutes(httpService *goserve.Service, handler Routes) {
	set := make(map[string]interface{})
	infos := []routesInfo{}

	httpService.Router.Get(common.GetUsersPath, handler.GetUsers)
	infos = append(infos, routesInfo{RouteName: "RetrieveUsers", RouteValue: common.GetUsersPath})

	set["routes"] = infos
	set["version"] = common.SERVICE_VERSION

	httpService.SetInfo(set)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	common.LOGGER.Info(fmt.Sprint("Invoked GetUsers()"))
	userList, err := h.UserRestAPI.GetUsers(ctx, models.UserRequest{})
	if err != nil {
		common.LOGGER.Error(err.Error())
		response.Error(w, response.ErrorResponse{
			ErrorStatus:  http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		})
		return
	}

	response.Success(w, userList)

}
