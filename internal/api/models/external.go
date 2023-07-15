package models

const DefaultNotificationLevel = "warning"

type NotifyMessage struct {
	Level                string `json:"level"`
	EmployeeAbbreviation string `json:"employeeAbbreviation"`
	Message              string `json:"message"`
}
