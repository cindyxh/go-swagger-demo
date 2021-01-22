# GO Demo

### how to code in go
https://golang.org/doc/code.html

### get started
```bash
# create go.mod
go mod init github.com/godemo
# build and install the binary under
# $GOPATH = $HOME/go
# $GOBIN = $GOPATH/bin 
go install github.com/godemo
# run
godemo
# list env vars for go
go env
# export path
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))

# go version
go version

```

### Go mod

- Usage:
    `go mod <command> [arguments]`

- The commands are:
    - download    download modules to local cache
    - edit        edit go.mod from tools or scripts
    - graph       print module requirement graph
    - init        initialize new module in current directory
    - tidy        add missing and remove unused modules
    - vendor      make vendored copy of dependencies
    - verify      verify dependencies have expected content
    - why         explain why packages or modules are needed

### Go tools
    - gopkgs
    - go-outline
    - gotests
    - dlv
    - golint
    - gopls
### Swagger
```bash
# install go-swagger
# https://goswagger.io/install.html
go get -u github.com/go-swagger/go-swagger/cmd/swagger
git config alias.swagger "docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
swagger version
# generate swagger api spec
swagger generate spec -o ./documentation/swagger.json --scan-models
# serve swagger api spec
swagger serve -F=swagger swagger.yaml
# or serve from api server
# http://server/api/
```
### Testing
```
curl http://localhost:5000/api/users/1
```
## Resources

- Swagger UI - cloned from https://github.com/swagger-api/swagger-ui/tree/master/dist
- Go Swagger Syntax & Documentation: https://goswagger.io/
- Operation 
  - https://goswagger.io/use/spec/operation.html
  - https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#operationObject