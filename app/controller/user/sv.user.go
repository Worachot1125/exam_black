package user

import (
	"app/app/model"
	"app/app/request"
	"app/app/response"
	"app/internal/logger"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Create(ctx context.Context, req request.CreateUser) (*model.User, bool, error) {
	var usr *model.User
	var dupErr bool

	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		// ---------- 1. ค้นหา role "Personnel" ----------
		var role model.Role
		err := tx.NewSelect().
			Model(&role).
			Where("name = ?", "Personnel").
			Limit(1).
			Scan(ctx)
		if err != nil {
			return fmt.Errorf("role 'Personnel' not found: %w", err)
		}

		// ---------- 2. สร้าง hash password ----------
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		// ---------- 3. สร้าง User ----------
		usr = &model.User{
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			User_Number: req.UserNumber,
			Password:    string(hash),
			Phone:       req.Phone,
			Role_ID:     role.ID,
			Address:     req.Address,
		}
		usr.SetCreatedNow()

		if _, err := tx.NewInsert().Model(usr).Returning("*").Exec(ctx); err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				dupErr = true
			}
			return err
		}

		// ---------- 4. สร้าง User_Role ----------
		ur := &model.User_Role{
			User_ID: usr.ID,
			Role_ID: role.ID,
		}
		if _, err := tx.NewInsert().Model(ur).Exec(ctx); err != nil {
			return err
		}

		return nil // commit
	})

	if err != nil {
		if dupErr {
			return nil, true, errors.New("user number already exists")
		}
		return nil, false, err
	}
	return usr, false, nil
}

func (s *Service) Update(ctx context.Context, req request.UpdateUser, id request.GetByIDUser) (*model.User, bool, error) {
	ex, err := s.db.NewSelect().Table("users").Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, false, err
	}

	if !ex {
		return nil, false, err
	}

	m := &model.User{
		ID:          id.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		User_Number: req.UserNumber,
		Phone:       req.Phone,
		Address:     req.Address,
	}
	logger.Info(m)
	m.SetUpdateNow()
	_, err = s.db.NewUpdate().Model(m).
		Set("first_name = ?first_name").
		Set("last_name = ?last_name").
		Set("user_number = ?user_number").
		Set("phone = ?phone").
		Set("address = ?address").
		Set("updated_at = ?updated_at").
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, true, errors.New("user number already exists")
		}
	}
	return m, false, err
}

func (s *Service) List(ctx context.Context, req request.ListUser) ([]response.UserResponse, int, error) {
	offset := (req.Page - 1) * req.Size

	m := []response.UserResponse{}
	query := s.db.NewSelect().
		TableExpr("users AS u").
		Column("u.id", "u.first_name", "u.last_name", "u.user_number", "u.phone", "u.address", "u.created_at", "u.updated_at").
		Where("deleted_at IS NULL")

	if req.Search != "" {
		search := fmt.Sprintf("%" + strings.ToLower(req.Search) + "%")
		if req.SearchBy != "" {
			search := strings.ToLower(req.Search)
			query.Where(fmt.Sprintf("LOWER(u.%s) LIKE ?", req.SearchBy), search)
		} else {
			query.Where("LOWER(first_name) LIKE ?", search)
		}
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	order := fmt.Sprintf("r.%s %s", req.SortBy, req.OrderBy)

	err = query.Order(order).Limit(req.Size).Offset(offset).Scan(ctx, &m)
	if err != nil {
		return nil, 0, err
	}
	return m, count, err
}

func (s *Service) Get(ctx context.Context, id request.GetByIDUser) (*response.UserResponse, error) {
	m := response.UserResponse{}
	err := s.db.NewSelect().
		TableExpr("users AS u").
		Column("u.id", "u.first_name", "u.last_name", "u.user_number", "u.phone", "u.address", "u.created_at", "u.updated_at").
		Where("id = ?", id.ID).Where("deleted_at IS NULL").Scan(ctx, &m)
	return &m, err
}

func (s *Service) Delete(ctx context.Context, id request.GetByIDUser) error {
	ex, err := s.db.NewSelect().Table("users").Where("id = ?", id.ID).Where("deleted_at IS NULL").Exists(ctx)
	if err != nil {
		return err
	}

	if !ex {
		return errors.New("user not found")
	}

	// data, err := s.db.NewDelete().Table("users").Where("id = ?", id.ID).Exec(ctx)
	_, err = s.db.NewDelete().Model((*model.User)(nil)).Where("id = ?", id.ID).Exec(ctx)
	return err
}
