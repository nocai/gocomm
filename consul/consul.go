package consul

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/nocai/gocomm"
	"os"
	"time"
)

func Register(l log.Logger, consulClient *api.Client, addr string, httpPort, grpcPort int, serviceName string) (httpRegister, grpcRegister *api.AgentServiceRegistration) {
	// KV
	kv(l, consulClient, addr, serviceName)

	// 服务注册
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: (5 * time.Second).String(),
		HTTP:                           fmt.Sprintf("http://%s:%d/echo?Req=%v", gocomm.LocalIP(), httpPort, serviceName),
		Timeout:                        "5s",
		Interval:                       "5s",
	}

	client := consul.NewClient(consulClient)
	if httpPort > 0 {
		// http register
		httpRegister = registration(httpPort, serviceName, check)
		if err := client.Register(httpRegister); err != nil {
			_ = level.Error(l).Log("msg", err)
			os.Exit(1)
		}
	}
	if grpcPort > 0 {
		// grpc register
		grpcRegister = registration(grpcPort, serviceName+"-grpc", check)
		if err := client.Register(grpcRegister); err != nil {
			_ = level.Error(l).Log("msg", err)
			os.Exit(1)
		}
	}
	//defer func() {
	//	// NOTE:服务注销
	//	if err := consulClient.Deregister(httpRegister); err != nil {
	//		_ = level.Error(l).Log("msg", err)
	//	}
	//}()

	return httpRegister, grpcRegister
}

func ConsulClient(l log.Logger, addr string) *api.Client {
	consulConfig := api.DefaultConfig()
	if len(addr) > 0 {
		consulConfig.Address = addr
	}
	consulApi, err := api.NewClient(consulConfig)
	if err != nil {
		_ = level.Error(l).Log("msg", err)
		os.Exit(1)
	}
	return consulApi
}

func registration(port int, serverName string, check *api.AgentServiceCheck) *api.AgentServiceRegistration {
	localIP := gocomm.LocalIP()
	return &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v:%v:%d", serverName, localIP, port),
		Name:    serverName,
		Address: localIP,
		Port:    port,
		Tags:    []string{serverName, "urlprefix-/" + serverName + " strip=/" + serverName},
		Check:   check,
	}
}

func kv(l log.Logger, consulApi *api.Client, consulAddr, servicePath string) {
	_ = level.Info(l).Log("msg", fmt.Sprintf("servicePath:%s", servicePath))

	kvp, _, err := consulApi.KV().Get(servicePath, nil)
	if err != nil || kvp == nil {
		_ = level.Error(l).Log("msg", fmt.Sprintf("err:%v, kvp:%v", err, kvp))
		os.Exit(-1)
	}
	if err = Unmarshal(kvp.Value); err != nil {
		_ = level.Error(l).Log("msg", err)
		os.Exit(-1)
	}

	go func() {
		if r := recover(); r != nil {
			_ = level.Error(l).Log("msg", r)
		}
		kvWatch(l, consulAddr, servicePath)
	}()
}

func kvWatch(l log.Logger, consulAddr, servicePath string) {
	var (
		plan *watch.Plan
		err  error
	)

	if plan, err = watch.Parse(map[string]interface{}{"type": "key", "key": servicePath,}); err != nil {
		_ = level.Error(l).Log("msg", err)
		os.Exit(1)
	}
	plan.HybridHandler = func(_ watch.BlockingParamVal, raw interface{}) {
		_ = level.Info(l).Log("msg", "update config file")
		if kvp, ok := raw.(*api.KVPair); ok {
			if err = Unmarshal(kvp.Value); err != nil {
				_ = level.Error(l).Log("msg", err)
			}
		}
	}

	if err := plan.Run(consulAddr); err != nil {
		_ = level.Error(l).Log("msg", err)
		os.Exit(1)
	}
}
