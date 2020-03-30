package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"time"
)

var conf *api.Config
var tlsConf api.TLSConfig
var Client *api.Client

// e9138d5b-c037-e88b-5cea-a381ae7be43e
// NewConsulClient returns a new client to Consul for the given address
func NewConsulClient(address, scheme, token, cert, key, caCert string) error {

	fmt.Println("=============> consul init")
	conf = api.DefaultConfig()

	conf.Scheme = scheme
	if len(address) > 0 {
		conf.Address = address

	}
	if len(token) > 0 {
		conf.Token = token
	}

	if scheme == "https" {
		tlsConf = api.TLSConfig{
			CertFile: cert,
			KeyFile:  key,
			CAFile:   caCert,
		}
		conf.TLSConfig = tlsConf
	}

	_client, err := NewClient()
	Client = _client
	if err != nil {
		return err
	}
	return nil
}

func NewClient() (*api.Client, error) {
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type watchResponse struct {
	ServiceEntry []*api.ServiceEntry
	Meta         *api.QueryMeta
	Err          error
}

type svcResponse struct {
	Services map[string][]string
	Meta     *api.QueryMeta
	Err      error
}

const waitTime time.Duration = 60

/**
循环侦听某一service变化
*/
func ServiceChange(serviceName string, isRun func(name string) bool, callback func(name string, entry []*api.ServiceEntry)) {
	//waitIndex初始版本
	var waitIndex uint64
	go func(lastIndex uint64) {
		for {
			//前提条件，service必须处于侦听状态(isRun返回true)
			if isRun(serviceName) {
				resp, err := serviceChange(serviceName, lastIndex)
				if err != nil {
					time.Sleep(1 * time.Second)
					continue
				}
				//有变化则回调处理，并更新版本
				if resp.Meta.LastIndex > lastIndex {
					go callback(serviceName, resp.ServiceEntry)
					lastIndex = resp.Meta.LastIndex
				}
			} else {
				return
			}
		}
	}(waitIndex)
}

/**
调用consul底层api查询某一service信息
*/
func serviceChange(serviceName string, lastIndex uint64) (watchResponse, error) {
	keys, meta, err := Client.Health().Service(serviceName, "", true, &api.QueryOptions{
		WaitTime:  waitTime * time.Second,
		WaitIndex: lastIndex,
	})
	if err != nil {
		return watchResponse{}, err
	}
	return watchResponse{keys, meta, err}, nil
}

/**
步骤1.2.1
循环查询consul微服务
*/
func Services(callback func(services map[string][]string)) {
	var waitIndex uint64
	//闭包函数，定义和使用一起，传参waitIndex
	go func(lastIndex uint64) {
		loopCount := 1
		fmt.Println("=============> ServicesLoopStart", loopCount)
		for {
			resp, err := services(lastIndex)
			loopCount++
			fmt.Println("=============> ServicesLoopStartLoopCount", loopCount, " lastIndex", lastIndex)
			if err != nil {
				//如果有错，休息一秒，继续循环
				time.Sleep(1 * time.Second)
				continue
			}
			if resp.Meta.LastIndex > lastIndex {
				go callback(resp.Services)
				lastIndex = resp.Meta.LastIndex
			}
		}
		fmt.Println("=============> ServicesLoopStart", loopCount)
	}(waitIndex)
}

/**
步骤1.2.2
调用consul底层api查询所有微服务
lastIndex：long pull索引
*/
func services(lastIndex uint64) (svcResponse, error) {
	//查询所有services服务( long pull consul 60秒超时）
	keys, meta, err := Client.Catalog().Services(&api.QueryOptions{
		WaitTime:  waitTime * time.Second, //超时时间60秒
		WaitIndex: lastIndex,
	})
	if err != nil {
		return svcResponse{}, err
	}
	//将查询获得的服务元素封装成svcResponse结构体返回
	return svcResponse{keys, meta, err}, nil
}

/**
1.3.1 循环侦听配置变化
*/
func KVWatch(key string, callback func(map[string]string)) {
	var waitIndex uint64
	//开启多线程循环侦听
	go func(lastIndex uint64) {
		for {
			val, meta, err := kvwatch(key, lastIndex)
			if err != nil {
				time.Sleep(1 * time.Second)
				continue
			}
			if meta.LastIndex > lastIndex {
				//fmt.Println(val)
				go callback(val)
				lastIndex = meta.LastIndex
			}
		}
	}(waitIndex)
}

/**
1.3.2 调用Consul底层Api，longpull方式侦听配置变化
*/
func kvwatch(key string, lastIndex uint64) (map[string]string, *api.QueryMeta, error) {
	val, meta, err := Client.KV().List(key, &api.QueryOptions{
		WaitTime:  waitTime * time.Second,
		WaitIndex: lastIndex,
	})
	if err != nil {
		return nil, nil, err
	}
	vals := make(map[string]string)
	for _, x := range val {
		if "/"+x.Key != key {
			k := ("/" + x.Key)[len(key):]
			vals[k] = string(x.Value)
		}
	}
	return vals, meta, err
}
