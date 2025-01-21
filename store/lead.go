package store

import (
	"github.com/minhnhut123456/golang_app_crm/model"
)

type LeadStore struct {
	Store
}

func (s *LeadStore) Create(lead *model.Lead) error {
	err := s.Store.db.Create(lead).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *LeadStore) Upsert( lead *model.Lead) error {
	err := s.db.Save(lead).Error
	if err !=nil{
		return err
	}

	return nil
}

func (s *LeadStore) Delete( leadID int64) error {
	err := s.db.Delete(&model.Lead{}, leadID).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *LeadStore) FindByID(leadID int64) (*model.Lead, error) {
	lead := model.Lead{}
	err := s.db.First(&lead, leadID).Error
	if err != nil {
		return nil, err
	}

	return &lead, nil
}