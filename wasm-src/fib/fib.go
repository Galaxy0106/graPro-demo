package main

import (
	"fmt"
	"log"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
	bindgen "github.com/second-state/wasmedge-bindgen/host/go"
)

var reqIndex = 35

func main() {
	var conf = wasmedge.NewConfigure(wasmedge.WASI)
	var vm = wasmedge.NewVMWithConfig(conf)
	vm.LoadWasmFile(os.Args[1])
	vm.Validate()
	bg := bindgen.Instantiate(vm)
	var res []interface{}
	res, err := vm.Execute("fib", uint32(reqIndex))
	// res, err := bg.Execute("fib", 20)
	if err != nil {
		log.Fatal("Job Failed!!!")
	}

	fmt.Println(res[0])

	bg.Release()
	vm.Release()
	conf.Release()
}
