package emergency_report

import (
	"app/app/request"
	"app/app/response"
	"app/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctl *Controller) Create(ctx *gin.Context) {
	// 1) ดึง user_id จาก ctx
	uidStr, ok := ctx.Get("user_id")
	if !ok {
		response.Unauthorized(ctx, "missing user_id")
		return
	}
	userID, err := uuid.Parse(uidStr.(string))
	if err != nil {
		response.BadRequest(ctx, "invalid user_id format")
		return
	}

	// 2) Bind JSON body แทนการใช้ ctx.PostForm()
	var req request.CreateEmergencyReport
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "ข้อมูลส่งมาไม่ถูกต้อง")
		return
	}

	// 3) เติม user_id จาก token/session
	req.User_ID = userID.String()

	// 4) ไม่บังคับกรอก emergency_type_id หรือ description
	//    (service จะจัดการเอง)

	// 5) เรียก service สร้างข้อมูล
	if _, mserr, err := ctl.Service.Create(ctx, req); err != nil {
		msg := "internal server error"
		if mserr {
			msg = err.Error()
		}
		response.InternalError(ctx, msg)
		return
	}

	response.Success(ctx, gin.H{"message": "สร้างเหตุฉุกเฉินเรียบร้อย"})
}

func (ctl *Controller) Update(ctx *gin.Context) {
	ID := request.GetByIDEmergencyReport{}
	if err := ctx.BindUri(&ID); err != nil {
		logger.Errf(err.Error())
		response.BadRequest(ctx, err.Error())
		return
	}

	body := request.UpdateEmergencyReport{}
	if err := ctx.Bind(&body); err != nil {
		logger.Errf(err.Error())
		response.BadRequest(ctx, err.Error())
		return
	}

	_, mserr, err := ctl.Service.Update(ctx, body, ID)
	if err != nil {
		ms := "internal server error"
		if mserr {
			ms = err.Error()
		}
		logger.Errf(err.Error())
		response.InternalError(ctx, ms)
		return
	}

	response.Success(ctx, nil)
}

func (ctl *Controller) List(ctx *gin.Context) {
	req := request.ListEmergencyReport{}
	if err := ctx.Bind(&req); err != nil {
		logger.Errf(err.Error())
		response.BadRequest(ctx, err.Error())
		return
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.Size == 0 {
		req.Size = 10
	}

	if req.OrderBy == "" {
		req.OrderBy = "asc"
	}

	if req.SortBy == "" {
		req.SortBy = "created_at"
	}

	data, total, err := ctl.Service.List(ctx, req)
	if err != nil {
		logger.Errf(err.Error())
		response.InternalError(ctx, err.Error())
		return
	}
	response.SuccessWithPaginate(ctx, data, req.Size, req.Page, total)
}

func (ctl *Controller) Get(ctx *gin.Context) {
	ID := request.GetByIDEmergencyReport{}
	if err := ctx.BindUri(&ID); err != nil {
		logger.Errf(err.Error())
		response.BadRequest(ctx, err.Error())
		return
	}

	data, err := ctl.Service.Get(ctx, ID)
	if err != nil {
		logger.Errf(err.Error())
		response.InternalError(ctx, err.Error())
		return
	}
	response.Success(ctx, data)
}

func (ctl *Controller) GetByUserIDEmergency(ctx *gin.Context) {
	var req request.GetByUserIDEmergency
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.BadRequest(ctx, "user_id ไม่ถูกต้อง")
		return
	}

	emergencies, err := ctl.Service.GetByUserIDEmergency(ctx, req)
	if err != nil {
		logger.Errf("Failed to get emergencies by user_id: %v", err)
		response.InternalError(ctx, "ไม่สามารถดึงข้อมูลได้")
		return
	}

	response.Success(ctx, emergencies)
}

func (ctl *Controller) Delete(ctx *gin.Context) {
	ID := request.GetByIDEmergencyReport{}
	if err := ctx.BindUri(&ID); err != nil {
		logger.Errf(err.Error())
		response.BadRequest(ctx, err.Error())
		return
	}

	err := ctl.Service.Delete(ctx, ID)
	if err != nil {
		logger.Errf(err.Error())
		response.InternalError(ctx, err.Error())
		return
	}
	response.Success(ctx, nil)
}
