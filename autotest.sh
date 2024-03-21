export GOARCH=amd64

# 执行单元测试
go test -gcflags=-l -v -covermode=set -coverprofile=.coverage.cov ./linear

# 生成测试覆盖率的 html 文件
go tool cover -html=.coverage.cov -o coverage.html
