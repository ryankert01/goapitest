package utils

import (
	"github.com/ryankert01/goapitest/models"

	"time"
)

func IsAdAlive(ad models.AD) bool {
	currentTime := time.Now()
	if currentTime.After(ad.StartAt) && currentTime.Before(ad.EndAt) {
		return true
	}
	return false
}
