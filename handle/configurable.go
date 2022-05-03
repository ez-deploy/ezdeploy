package handle

import (
	"crypto/tls"
	"net/http"

	"github.com/ez-deploy/ezdeploy/restapi/operations"
)

// ConfigurableImpl impl restapi.Configurable interface, but do nothing.
type ConfigurableImpl struct{}

func (c *ConfigurableImpl) ConfigureFlags(api *operations.EzDeployApiserverAPI) {
	return
}

func (c *ConfigurableImpl) ConfigureTLS(tlsConfig *tls.Config) {
	return
}

func (c *ConfigurableImpl) ConfigureServer(s *http.Server, scheme, addr string) {
	return
}

func (c *ConfigurableImpl) CustomConfigure(api *operations.EzDeployApiserverAPI) {
	return
}

func (c *ConfigurableImpl) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func (c *ConfigurableImpl) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
