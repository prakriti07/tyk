package openzipkin

import (
	"github.com/prakriti07/tyk/config"
)

// Load retusn a zipkin configuration from the opts.
func Load(opts map[string]interface{}) (*config.ZipkinConfig, error) {
	var c config.ZipkinConfig
	if err := config.DecodeJSON(&c, opts); err != nil {
		return nil, err
	}
	return &c, nil
}
