package main

import (
	"context"

	"github.com/minhnhut123456/golang_app_crm/crm"
	"github.com/minhnhut123456/golang_app_crm/store"
)

func main() {
	ctx := context.Background()
	stores, err := store.NewStores(ctx)
	
	if err != nil {
		panic(err)
	}

	app := crm.NewApp(stores)
	
	err = app.Server.Start()
	if(err != nil){
		panic("Error when start application")
	}
}