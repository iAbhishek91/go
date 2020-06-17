# Go Module

* introduced on 1.13.
* default mode will be go modules.
* dependencies management are done almost same as node
  * global repo: [godoc](https://godoc.org/) //similar to npm.org
  * initialize a go module: **go mode init** //similar to npm init
  * metadata file: **go.mod** //similar to package.json or pom.xml
  * auto generated file: **go.sum** //similar to package.lock - should be checked in version control
  * command used: **go test** //similar to npm install
* version of go modules uses semantic versioning like major.minor.patch
* **direct and indirect** modules, direct are which are explicitly called.

## Refer

002-hello-module
002-module-with-dependencies
