package user

import (
	"app/app/request"
	"app/app/response"
	"app/internal/logger"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) Create(ctx *gin.Context) {
	body := request.CreateUser{}

	if err := ctx.Bind(&body); err != nil {
		logger.Errf(err.Error())
		response.BadRequest(ctx, err.Error())
		return
	}

	_, mserr, err := ctl.Service.Create(ctx, body)
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

func (ctl *Controller) Update(ctx *gin.Context) {
	ID := request.GetByIDUser{}
	if err := ctx.BindUri(&ID); err != nil {
		logger.Errf(err.Error())
		response.BadRequest(ctx, err.Error())
		return
	}

	body := request.UpdateUser{}
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
	req := request.ListUser{}
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
	ID := request.GetByIDUser{}
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

func (ctl *Controller) Delete(ctx *gin.Context) {
	ID := request.GetByIDUser{}
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
