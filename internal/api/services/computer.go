package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	apimodels "greenbone-task/internal/api/models"
	dbmodels "greenbone-task/internal/persistence/models/computer"
	"greenbone-task/internal/persistence/repository"
)

type computerService struct {
	logger *zap.Logger
	notify Notifier
}

func NewComputerService(logger *zap.Logger, notify Notifier) ServerInterface {
	return &computerService{
		logger: logger,
		notify: notify,
	}
}
func (cs *computerService) GetComputers(c *gin.Context) {
	repo := repository.GetComputer()

	if computers, err := repo.All(); err != nil {
		cs.newError(c, http.StatusNotFound, errors.New("computer not found"))
		cs.logger.Error(err.Error())
	} else {
		var respComputers = make([]apimodels.ComputerDto, len(*computers))
		for i, computer := range *computers {
			respComputers[i] = apimodels.DbToApiComputer(computer)
		}
		resp := apimodels.GetComputersResponse{Computers: respComputers}
		c.JSON(http.StatusOK, resp)
	}
}

func (cs *computerService) CreateComputer(c *gin.Context) {
	repo := repository.GetComputer()
	var computerInput dbmodels.Computer
	if err := c.BindJSON(&computerInput); err != nil {
		cs.newError(c, http.StatusBadRequest, err)
		return
	}
	if err := repo.Add(&computerInput); err != nil {
		cs.newError(c, http.StatusBadRequest, err)
		cs.logger.Error(err.Error())
	} else {
		cs.logger.Info("saved data", zap.Any("computerData", computerInput))
		c.JSON(http.StatusCreated, apimodels.DbToApiComputer(computerInput))
	}

	cs.notify.CheckNotifyAdmin(computerInput.AssignedEmployee)
}

func (cs *computerService) GetComputer(c *gin.Context, computerName string) {
	repo := repository.GetComputer()
	if computer, err := repo.Get(computerName); err != nil {
		cs.newError(c, http.StatusNotFound, errors.New("computer not found"))
		cs.logger.Error(err.Error())
	} else {
		c.JSON(http.StatusOK, apimodels.DbToApiComputer(*computer))
	}
}

func (cs *computerService) UpdateComputer(c *gin.Context, computerName string) {
	repo := repository.GetComputer()

	storedComputer, err := repo.Get(computerName)
	if err != nil {
		cs.newError(c, http.StatusNotFound, errors.New("computer not found"))
		cs.logger.Error(err.Error())
		return
	}

	if err := c.BindJSON(&storedComputer); err != nil {
		abortErr := c.AbortWithError(http.StatusBadRequest, err)
		if abortErr != nil {
			cs.logger.Error(abortErr.Error())
		}
		return
	}

	if err := repo.Update(storedComputer); err != nil {
		cs.newError(c, http.StatusNotFound, err)
		cs.logger.Error(err.Error())
	} else {
		c.JSON(http.StatusOK, apimodels.DbToApiComputer(*storedComputer))
	}

	cs.notify.CheckNotifyAdmin(storedComputer.AssignedEmployee)
}

func (cs *computerService) DeleteComputer(c *gin.Context, computerName string) {
	s := repository.GetComputer()
	if computer, err := s.Get(computerName); err != nil {
		cs.newError(c, http.StatusNotFound, errors.New("computer not found"))
		cs.logger.Error(err.Error())
	} else {
		if err := s.Delete(computer); err != nil {
			cs.newError(c, http.StatusNotFound, err)
			cs.logger.Error(err.Error())
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}

func (cs *computerService) GetComputersByEmployee(c *gin.Context, assignedEmployee string) {
	repo := repository.GetComputer()
	var query dbmodels.Computer
	query.AssignedEmployee = assignedEmployee
	if computers, err := repo.Query(&query); err != nil {
		cs.newError(c, http.StatusNotFound, errors.New("no computers found"))
		cs.logger.Error(err.Error())
	} else {
		var respComputers = make([]apimodels.ComputerDto, len(*computers))
		for i, computer := range *computers {
			respComputers[i] = apimodels.DbToApiComputer(computer)
		}
		resp := apimodels.GetComputersResponse{Computers: respComputers}
		c.JSON(http.StatusOK, resp)
	}
}

func (*computerService) newError(c *gin.Context, status int, err error) {
	er := apimodels.ServiceErrorResponse{
		Code: status,
		Name: err.Error(),
	}
	c.JSON(http.StatusNotFound, er)
}
