GO_MOD_NAME=github.com/ez-deploy/ezdeploy

dev:
	@go build -o ezdeploy cmd/ez-deploy-apiserver-server/main.go
	@./ezdeploy --port=8000

	@echo "visit at http://localhost:8888" 

build:
	go build -o ezdeploy cmd/ez-deploy-apiserver-server/main.go

# only used by `make gen_server`, store some models.
pre_gen_server:
	@cp ./models/token.go ./token.go.back
	@cp ./models/user_info.go ./user_info.go.back
	@cp ./models/environment_variable.go ./environment_variable.go.back
	@cp ./models/project_info.go ./project_info.go.back
	@cp ./models/service_info.go ./service_info.go.back
	@cp ./models/service_version.go ./service_version.go.back
	@cp ./models/authority_role.go ./authority_role.go.back
	@cp ./models/role_permission.go ./role_permission.go.back

# only used by `make gen_server`, restore some models.
post_gen_server:
	@mv ./token.go.back ./models/token.go
	@mv ./user_info.go.back ./models/user_info.go
	@mv ./environment_variable.go.back ./models/environment_variable.go
	@mv ./project_info.go.back ./models/project_info.go
	@mv ./service_info.go.back ./models/service_info.go
	@mv ./service_version.go.back ./models/service_version.go
	@mv ./authority_role.go.back ./models/authority_role.go
	@mv ./role_permission.go.back ./models/role_permission.go

gen_server:
	@make pre_gen_server
	@make clean_server

	@swagger generate server \
		--implementation-package=${GO_MOD_NAME}/handle \
		--principal=github.com/ez-deploy/ezdeploy/models.AuthInfo \
		-f ./swagger.yaml
	@go mod tidy

	@make post_gen_server

clean_server:
	@rm -rf cmd models restapi