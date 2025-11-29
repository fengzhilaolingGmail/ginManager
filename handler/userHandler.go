package handler

import (
	"fmt"
	"ginManager/dto"
	"ginManager/service"
	"ginManager/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{svc: service.NewUserService()}
}

// Page 用户分页  /api/user/list
func (h *UserHandler) Page(c *gin.Context) {
	var req dto.UserListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数错误", err))
		return
	}
	list, total, err := h.svc.Page(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, dto.SuccessPage(list, total))
}

// Create 新增用户  /api/user/add
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.UserAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数错误", err))
		fmt.Println(err)
		return
	}
	if err := h.svc.Create(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// Update 编辑用户  /api/user/edit
func (h *UserHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var req dto.UserUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数错误", err))
		fmt.Println(err)
		return
	}
	if err := h.svc.Update(c.Request.Context(), id, &req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// UpdateStatus 开关账号  /api/user/status/:id/:status
func (h *UserHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	status, _ := strconv.ParseUint(c.Param("status"), 10, 8)
	if err := h.svc.UpdateStatus(c.Request.Context(), id, uint8(status)); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// UpdatePassword 修改密码  /api/user/pwd
type pwdReq struct {
	OldPwd string `json:"oldPwd" binding:"required,min=6"`
	NewPwd string `json:"newPwd" binding:"required,min=6"`
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	userID := c.GetUint64("userID") // 中间件写入
	var req pwdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("参数错误", err))
		return
	}
	// 验证旧密码
	user, err := h.svc.GetByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	if !utils.CheckPassword(req.OldPwd, user.PasswordHash) {
		c.JSON(http.StatusOK, dto.FailMsg("旧密码错误", err))
		return
	}
	if err = h.svc.UpdatePassword(c.Request.Context(), userID, req.NewPwd); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg(err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// Delete 删除用户  /api/user/del/:id
func (h *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("删除失败!", err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil))
}

// Info 当前登录用户详情  /api/user/info
func (h *UserHandler) Info(c *gin.Context) {
	userID := c.GetUint64("userID")
	user, err := h.svc.GetByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusOK, dto.FailMsg("查询失败", err))
		return
	}
	c.JSON(http.StatusOK, dto.Success(user))
}
