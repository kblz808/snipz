package handlers

import (
	"net/http"
	"snipz/internal/http/middlewares"
	"snipz/internal/services"
	"snipz/internal/storage/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type SnippetHandler struct {
	service services.SnippetService
}

func NewSnippetHandler(service services.SnippetService) *SnippetHandler {
	return &SnippetHandler{service}
}

type createSnippetRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Language string `json:"language"`

	ExpiresAt   time.Time `json:"expires_at"`
	Visibility  string    `json:"visibility"`
	IsEncrypted bool      `json:"is_encrypted"`
	Password    string    `json:"password"`
	Tags        []string  `json:"tags"`
}

func (handler *SnippetHandler) CreateSnippet(ctx *gin.Context) {
	var req createSnippetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	payload, exists := ctx.Get(middlewares.AuthorizationPayloadKey)
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "authorization payload not found"})
		return
	}
	authPayload := payload.(*services.TokenPayload)

	// convert req into repository.Snippet
	snippet := &repository.Snippet{
		Title:       req.Title,
		Content:     req.Content,
		Language:    req.Language,
		ExpiresAt:   req.ExpiresAt,
		Visibility:  repository.Visibility(req.Visibility),
		IsEncrypted: req.IsEncrypted,
		Password:    req.Password,
		Tags:        req.Tags,
	}

	snippet, err := handler.service.CreateSnippet(ctx, snippet, authPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"snippet": snippet})
}

type getSnippetRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (handler *SnippetHandler) GetSnippet(ctx *gin.Context) {
	var req getSnippetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	snippet, err := handler.service.GetSnippet(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"snippet": snippet})
}

type listSnippetsRequest struct {
	Skip  uint64 `form:"skip"`
	Limit uint64 `form:"limit"`
}

func (handler *SnippetHandler) ListSnippets(ctx *gin.Context) {
	var req listSnippetsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	snippets, err := handler.service.GetSnippets(ctx, req.Skip, req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"snippets": snippets})
}

type searchSnippetRequrst struct {
	Title    string `form:"title"`
	Language string `form:"language"`
	Skip     uint64 `form:"skip"`
	Limit    uint64 `form:"limit"`
}

func (handler *SnippetHandler) SearchSnippets(ctx *gin.Context) {
	var req searchSnippetRequrst
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
}
