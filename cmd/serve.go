package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"vmware-rest-proxy/internal"
	"vmware-rest-proxy/internal/endpoints"
)

func main() {
	c := internal.Config{
		Resty: resty.New(),
	}

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
		c.Resty.SetBaseURL(b)
	}

	bindAddress := "0.0.0.0:8080"

	if a, found := os.LookupEnv("BIND_ADDRESS"); found {
		bindAddress = a
	}

	if e, found := os.LookupEnv("TLS_INSECURE_SKIP_VERIFY"); found && e == "true" {
		logrus.Warn("Disabling TLS verification")
		c.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if p, found := os.LookupEnv("VCENTER_PROXY_URL"); found {
		logrus.Debug("Setting proxy URL")
		c.Resty.SetProxy(p)
	}

	logrus.Debug("Starting server")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	logrus.Debug("Disabling trusted proxies because it's recommended by gin")
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Error disabling trusted proxies: %s", err)
	}
	e := []endpoints.Endpoint{&endpoints.VMSEndpoint{}, &endpoints.StatusEndpoint{}}
	for _, endpoint := range e {
		endpoint.Register(r, c)
	}
	_ = r.Run(bindAddress)
}
