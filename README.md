This project was written to demonstrate how to run a golang web3 implementation in a web worker

Required:
- golang
- node

At the moment, the gopherjs does not work with 1.18:
- download module                     `go get golang.org/dl/go1.17.1`
- open module folder in console       `cd ~/go/pkg/mod/golang.org/dl@v0.0.0-20220315170520-faa7218da89a/go1.17.1`
- install module in bin folder        `go install .`
- install source-map-support module   `npm install --global source-map-support`
- compile golang into a js web worker `GOPHERJS_GOROOT="$(go1.17.1 env GOROOT)" gopherjs build main.go`
- serve demo on any server            `GOPHERJS_GOROOT="$(go1.17.1 env GOROOT)" gopherjs serve`

At the moment, the gopherjs does not support cgo
