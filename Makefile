buildlambdaone:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -gcflags "all=-trimpath=${PWD}" -o ./build/lambdaone/bootstrap ./lambdaone/main.go

buildlambdatwo:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -gcflags "all=-trimpath=${PWD}" -o ./build/lambdatwo/bootstrap ./lambdatwo/main.go

buildcdk:
	cd deployment && go build