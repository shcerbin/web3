package main

import (
	"github.com/chenzhijie/go-web3"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Get("console").Call("log", "Web worker started")

	web3.NewWeb3("https://rpc.flashbots.net")
}
