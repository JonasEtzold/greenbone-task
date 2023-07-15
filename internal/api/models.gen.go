// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

// ComputerDto An object that describes required information about a computer.
type ComputerDto struct {
	// AssignedEmployee The abbrevation of the employee assigned to the computer. The employee abbreviation consists of 3 letters. For example Max Mustermann should be mmu.
	AssignedEmployee *string `json:"assignedEmployee,omitempty"`

	// Description Additional information about the computer.
	Description *string `json:"description,omitempty"`

	// Ip The IP V4 address of the computer within the company network.
	Ip string `json:"ip"`

	// Mac The MAC address of the computer within the company network.
	Mac string `json:"mac"`

	// Name The internal name of the computer.
	Name string `json:"name"`
}

// GetComputersResponse An list of object that describes computers used in the company.
type GetComputersResponse struct {
	// Computers A list of assigned computers.
	Computers []ComputerDto `json:"computers"`
}

// ServiceErrorResponse This schema represents the default response in case of an API call resulting in an error.
type ServiceErrorResponse struct {
	// Code The error code.
	Code string `json:"code"`

	// Details An optional string of explanatatory details about the occured errors.
	Details *string `json:"details,omitempty"`

	// Name The name of the error.
	Name string `json:"name"`
}

// CreateComputerJSONRequestBody defines body for CreateComputer for application/json ContentType.
type CreateComputerJSONRequestBody = ComputerDto

// UpdateComputerJSONRequestBody defines body for UpdateComputer for application/json ContentType.
type UpdateComputerJSONRequestBody = ComputerDto
