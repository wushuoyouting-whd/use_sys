.PHONY: run build swagger clean

# 运行应用
run:
	go run cmd/server/main.go

# 构建应用
build:
	go build -o bin/server cmd/server/main.go

# 生成 Swagger 文档
swagger:
	swag init -g cmd/server/main.go -o docs

# 安装依赖
deps:
	go mod download
	go mod tidy

# 清理
clean:
	rm -rf bin/
	rm -rf docs/

# 安装 Swagger 工具（首次使用）
install-swagger:
	go install github.com/swaggo/swag/cmd/swag@latest

