package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
)


type VcapServices struct {
	ScaleioServiceBrokerVf []struct {
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
	} `json:"scaleio-service-broker-vf"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello world")

	vcap_services := &VcapServices{}
	err := json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &vcap_services)
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("Life is wrong and the unmarshal failed! %s", err))
	}
	volumePath := vcap_services.ScaleioServiceBrokerVf[0].VolumeMounts[0].ContainerPath
  fmt.Fprintln(w, fmt.Sprintf("%s",volumePath))
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/test.txt", volumePath))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "file content: %s", string(content))
}

func write(w http.ResponseWriter, r *http.Request) {
	s := "Chris and I were here"
	vcap_services := &VcapServices{}
	err := json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &vcap_services)
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("Life is wrong and the unmarshal failed! %s", err))
	}
	volumePath := vcap_services.ScaleioServiceBrokerVf[0].VolumeMounts[0].ContainerPath
	err = ioutil.WriteFile(fmt.Sprintf("%s/test.txt", volumePath), []byte(s), os.ModePerm)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "saved %s to %s successfully", s, volumePath)
}

func main() {
	http.HandleFunc("/write", write)
	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
