package handle

import (
	"crypto/tls"
	"net/http"

	"github.com/ez-deploy/ezdeploy/restapi/operations"
)

// ConfigurableImpl impl restapi.Configurable interface, but do nothing.
type ConfigurableImpl struct{}

func (c *ConfigurableImpl) ConfigureFlags(api *operations.EzDeployApiserverAPI) {}

func (c *ConfigurableImpl) ConfigureTLS(tlsConfig *tls.Config) {}

func (c *ConfigurableImpl) ConfigureServer(s *http.Server, scheme, addr string) {}

func (c *ConfigurableImpl) CustomConfigure(api *operations.EzDeployApiserverAPI) {}

func (c *ConfigurableImpl) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func (c *ConfigurableImpl) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
