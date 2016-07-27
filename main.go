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
  fmt.Fprintln(w, fmt.Sprintf("CF IP: %s",os.Getenv("CF_INSTANCE_ADDR")))
	fmt.Fprintln(w, fmt.Sprintf("CF Instance Number: %s",os.Getenv("CF_INSTANCE_INDEX")))

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
	instanceIP := os.Getenv("CF_INSTANCE_ADDR")
	instanceIndex := os.Getenv("CF_INSTANCE_INDEX")
	fmt.Fprintln(w, fmt.Sprintf("CF IP: %s",instanceIP))
	fmt.Fprintln(w, fmt.Sprintf("CF Instance Number: %s",instanceIndex))

	s := fmt.Sprintf("BOT was here and wrote from instanceID %s, with IP %s! \n", instanceIndex, instanceIP)
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
