
package:
	@if [ "$$OPENAPI_TOKEN" = "" ]; then echo "OPENAPI_TOKEN is not set" && exit 1; fi
	mkdir -p dist
	GOOS=linux arch=amd64 go build -o dist/chatgptsvr \
		-ldflags "-X main.openaiToken=$$OPENAPI_TOKEN " \
		*.go
