package adapters

type LimiterControllerRequest struct {
	IP string
}

type LimiterControllerResponse struct {
	Code int
}

type LimiterController struct {
}

func (r LimiterController) HandleRequest(LimiterControllerRequest) LimiterControllerResponse {
	return LimiterControllerResponse{Code: 200}
}
