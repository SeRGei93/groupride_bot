package utils

import (
	"fmt"
	"time"
)

type AwaitRideStep string

const (
	StepLink        AwaitRideStep = "link"
	StepTime        AwaitRideStep = "time"
	StepDescription AwaitRideStep = "description"
)

var (
	awaitingMessage  = make(map[int64]time.Time)
	awaitingRideStep = make(map[int64]AwaitRideStep)
)

func GetAwaitRideStep(userID int64) (AwaitRideStep, error) {
	step, ok := awaitingRideStep[userID]
	if !ok {
		return "", fmt.Errorf("await ride not found")
	}

	return step, nil
}

func NextAwaitRideStep(userID int64) error {
	step, ok := awaitingRideStep[userID]
	if !ok {
		return fmt.Errorf("await ride not found")
	}

	switch step {
	case StepLink:
		awaitingRideStep[userID] = StepTime
	case StepTime:
		awaitingRideStep[userID] = StepDescription
	}

	return nil
}

// SetAwaiting отмечает пользователя как ожидающего ввода и запоминает текущее время
func SetAwaiting(userID int64, seconds int) {
	awaitingMessage[userID] = time.Now().Add(time.Duration(seconds) * time.Second)
	awaitingRideStep[userID] = StepLink
}

// IsAwaiting проверяет, ожидается ли сообщение от пользователя и не истек ли таймаут
func IsAwaiting(userID int64) bool {
	end, exists := awaitingMessage[userID]
	if !exists {
		return false
	}

	if time.Now().After(end) {
		delete(awaitingMessage, userID)
		delete(awaitingRideStep, userID)
		return false
	}

	return true
}

// CleanupOldAwaiting очищает все записи старше времени ожидания
func CleanupOldAwaiting() {
	for id, t := range awaitingMessage {
		if time.Now().After(t) {
			delete(awaitingMessage, id)
			delete(awaitingRideStep, id)
		}
	}
}

// DeleteAwaiting удаляет запись вручную
func DeleteAwaiting(userID int64) {
	delete(awaitingMessage, userID)
	delete(awaitingRideStep, userID)
}
