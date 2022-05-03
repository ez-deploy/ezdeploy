GO_MOD_NAME=github.com/ez-deploy/ezdeploy

build:
	go build -o ezdeploy cmd/ez-deploy-apiserver-server/main.go

gen_server: clean_server
	@swagger generate server \
		--implementation-package=${GO_MOD_NAME}/handle \
		--principal=github.com/ez-deploy/ezdeploy/models.AuthInfo \
		-f ./swagger.yaml
	@go mod tidy

clean_server:
	@rm -rf cmd models restapi