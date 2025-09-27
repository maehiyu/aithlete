package handler

import (
	"api/application/dto"
	"api/application/service/command"
	"api/application/service/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppointmentHandler は予約関連のHTTPリクエストを処理します。
type AppointmentHandler struct {
	commandService *command.AppointmentCommandService
	queryService   *query.AppointmentQueryService
}

// NewAppointmentHandler はAppointmentHandlerの新しいインスタンスを生成します。
func NewAppointmentHandler(cs *command.AppointmentCommandService, qs *query.AppointmentQueryService) *AppointmentHandler {
	return &AppointmentHandler{commandService: cs, queryService: qs}
}

// HandleCreateAppointment は新しい予約を作成するハンドラです。
func (h *AppointmentHandler) HandleCreateAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var req dto.AppointmentCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: 本来は認証情報からUserIDを取得すべき
		// userID, _ := c.Get("userId")
		// req.UserID = userID.(string)

		res, err := h.commandService.CreateAppointment(ctx, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, res)
	}
}

// HandleGetAppointmentByID はIDで予約を取得するハンドラです。
func (h *AppointmentHandler) HandleGetAppointmentByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id := c.Param("id")
		res, err := h.queryService.GetByID(ctx, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if res == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

// HandleListAppointments は予約のリストを取得するハンドラです。
// クエリパラメータで絞り込みます (例: /appointments?user_id=...)
func (h *AppointmentHandler) HandleListAppointments() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userID := c.Query("user_id")
		coachID := c.Query("coach_id")
		chatID := c.Query("chat_id")

		var appointments []*dto.AppointmentResponse
		var err error

		if userID != "" {
			appointments, err = h.queryService.GetByUserID(ctx, userID)
		} else if coachID != "" {
			appointments, err = h.queryService.GetByCoachID(ctx, coachID)
		} else if chatID != "" {
			appointments, err = h.queryService.GetByChatID(ctx, chatID)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user_id, coach_id, or chat_id is required"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, appointments)
	}
}

// HandleUpdateAppointment は予約を更新するハンドラです。
func (h *AppointmentHandler) HandleUpdateAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id := c.Param("id")
		var req dto.AppointmentUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := h.commandService.UpdateAppointment(ctx, id, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// HandleDeleteAppointment は予約を削除するハンドラです。
func (h *AppointmentHandler) HandleDeleteAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id := c.Param("id")
		if err := h.commandService.DeleteAppointment(ctx, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
