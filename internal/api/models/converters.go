package models

import (
	models "greenbone-task/internal/persistence/models/computer"
)

func DbToApiComputer(dbc models.Computer) ComputerDto {
	return ComputerDto{
		AssignedEmployee: &dbc.AssignedEmployee,
		Description:      &dbc.Description,
		Ip:               dbc.Ip,
		Mac:              dbc.Mac,
		Name:             dbc.Name,
	}
}
