package store

import (
	"context"

	"github.com/minhnhut123456/golang_app_crm/database"
	"gorm.io/gorm"
)

type Stores struct {
	db *gorm.DB
	LeadStore *LeadStore
}

func NewStores(ctx context.Context) (*Stores,error){
	m, err := database.NewMySql("config/crm_db.yaml", database.WithMigation("migration"))
	if(err !=nil){
		return nil, err
	}

	err = m.Connect()
	if(err !=nil){
		return nil, err
	}

	err = m.MigrateUp()
	if(err !=nil){
		return nil, err
	}

	return &Stores{LeadStore: &LeadStore{Store: NewStore(m.DB)}, db: m.DB}, nil
}

// eliminate the Close() method on version 1.20 because GORM supports connection pooling
// func (s *Stores) Close(){
	
// }