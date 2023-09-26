package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"vmware-rest-proxy/internal"
	"vmware-rest-proxy/internal/endpoints"
)

func main() {
	c := internal.Config{}

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
		c.BaseUrl = b
	}

	if a, found := os.LookupEnv("BIND_ADDRESS"); !found {
		c.BindAddress = "0.0.0.0:8080"
		logrus.Info("BIND_ADDRESS not specified, using 0.0.0.0:8080 as the bind address.")
	} else {
		c.BindAddress = a
	}

	logrus.Debug("Starting server")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	logrus.Debug("Disabling trusted proxies because it's recommended by gin")
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Error disabling trusted proxies: %s", err)
	}
	e := []endpoints.Endpoint{&endpoints.VMSEndpoint{}}
	for _, endpoint := range e {
		endpoint.Register(r, c)
	}
	_ = r.Run(c.BindAddress)
}
