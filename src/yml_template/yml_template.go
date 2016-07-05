package yml_template

import (
//    "fmt"
//    "html"
//    "log"
    "net/http"
    "errors"
    "os"
    "text/template"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "path/filepath"
//    "github.com/gorilla/mux"
)

// tpl_url := "https://raw.githubusercontent.com/k8sp/cloud-configs/liangjiameng/src/yml_template/cloud-config.tpl"
// config_url := "https://raw.githubusercontent.com/k8sp/cloud-configs/liangjiameng/src/yml_template/cluster.conf"

type Config struct {
    Nodes map[string]Options
}

type Options struct {
    Own_ip string
    Etcd2 string
    Hostname string
    Nic_name string
}

func main1() {
    tplFile, _ := filepath.Abs("./cloud-config.tpl")
    configFile, _ := filepath.Abs("./cluster.conf")
    configTpl, err := ioutil.ReadFile(tplFile)
    Check(err)
    clusterConfig, err := ioutil.ReadFile(configFile)
    Check(err)

    BatchCreate(string(configTpl), clusterConfig)
//  CreateCloudConfig(string(configTpl), clusterConfig, "00:25:90:c0:f6:ee")
}

func BatchCreate(tmpl string, data []byte) {
    var config Config
    err := yaml.Unmarshal(data, &config)
    Check(err)
//  fmt.Printf("Nodes: %#+v\n", config.Nodes["00:25:90:c0:f6:ee"])
    t := template.Must(template.New("configTpl").Parse(tmpl))

    for k, v := range config.Nodes {
        f, err := os.Create(k + ".yml")
        Check(err)
        err = t.Execute(f, v)
        Check(err)
//      fmtt.Printf("keys:%s\tvalues:%s\n", k, v.Own_ip)
    }

}

func CreateCloudConfig(tmpl string, data []byte, mac string) {
    if mac == "" {
         BatchCreate(tmpl, data)
         return
    }
    var config Config
    err := yaml.Unmarshal(data, &config)
    Check(err)
//  fmt.Printf("Nodes: %#+v\n", config.Nodes["00:25:90:c0:f6:ee"])
    t := template.Must(template.New("configTpl").Parse(tmpl))
    if v, ok := config.Nodes[mac]; ok {
        //do something here
        f, err := os.Create(mac + ".yml")
        Check(err)
        err = t.Execute(f, v)
        Check(err)
    }
}

func CreateFromUrl(tmpl_url string, config_url string, mac string) {
    tmpl, err := http_get(tmpl_url)
    Check(err)
    data, err := http_get(config_url)
    Check(err)
//  BatchCreate(string(tmpl), data)
    CreateCloudConfig(string(tmpl), data, mac)
}


func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func http_get(url string) ([]byte, error) {
    var body []byte
    resp, err := http.Get(url)
    if err != nil {
        return body, errors.New("HTTP ERROR!")
    }
    defer resp.Body.Close()
    body, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        return body, errors.New("ERROR: Cannot read data from http response.")
    }
    return body, nil
}
