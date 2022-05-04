GO_MOD_NAME=github.com/ez-deploy/ezdeploy

dev:
	@go build -o ezdeploy cmd/ez-deploy-apiserver-server/main.go
	@./ezdeploy --port=8000

	@echo "visit at http://localhost:8888" 

build:
	go build -o ezdeploy cmd/ez-deploy-apiserver-server/main.go

gen_server:
	@cp ./models/token.go ./token.go.back
	@cp ./models/user_info.go ./user_info.go.back

	make clean_server

	@swagger generate server \
		--implementation-package=${GO_MOD_NAME}/handle \
		--principal=github.com/ez-deploy/ezdeploy/models.AuthInfo \
		-f ./swagger.yaml
	@go mod tidy

	@mv ./user_info.go.back ./models/user_info.go
	@mv ./token.go.back ./models/token.go

clean_server:
	@rm -rf cmd models restapi