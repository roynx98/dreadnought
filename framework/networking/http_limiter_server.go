package networking

import (
	"adeptus-limitarius/adapters"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HttpLimiterServer struct {
	LimiterController adapters.LimiterController
}

func (server HttpLimiterServer) Start(targetHost *url.URL) {
	proxy := httputil.NewSingleHostReverseProxy(targetHost)
	limiterController := adapters.LimiterController{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Host = targetHost.Host

		limiterRequest := adapters.LimiterControllerRequest{IP: r.RemoteAddr}

		shouldLimit := limiterController.HandleRequest(limiterRequest)

		if shouldLimit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		proxy.ServeHTTP(w, r)
	})

	log.Println("Reverse proxy running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ProvideLimiterServer(limiterController adapters.LimiterController) HttpLimiterServer {
	return HttpLimiterServer{LimiterController: limiterController}
}
