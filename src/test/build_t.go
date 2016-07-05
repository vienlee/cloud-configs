package main

import (
//    "io/ioutil"
//    "path/filepath"
    "yml_template"
)

func main() {
    tpl_url := "https://raw.githubusercontent.com/k8sp/cloud-configs/liangjiameng/src/yml_template/cloud-config.tpl"
    config_url := "https://raw.githubusercontent.com/k8sp/cloud-configs/liangjiameng/src/yml_template/cluster.conf"
    yml_template.CreateFromUrl(tpl_url, config_url, "00:25:90:c0:f6:ee")
/*    tplFile, _ := filepath.Abs("./cloud-config.tpl")
    configFile, _ := filepath.Abs("./cluster.conf")
    configTpl, err := ioutil.ReadFile(tplFile)
    yml_template.Check(err)
    clusterConfig, err := ioutil.ReadFile(configFile)
    yml_template.Check(err)

    yml_template.BatchCreate(string(configTpl), clusterConfig)
    yml_template.CreateCloudConfig(string(configTpl), clusterConfig, "00:25:90:c0:f6:ee")
*/
}

