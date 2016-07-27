package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
  "encoding/json"
)

const FILEPATH = "/mnt/data/goapp.txt"



type VcapServices struct {
	serviceName []struct {
		Credentials struct {
			Database string `json:"database"`
			Host string `json:"host"`
			Password string `json:"password"`
			Port int `json:"port"`
			URI string `json:"uri"`
			Username string `json:"username"`
		} `json:"credentials"`
		SyslogDrainURL interface{} `json:"syslog_drain_url"`
		VolumeMounts []struct {
			ContainerPath string `json:"container_path"`
			Mode string `json:"mode"`
		} `json:"volume_mounts"`
		Label string `json:"label"`
		Provider interface{} `json:"provider"`
		Plan string `json:"plan"`
		Name string `json:"name"`
		Tags []interface{} `json:"tags"`
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	//serviceInstance := os.Getenv("CF_SERVICE_NUM")
	fmt.Fprintln(w, "hello world")
	vcap_services := &VcapServices{}
	volumePath := vcapServices["scaleio-service-broker-vf"].([]interface{})[0].(map[string]interface{})["volume_mounts"].([]interface{})[0].(map[string]interface{})["container_path"]
  // err := json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &vcap_services)
	// if err != nil {
	// 	fmt.Fprintln(w, fmt.Sprintf("Life is wrong and the unmarshal failed! %s", err))
	// }
  fmt.Fprintln(w, volumePath)//vcap_services.serviceName[0].VolumeMounts[0].ContainerPath)
	fmt.Fprintln(w, "VCAP_SERVICES:", os.Getenv("VCAP_SERVICES"))
}

func read(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile(FILEPATH)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "file content: %s", string(content))
}

func write(w http.ResponseWriter, r *http.Request) {
	s := "Chris and I were here"
	err := ioutil.WriteFile(FILEPATH, []byte(s), os.ModePerm)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "saved %s to %s successfully", s, FILEPATH)
}

func main() {
	http.HandleFunc("/read", read)
	http.HandleFunc("/write", write)
	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
