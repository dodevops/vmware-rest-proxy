# VMware REST proxy

REST server that proxies request through to a vCenter web service making it easier to request certain details.

# Usage

Start the server by running

    go run cmd/serve.go

# Configuration

The following environment variables are used for configuration:

* BASE_URL (required): The base URL of the vCenter to connect to like https://vcenter.company.com
* BIND_ADDRESS: Bind address to bind the server to [0.0.0.0:8080]
* LOG_LEVEL: Maximum log level to use (see (https://pkg.go.dev/github.com/sirupsen/logrus#readme-level-logging)) [INFO]
* TLS_INSECURE_SKIP_VERIFY: If set, will disable TLS verification for the API client
* VCENTER_PROXY_URL: Connect to the vCenter using this proxy
* EXTERNAL_BASE_URL: Base URL the service is hosted on. Will be used for the Swagger API docs

# APIs

The service includes a generated OpenAPI documentation available at `/swagger`.

See [the generated API documentation](api.md)