package client
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/keir-rex/lifx-cli/cmd/lifx/config"
	"net/http"
	"os"
)

const base_path = "https://api.lifx.com/v1"

type Product struct {
	Name            string `json:"name"`
	Identifier      string `json:"identifier"`
	LastSeen        string `json:"last_seen"`
	SecondSinceSeen string `json:'second_since_seen"`
}

type Light struct {
	Id      string  `json:"id"`
	Product Product `json:"product"`
}

type Lights []Light

var configuration map[string]string

func init() {
	configuration, err := config.Get()

	if err != nil || configuration["token"] == "" {
		config.InitializeConfig()
		fmt.Println("Set your token in ~/.lifx/conf.json")
		os.Exit(1)
	}
}

func Verbose(debug bool, body []byte, err error) {
	if err != nil {
		panic(err)
	}
	if debug {
		fmt.Println(string("Testing"))
		fmt.Println(string(body))
	}
}

func Config() {
	configuration, _ = config.Get()
	fmt.Printf("Configuration:\n\n")
	for k, v := range configuration {
		fmt.Printf(k)
		fmt.Println(v)
	}
	fmt.Printf("\n")
}

func SelectLight(debug bool, selector string) {
	configuration, _ = config.Get()
	
	if selector != "" {
		configuration = config.Set("selector", selector)
	}
	fmt.Printf("Selected:\t%s\n", configuration["selector"])
}

func List(debug bool) {
	configuration, err := config.Get()
	if err != nil || configuration["token"] == "" {
		config.InitializeConfig()
		fmt.Println("Set your token in ~/.lifx/conf.json")
		os.Exit(1)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", base_path+"/lights/all", nil)
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var lights Lights

	json.Unmarshal(body, &lights)

	for i := 0; i < len(lights); i++ {
		fmt.Printf("Id:\t%s\n", lights[i].Id)
		fmt.Printf("Name:\t%s\n", lights[i].Product.Name)
		fmt.Printf("\n")
	}
}

func Toggle(debug bool) {
	configuration, _ = config.Get()
	client := &http.Client{}
	s := configuration["selector"]
	req, err := http.NewRequest("POST", base_path+"/lights/"+s+"/toggle", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}

func On(debug bool) {
	configuration, _ = config.Get()
	client := &http.Client{}
	s := configuration["selector"]
	payload := bytes.NewBuffer([]byte(`{"power":"on"}`))
	req, err := http.NewRequest("PUT", base_path+"/lights/"+s+"/state", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}

func Off(debug bool) {
	configuration, _ = config.Get()
	client := &http.Client{}
	s := configuration["selector"]
	payload := bytes.NewBuffer([]byte(`{"power":"off"}`))
	req, err := http.NewRequest("PUT", base_path+"/lights/"+s+"/state", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}

func Set(selector string, state string) {
	SelectLight(true, selector)
	if state == "on" {
		fmt.Println("Light on")
		On(true)
		return
	}
	if state == "off" {
		fmt.Println("Light off")
		Off(true)
		return
	}
	fmt.Println("Selected invalidate light state")
}

func Brightness(debug bool, b string) {
	configuration, _ = config.Get()
	client := &http.Client{}
	s := configuration["selector"]
	payload := bytes.NewBuffer([]byte(`{"brightness":` + b + `}`))
	req, err := http.NewRequest("PUT", base_path+"/lights/"+s+"/state", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}

func Color(debug bool, c string) {
	configuration, _ = config.Get()
	client := &http.Client{}
	s := configuration["selector"]
	payload := bytes.NewBuffer([]byte(`{"color":"` + c + `"}`))
	req, err := http.NewRequest("PUT", base_path+"/lights/"+s+"/state", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}
