build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/lambdaenvs-linux-amd64
	GOARCH=amd64 GOOS=windows go build -o ./bin/lambdaenvs-windows-amd64
	GOARCH=amd64 GOOS=darwin go build -o ./bin/lambdaenvs-darwin-amd64
	GOARCH=arm64 GOOS=darwin go build -o ./bin/lambdaenvs-darwin-arm64