package main

import (
    "io/ioutil"
    "path/filepath"
    "yml_template"
)

func main() {
    tplFile, _ := filepath.Abs("./cloud-config.tpl")
    configFile, _ := filepath.Abs("./cluster.conf")
    configTpl, err := ioutil.ReadFile(tplFile)
    yml_template.Check(err)
    clusterConfig, err := ioutil.ReadFile(configFile)
    yml_template.Check(err)

    yml_template.BatchCreate(string(configTpl), clusterConfig)
    yml_template.CreateCloudConfig(string(configTpl), clusterConfig, "00:25:90:c0:f6:ee")
}

