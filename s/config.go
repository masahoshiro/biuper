package biuper

import (
	"os"
	"io/ioutil"
	"github.com/olebedev/config"
)

func readFile(path string) string {  
	fi, err := os.Open(path)
	if err != nil { panic(err)}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)  
}


func ReadConf(key string) string {

	cfg, _ := config.ParseYaml(readFile("config.yaml"))
	value, _ := cfg.String(key)
	return value
}
