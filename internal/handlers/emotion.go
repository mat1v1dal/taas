package handlers

import (
	"net/http"
	"taas/internal/db"
	"taas/internal/models"

	"github.com/gin-gonic/gin"
)

type EmotionInput struct {
	UserID     uint   `json:"user_id" binding:"required"`
	Emotion    string `json:"emotion" binding:"required"`
	Note       string `json:"note"`
	Contextual string `json:"contextual_info"`
	Intensity  int    `json:"intensity" binding:"required"`
}

func PostEmotion(c *gin.Context) {
	var input EmotionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	em := models.Emotion{
		UserID:         input.UserID,
		Mood:           input.Emotion,
		Intensity:      input.Intensity,
		ContextualInfo: input.Contextual,
	}

	if err := db.DB.Create(&em).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la emoción"})
		return
	}

	c.JSON(http.StatusCreated, em)
}

func GetEmotions(c *gin.Context) {
	userID := c.Query("user_id")

	var emotions []models.Emotion
	if err := db.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&emotions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener emociones"})
		return
	}

	c.JSON(http.StatusOK, emotions)
}
