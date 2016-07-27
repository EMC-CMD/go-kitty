package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
)


type VcapServices struct {
	EMCPersistenceCI []struct {
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
	} `json:"EMC-Persistence-CI"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "hello world")

	vcap_services := &VcapServices{}
	err := json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &vcap_services)
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("Life is wrong and the unmarshal failed! %s", err))
	}
	volumePath := vcap_services.EMCPersistenceCI[0].VolumeMounts[0].ContainerPath
  fmt.Fprintln(w, fmt.Sprintf("%s",volumePath))
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/test.txt", volumePath))
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, fmt.Sprintf("file content: %s", string(content)))
}

func write(w http.ResponseWriter, r *http.Request) {
	s := "Chris and I were here \n"
	vcap_services := &VcapServices{}
	err := json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &vcap_services)
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("Life is wrong and the unmarshal failed! %s", err))
	}
	volumePath := vcap_services.EMCPersistenceCI[0].VolumeMounts[0].ContainerPath
	f, err := os.OpenFile(fmt.Sprintf("%s/test.txt", volumePath), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
	    panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(s); err != nil {
	    panic(err)
	}
	fmt.Fprintln(w, fmt.Sprintf("saved %s to %s successfully", s, volumePath))
}

func main() {
	http.HandleFunc("/write", write)
	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
