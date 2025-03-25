package adapters

import (
	"dreadnought/cases"
	"dreadnought/entities"
)

type LimiterControllerRequest struct {
	IP string
}

type LimiterController struct {
	interactor cases.LimiterInteractor
}

func (controller LimiterController) HandleRequest(request LimiterControllerRequest) bool {
	rule := entities.LimitRule{IP: request.IP, Strategy: "bucket"}

	return controller.interactor.ShouldLimit(rule)
}

func ProvideLimiterController(interactor cases.LimiterInteractor) LimiterController {
	return LimiterController{interactor: interactor}
}
