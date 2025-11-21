package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"subscribe-test-service/models"
)

// @Summary Create a subscription
// @Description Create a new subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param input body models.SubscriptionCreateRequest true "Subscription info"
// @Success 201 {object} models.SubscriptionResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /subscriptions [post]
func (h *Handler) createSubscription(c *gin.Context) {
	var input models.SubscriptionCreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	startDate, err := time.Parse("01-2006", input.StartDate)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid start_date format. use MM-YYYY")
		return
	}

	var endDate *time.Time
	if input.EndDate != "" {
		parsedEndDate, err := time.Parse("01-2006", input.EndDate)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid end_date format. use MM-YYYY")
			return
		}
		endDate = &parsedEndDate
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	sub := models.Subscription{
		UserID:      userID,
		ServiceName: input.ServiceName,
		Price:       input.Price,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	subID, err := h.services.Subscription.Create(sub)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	createdSub, err := h.services.Subscription.GetByID(uint(subID))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdSub)
}

// @Summary Get a subscription by ID
// @Description Get a subscription by its ID
// @Tags subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} models.SubscriptionResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /subscriptions/{id} [get]
func (h *Handler) getSubscription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	sub, err := h.services.Subscription.GetByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sub)
}

// @Summary Get all subscriptions
// @Description Get a list of all subscriptions
// @Tags subscriptions
// @Produce json
// @Param user_id query string false "User ID"
// @Param service_name query string false "Service Name"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {array} models.SubscriptionResponse
// @Failure 500 {object} errorResponse
// @Router /subscriptions [get]
func (h *Handler) getAllSubscriptions(c *gin.Context) {
	var filter models.GetAllSubscriptionsFilter

	if userIDStr := c.Query("user_id"); userIDStr != "" {
		userID, err := uuid.Parse(userIDStr)
		if err == nil {
			filter.UserID = &userID
		}
	}
	if serviceName := c.Query("service_name"); serviceName != "" {
		filter.ServiceName = &serviceName
	}
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		startDate, err := time.Parse(time.RFC3339, startDateStr)
		if err == nil {
			filter.StartDate = &startDate
		}
	}
	if endDateStr := c.Query("end_date"); endDateStr != "" {
		endDate, err := time.Parse(time.RFC3339, endDateStr)
		if err == nil {
			filter.EndDate = &endDate
		}
	}

	subs, err := h.services.Subscription.GetAll(filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, subs)
}

// @Summary Update a subscription
// @Description Update a subscription's details
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param input body models.SubscriptionUpdateRequest true "Subscription update info"
// @Success 200 {object} models.SubscriptionResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /subscriptions/{id} [put]
func (h *Handler) updateSubscription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input models.SubscriptionUpdateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	sub, err := h.services.Subscription.GetByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "subscription not found")
		return
	}

	if input.ServiceName != nil {
		sub.ServiceName = *input.ServiceName
	}
	if input.Price != nil {
		sub.Price = *input.Price
	}
	if input.StartDate != nil {
		startDate, err := time.Parse("01-2006", *input.StartDate)
		if err == nil {
			sub.StartDate = startDate
		}
	}
	if input.EndDate != nil {
		endDate, err := time.Parse("01-2006", *input.EndDate)
		if err == nil {
			sub.EndDate = &endDate
		}
	}

	if err := h.services.Subscription.Update(sub); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	updatedSub, err := h.services.Subscription.GetByID(uint(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedSub)
}

// @Summary Delete a subscription
// @Description Delete a subscription by its ID
// @Tags subscriptions
// @Param id path int true "Subscription ID"
// @Success 204 "No Content"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /subscriptions/{id} [delete]
func (h *Handler) deleteSubscription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	if err := h.services.Subscription.Delete(uint(id)); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
