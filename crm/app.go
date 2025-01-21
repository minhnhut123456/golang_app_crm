package crm

import (
	"net/http"

	"github.com/minhnhut123456/golang_app_crm/netkit"
	"github.com/minhnhut123456/golang_app_crm/store"
)

var (
	RouteCreateLead = &netkit.Route{
		Name:         "create_lead",
		Method:       http.MethodPost,
		Path:         "/lead/create",
	}
	RouteGetLead = &netkit.Route{
		Name:         "get_lead",
		Method:       http.MethodGet,
		Path:         "/lead/{id}",
	}
)

// demo apply golang functional pattern
type AppOption func (*App)

type App struct {
	Server *netkit.HTTPServer
	config1 string
	config2 string
}

func NewApp(stores *store.Stores, options ...AppOption) *App {
	leadHandler := NewLeadHandler(stores)

	app := &App{
		Server: netkit.NewHTTPServer([]*netkit.RouteHandler{
			{Route: RouteCreateLead, Handler: leadHandler.handleCreateLead},
			{Route: RouteGetLead, Handler: leadHandler.handleGetLead},
		}),
	}

	for _,o := range options {
		o(app)
	}
	return app
}

func WithConfig1(config1 string) AppOption {
	return func(a *App) {
		a.config1 = config1
	}
}

func WithConfig2(config2 string) AppOption {
	return func(a *App) {
		a.config2 = config2
	}
}
