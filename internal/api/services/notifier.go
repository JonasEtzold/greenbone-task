package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	apimodels "greenbone-task/internal/api/models"
	models "greenbone-task/internal/persistence/models/computer"
	"greenbone-task/internal/persistence/repository"
)

type Notifier interface {
	CheckNotifyAdmin(employeeAbbreviation string)
}

type notifierService struct {
	logger *zap.Logger
}

func NewNotifierService(logger *zap.Logger) Notifier {
	return &notifierService{
		logger: logger,
	}
}

func (ns *notifierService) CheckNotifyAdmin(employeeAbbreviation string) {
	repo := repository.GetComputer()
	var query models.Computer
	query.AssignedEmployee = employeeAbbreviation

	if computers, err := repo.Query(&query); err != nil {
		ns.logger.Error(err.Error())
	} else {
		if len(*computers) >= 3 {
			notification := apimodels.NotifyMessage{
				Level:                apimodels.DefaultNotificationLevel,
				EmployeeAbbreviation: employeeAbbreviation,
				Message:              fmt.Sprintf(`Employee has %d devices assigned.`, len(*computers)),
			}
			rawNotification, err := json.Marshal(notification)
			if err != nil {
				ns.logger.Error("failed to marshal notification message", zap.Error(err))
			}

			notificationUrl := viper.GetString("notification_url")
			_, postErr := http.Post(notificationUrl, "application/json", bytes.NewReader(rawNotification))
			if postErr != nil {
				ns.logger.Error("failed to sent notification", zap.Error(postErr))
				return
			}
			ns.logger.Info("Notification successfully sent", zap.Any("payload", notification))
		}
	}
}
