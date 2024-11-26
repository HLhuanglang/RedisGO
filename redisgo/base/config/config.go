package config

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Config ...
type Config struct {
	Bind  string   `cfg:"bind"`  //监听地址
	Port  int      `cfg:"port"`  //服务监听的端口
	Peers []string `cfg:"peers"` //集群节点
}

const (
	cfgYes = "yes"
)

var RedisGoCfg *Config

func init() {
	//默认配置
	RedisGoCfg = &Config{
		Bind: "127.0.0.1",
		Port: 6379,
	}
}

// SetupConfig 从文件中读取配置
func SetupConfig(cfgPath string) {
	//打开文件
	f, err := os.OpenFile(cfgPath, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//读取配置文件,生成map
	cfgMap := make(map[string]string)
	fscaner := bufio.NewScanner(f)
	for fscaner.Scan() {
		line := fscaner.Text()
		if len(line) > 0 {
			line = strings.TrimLeft(line, " ")
			if strings.HasPrefix(line, "#") {
				continue
			}
		}
		//exampleline
		//ip    127.0.0.1
		idx := strings.IndexAny(line, " ")
		if idx > 0 && idx < len(line)-1 {
			key := line[:idx]
			value := strings.Trim(line[idx+1:], " ")
			cfgMap[strings.ToLower(key)] = value
		}
	}
	if err := fscaner.Err(); err != nil {
		panic(err)
	}

	//反射能够获取对象的字段名、值、tag等信息
	log.Printf("cfgMap:%+v\n", cfgMap)
	t := reflect.TypeOf(RedisGoCfg)
	v := reflect.ValueOf(RedisGoCfg)
	n := t.Elem().NumField()
	for i := 0; i < n; i++ {
		field := t.Elem().Field(i)
		fieldVal := v.Elem().Field(i)
		key, ok := field.Tag.Lookup("cfg")
		if !ok {
			log.Printf("field:%s not found tag cfg\n", field.Name)
			continue
		}
		value, ok := cfgMap[strings.ToLower(key)]
		if !ok {
			log.Printf("key:%s not found in cfgMap\n", key)
			continue
		}
		switch field.Type.Kind() { //kind表示种类,也就是底层数据类型如int、string等
		case reflect.String:
			fieldVal.SetString(value)
		case reflect.Int:
			intValue, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				fieldVal.SetInt(intValue)
			}
		case reflect.Bool:
			boolValue := cfgYes == value
			fieldVal.SetBool(boolValue)
		case reflect.Slice:
			if field.Type.Elem().Kind() == reflect.String {
				slice := strings.Split(value, ",")
				fieldVal.Set(reflect.ValueOf(slice))
			}
		}
	}
	log.Printf("RedisGoCfg:%+v\n", RedisGoCfg)
}
