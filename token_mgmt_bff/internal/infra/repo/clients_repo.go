package repo

import "dxps.io/token_mgmt_bff/internal/domain/model"

type ClientsRepo struct {
	store []*model.Client
}
