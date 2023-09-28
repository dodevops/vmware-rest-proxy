package internal

import (
	"github.com/go-resty/resty/v2"
)

// Config holds shared configuration data
type Config struct {
	// The prepared resty client to use
	Resty *resty.Client
}
