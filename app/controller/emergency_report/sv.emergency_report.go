package emergency_report

import (
	"app/app/model"
	"app/app/request"
	"app/app/response"
	"app/internal/logger"
	"context"
	"errors"
	"fmt"
)

func (s *Service) Create(ctx context.Context, req request.CreateEmergencyReport) (*model.Emergency_report, bool, error) {

	// ถ้าไม่ส่ง emergency_type_id ให้เป็น "" (หรือ NULL ภายหลัง)
	emTypeID := req.Emergency_Type_ID
	if emTypeID == "" {
		emTypeID = "" // หรือเก็บ NULL ได้ถ้า column รับ
	}

	m := &model.Emergency_report{
		User_ID:           req.User_ID, // ⇦ string
		Emergency_Type_ID: emTypeID,    // ⇦ string
		Description:       req.Description,
		Image_URL:         req.Image_URL,
		Status:            "Pending", // ค่าเริ่มต้น
	}

	_, err := s.db.NewInsert().Model(m).Exec(ctx)
	return m, false, err
}

func (s *Service) Update(
	ctx context.Context,
	req request.UpdateEmergencyReport,
	id request.GetByIDEmergencyReport,
) (*model.Emergency_report, bool, error) {

	// ---------- 1) หาแถวเดิม ----------
	var m model.Emergency_report
	if err := s.db.
		NewSelect().
		Model(&m).
		Where("id = ?", id.ID).
		Where("deleted_at IS NULL").
		Scan(ctx); err != nil {

		if errors.Is(err, nil) {
			return nil, true, errors.New("emergency_report not found")
		}
		return nil, false, err
	}

	// ---------- 2) อัปเดตเฉพาะฟิลด์ที่ส่งมา ----------
	if req.Description != "" {
		m.Description = req.Description
	}
	if req.Image_URL != "" {
		m.Image_URL = req.Image_URL
	}
	if req.Status != "" {
		m.Status = req.Status
	}

	m.SetUpdateNow() // updated_at = now

	// ---------- 3) UPDATE ----------
	_, err := s.db.NewUpdate().
		Model(&m).
		Column("description", "image_url", "status", "updated_at").
		WherePK(). // m.ID มีค่าแล้ว → ไม่ error
		Exec(ctx)
	if err != nil {
		return nil, false, err
	}

	return &m, false, nil
}

func (s *Service) List(ctx context.Context, req request.ListEmergencyReport) ([]response.EmergencyReportResponse, int, error) {
	offset := (req.Page - 1) * req.Size

	m := []response.EmergencyReportResponse{}
	query := s.db.NewSelect().
		TableExpr("emergency_reports AS er").
		Column("er.id", "er.user_id", "er.emergency_type_id", "er.description", "er.image_url", "er.status", "er.created_at", "er.updated_at").
		Where("deleted_at IS NULL")

	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	order := fmt.Sprintf("er.%s %s", req.SortBy, req.OrderBy)

	err = query.Order(order).Limit(req.Size).Offset(offset).Scan(ctx, &m)
	if err != nil {
		return nil, 0, err
	}
	return m, count, err
}

func (s *Service) Get(ctx context.Context, id request.GetByIDEmergencyReport) (*response.EmergencyReportResponse, error) {
	m := response.EmergencyReportResponse{}
	err := s.db.NewSelect().
		TableExpr("emergency_reports AS er").
		Column("er.id", "er.user_id", "er.student_id", "er.emergency_type_id", "er.description", "er.image_url", "er.status", "er.created_at", "er.updated_at").
		Where("id = ?", id.ID).Where("deleted_at IS NULL").Scan(ctx, &m)
	return &m, err
}

func (s *Service) GetByUserIDEmergency(ctx context.Context, req request.GetByUserIDEmergency) ([]model.Emergency_report, error) {
	var emergencies []model.Emergency_report
	err := s.db.NewSelect().
		Model(&emergencies).
		Where("user_id = ? AND deleted_at IS NULL", req.UserID).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		logger.Errf("Database query failed: %v", err) // เพิ่มบรรทัดนี้
		return nil, err
	}
	return emergencies, nil
}

func (s *Service) Delete(ctx context.Context, id request.GetByIDEmergencyReport) error {
	ex, err := s.db.NewSelect().Table("emergency_reports").Where("id = ?", id.ID).Where("deleted_at IS NULL").Exists(ctx)
	if err != nil {
		return err
	}

	if !ex {
		return errors.New("emergency_report not found")
	}

	// data, err := s.db.NewDelete().Table("users").Where("id = ?", id.ID).Exec(ctx)
	_, err = s.db.NewDelete().Model((*model.Emergency_report)(nil)).Where("id = ?", id.ID).Exec(ctx)
	return err
}
