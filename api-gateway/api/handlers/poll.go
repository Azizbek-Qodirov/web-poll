package handlers

import (
	pb "auth-service/genprotos"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePoll godoc
// @Summary Create a new poll
// @Description Create a new poll with poll number, title, and options
// @Tags polls
// @Accept json
// @Produce json
// @Param poll body pb.PollCreateReq true "Poll creation request"
// @Success 201 {object} pb.Void
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /poll [post]
// @Security BearerAuth
func (h *HTTPHandler) CreatePoll(c *gin.Context) {
	var req pb.PollCreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	_, err := h.Poll.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &pb.Void{})
}

// UpdatePoll godoc
// @Summary Update an existing poll
// @Description Update an existing poll with poll number, title, and options
// @Tags polls
// @Accept json
// @Produce json
// @Param poll body pb.PollUpdateReq true "Poll update request"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /poll [put]
// @Security BearerAuth
func (h *HTTPHandler) UpdatePoll(c *gin.Context) {
	var req pb.PollUpdateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	_, err := h.Poll.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &pb.Void{})
}

// DeletePoll godoc
// @Summary Delete an existing poll
// @Description Delete a poll by its ID
// @Tags polls
// @Accept json
// @Produce json
// @Param id path string true "Poll ID"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /poll/{id} [delete]
// @Security BearerAuth
func (h *HTTPHandler) DeletePoll(c *gin.Context) {
	id := c.Param("id")
	req := pb.ByID{Id: id}

	_, err := h.Poll.Delete(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &pb.Void{})
}

// GetPollByID godoc
// @Summary Get a poll by ID
// @Description Get a poll by its ID
// @Tags polls
// @Accept json
// @Produce json
// @Param id path string true "Poll ID"
// @Success 200 {object} pb.PollGetByIDRes
// @Failure 404 {object} string "Poll not found"
// @Failure 500 {object} string "Server error"
// @Router /poll/{id} [get]
// @Security BearerAuth
func (h *HTTPHandler) GetPollByID(c *gin.Context) {
	id := c.Param("id")
	req := pb.ByID{Id: id}

	res, err := h.Poll.GetByID(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllPolls godoc
// @Summary Get all polls
// @Description Get all polls with pagination
// @Tags polls
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.PollGetAllRes
// @Failure 500 {object} string "Server error"
// @Router /polls [get]
// @Security BearerAuth
func (h *HTTPHandler) GetAllPolls(c *gin.Context) {
	var req pb.PollGetAllReq
	if limit := c.Query("limit"); limit != "" {
		req.Pagination.Limit = 1
	}
	if offset := c.Query("offset"); offset != "" {
		req.Pagination.Offset = 10
	}

	res, err := h.Poll.GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
