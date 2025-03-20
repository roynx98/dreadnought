package adapters

import (
	"adeptus-limitarius/cases"
)

type LimiterControllerRequest struct {
	IP string
}

type LimiterControllerResponse struct {
	Code int
}

type LimiterController struct {
	interactor cases.LimiterInteractor
}

func (controller LimiterController) HandleRequest(LimiterControllerRequest) bool {
	rule := cases.LimitRule{IP: "1.1.1.1", Strategy: "bucket"}

	return controller.interactor.ShouldLimit(rule)
}

func ProvideLimiterController(interactor cases.LimiterInteractor) LimiterController {
	return LimiterController{interactor: interactor}
}
