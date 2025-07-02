package emergency_type

import (
	"app/app/model"
	"app/app/request"
	"app/app/response"
	"context"
	"fmt"
	"strings"
)

func (s *Service) Create(ctx context.Context, req request.CreateEmergencyType) (*model.Emergency_Type, bool, error) {

	m := &model.Emergency_Type{
		Name: req.Name,
	}

	_, err := s.db.NewInsert().Model(m).Exec(ctx)

	return m, false, err
}

func (s *Service) List(ctx context.Context, req request.ListEmergencyType) ([]response.EmergencyTypeResponse, int, error) {
	offset := (req.Page - 1) * req.Size

	// ---------- 1. validate / default ----------
	if req.SortBy == "" {
		req.SortBy = "name" // หรือ "id"
	}
	order := strings.ToUpper(req.OrderBy)
	if order != "DESC" {
		order = "ASC"
	}

	// ---------- 2. query ----------
	var result []response.EmergencyTypeResponse
	query := s.db.NewSelect().
		TableExpr("emergency_types AS r").
		Column("r.id", "r.name")

	// --- search ---
	if req.Search != "" {
		like := "%" + strings.ToLower(req.Search) + "%"
		if req.SearchBy != "" {
			query.Where(fmt.Sprintf("LOWER(r.%s) LIKE ?", req.SearchBy), like)
		} else {
			query.Where("LOWER(r.name) LIKE ?", like)
		}
	}

	// --- count ---
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// --- order / paginate ---
	query.
		OrderExpr(fmt.Sprintf("r.%s %s", req.SortBy, order)).
		Limit(req.Size).
		Offset(offset)

	if err := query.Scan(ctx, &result); err != nil {
		return nil, 0, err
	}
	return result, count, nil
}
