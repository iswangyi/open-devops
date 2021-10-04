package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	MysqlS []*MySQLConf `yaml:"mysql"`
}
type MySQLConf struct {
	Name  string `yaml:"name"`
	Addr  string `yaml:"addr"`
	Max   int    `yaml:"max"`
	Idle  int    `yaml:"idle"`
	Debug bool   `yaml:"debug"`
}

// LoadConf IO读取配置文件
func LoadConf(fileName string) ([]byte, error) {
	cnf, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read config err:", err)
		return nil, err
	}
	return cnf, nil
}

// Load 解析配置文件
func Load(fileName string) (*Conf, error) {
	cnf, err := LoadConf(fileName)
	Conf := &Conf{}
	err = yaml.Unmarshal(cnf, Conf)
	if err != nil {
		fmt.Println("配置解析失败")
		return nil, err
	}
	return Conf, nil
}
