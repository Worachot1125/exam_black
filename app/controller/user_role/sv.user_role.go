package user_role

import (
	"app/app/model"
	"app/app/request"
	"context"
)

func (s *Service) Create(ctx context.Context, req request.CreateUser_Role) (*model.User_Role, bool, error) {
	m := &model.User_Role{
		User_ID: req.User_ID,
		Role_ID: req.Role_ID,
	}
	_, err := s.db.NewInsert().Model(m).Exec(ctx)

	return m, false, err
}
