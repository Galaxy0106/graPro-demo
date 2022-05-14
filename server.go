package main

import (
	"encoding/json"
	"graPro-demo/producer"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path"

	"github.com/suborbital/vektor/vk"
)

type deployProcessJson struct {
	UserName    string `json:"UserName"`
	ProcessName string `json:"ProcessName"`
}

type invokeProcessJson struct {
	UserName    string `json:"UserName"`
	ProcessName string `json:"ProcessName"`
	WasmName    string `json:"WasmName"`
	Period      int    `json:"Period"`
	Quota       int    `json:"Quota"`
	Ns          int    `json:"Ns"`
}

// Deploy func
func HandleDeploy(r *http.Request, ctx *vk.Ctx) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	// Compile the src code
	var resp deployProcessJson
	err := json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}
	executable := path.Join(wasm_app_dir, resp.ProcessName, resp.ProcessName+".go")
	// Executable file is located at the same dir as main.go.
	cmd := exec.Command("go", "build", executable)

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	return executable, nil
}

func HandleInvoke(r *http.Request, ctx *vk.Ctx) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	var resp invokeProcessJson
	err := json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}
	// Start client
	client := producer.StartClient()
	// Create Task from post body
	task, _ := producer.NewTask(resp.ProcessName, resp.WasmName, resp.Period, resp.Quota, resp.Ns)
	// Task Enqueue
	producer.EnQueue(client, task)
	// Close Client
	producer.CloseClient(client)
	return "Enqueue successfully", nil
}

func NewServer() {
	server := vk.New(
		vk.UseAppName("Vektor API HTTP-only"),
		vk.UseHTTPPort(8000),
	)

	asyncGroup := vk.Group("/asyncFunc")
	asyncGroup.POST("/deploy", HandleDeploy)
	asyncGroup.POST("/invoke", HandleInvoke)
	server.AddGroup(asyncGroup)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
