# Qucik Go CLI

* **go | go help**: displays the CLI option available.
  * **go help <command>**: displays the details of the command.

* **go env**: go related env variables are printed.

* **go fmt <file_name>**: automatically format a code.
  * **go fmt ./...**: format everything in a directory recursively.

* **go run <file_name>**: build and execute a file.

* **go build**: builds the file with the name of the current dir, in the current dir. Assuming it as the project name. It can be single file or  a entire project.

* **go install**: build the file with name installation of the project in GOBIN directory.

* **go mod init**: create a new module, initializing the go.mod file.
* **go mod init example.com/hello-world**: path is mentioned for module. by default it is go/curr_dir_name

* **go mod tidy**: removed unused dependencies.

* **go test**: test the module. also downloads all the used dependencies.

* **go list -m all**: current module dependencies

* **go get golang.org/x/text**: go changes the required version a dependencies.
