package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"net/url"
	"os"
	"vmware-rest-proxy/cmd/docs"
	api2 "vmware-rest-proxy/internal/api"
	"vmware-rest-proxy/internal/endpoints"
)

// REST server that proxies request through to a vCenter web service making it easier to request certain details.
// @title vmware-rest-proxy
// @version 0.1.0
// @contact.name   DO!DevOps
// @contact.url    http://dodevops.io
// @contact.email  info@dodevops.io
//
// @securityDefinitions.basic  BasicAuth
func main() {
	r := resty.New()

	if l, found := os.LookupEnv("LOG_LEVEL"); found {
		if lv, err := logrus.ParseLevel(l); err != nil {
			log.Fatalf("Can not parse log level %s", l)
		} else {
			logrus.SetLevel(lv)
		}
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	logrus.Debug("Checking configuration")

	if b, found := os.LookupEnv("BASE_URL"); !found {
		log.Fatal("Please set BASE_URL to the base url of the vCenter you'd like to access.")
	} else {
		r.SetBaseURL(b)
	}

	bindAddress := "0.0.0.0:8080"

	if a, found := os.LookupEnv("BIND_ADDRESS"); found {
		bindAddress = a
	}

	if e, found := os.LookupEnv("TLS_INSECURE_SKIP_VERIFY"); found && e == "true" {
		logrus.Warn("Disabling TLS verification")
		r.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if p, found := os.LookupEnv("VCENTER_PROXY_URL"); found && p != "" {
		logrus.Debug("Setting proxy URL")
		r.SetProxy(p)
	}

	externalBaseUrl, _ := url.Parse("http://localhost:8080")
	if p, found := os.LookupEnv("EXTERNAL_BASE_URL"); found && p != "" {
		logrus.Debug("Setting external base URL")
		if b, err := url.Parse(p); err != nil {
			log.Fatalf("Can not parse external base url %s: %s", p, err)
		} else {
			externalBaseUrl = b
		}
	}

	logrus.Debug("Starting server")

	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	logrus.Debug("Disabling trusted proxies because it's recommended by gin")
	if err := s.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Error disabling trusted proxies: %s", err)
	}
	api := api2.DefaultVSphereProxyApi{Resty: r}
	for _, endpoint := range []endpoints.Endpoint{&endpoints.VMSEndpoint{API: api}, &endpoints.HostsEndpoint{API: api}, &endpoints.StatusEndpoint{}, &endpoints.DataStoreEndpoint{API: api}} {
		endpoint.Register(s)
	}
	docs.SwaggerInfo.Title = "vmware-rest-proxy"
	docs.SwaggerInfo.Description = fmt.Sprintf("This is the API of the vmware-rest-proxy pointing at %s", r.BaseURL)
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = externalBaseUrl.Host
	docs.SwaggerInfo.Schemes = []string{externalBaseUrl.Scheme}
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	s.GET("/swagger", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	_ = s.Run(bindAddress)
}
