package repo

import (
	"dxps.io/token_mgmt_bff/internal/domain/model"
	"dxps.io/token_mgmt_bff/internal/errs"
)

type ClientsRepo struct {
	store []*model.Client
}

func NewClientsRepo() *ClientsRepo {
	return &ClientsRepo{
		store: []*model.Client{
			{ID: "joe", Secret: "black"},
		},
	}
}

func (r *ClientsRepo) Get(clientID string) (*model.Client, error) {

	for _, c := range r.store {
		if c.ID == clientID {
			return c, nil
		}
	}
	return nil, errs.ErrNotFound
}
