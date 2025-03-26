package networking

import (
	"dreadnought/adapters"
	"dreadnought/framework/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HttpLimiterServer struct {
	limiterController adapters.LimiterController
	configManager     config.ConfigManager
}

func (server HttpLimiterServer) Start() {
	target, _ := url.Parse(server.configManager.Config.TargetHost)

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Host = target.Host

		limiterRequest := adapters.LimiterControllerRequest{IP: r.RemoteAddr}

		shouldLimit := server.limiterController.HandleRequest(limiterRequest)

		if shouldLimit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		proxy.ServeHTTP(w, r)
	})
	port := server.configManager.Config.Port

	log.Println("Reverse proxy running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ProvideLimiterServer(configManager config.ConfigManager, limiterController adapters.LimiterController) HttpLimiterServer {
	return HttpLimiterServer{configManager: configManager, limiterController: limiterController}
}
