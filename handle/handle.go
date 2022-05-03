package handle

import "github.com/ez-deploy/ezdeploy/handle/user"

// handerImpl impl restapi.Handler interface.
type handlerImpl struct {
	*ConfigurableImpl
	*user.UserOperationImpl
}

func New() *handlerImpl {
	return &handlerImpl{
		ConfigurableImpl:  &ConfigurableImpl{},
		UserOperationImpl: &user.UserOperationImpl{},
	}
}
