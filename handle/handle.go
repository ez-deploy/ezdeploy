package handle

import (
	"log"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/handle/project"
	"github.com/ez-deploy/ezdeploy/handle/user"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "github.com/go-sql-driver/mysql"
)

// handerImpl impl restapi.Handler interface.
type handlerImpl struct {
	*ConfigurableImpl
	*user.UserOperationImpl
	*project.ProjectOperationImpl
}

const dsn = "kratos:123456@tcp(localhost:3306)/ezdeploy?charset=utf8mb4&parseTime=True"
const kubeconfigPath = "/home/wangsaiyu/.kube/config"

func New() *handlerImpl {
	tables, err := db.NewTables(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// create kubernetes client.
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return &handlerImpl{
		ConfigurableImpl:     &ConfigurableImpl{},
		UserOperationImpl:    &user.UserOperationImpl{Tables: tables},
		ProjectOperationImpl: &project.ProjectOperationImpl{Tables: tables, ClientSet: clientset},
	}
}
