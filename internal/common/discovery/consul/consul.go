package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

type Registry struct {
	client *api.Client
}

var (
	consulClient *Registry
	once sync.Once
	initErr error
)

func New(consulAddr string)(*Registry,error){
	once.Do(func() {
		config := api.DefaultConfig()
		config.Address = consulAddr
		client , err := api.NewClient(config)
		if err != nil {
			initErr = err
			return 
		}
		consulClient = &Registry{
			client: client,
		}
	})
	if initErr != nil {
		return nil , initErr
	}
	return consulClient , nil
}

func(r *Registry)Register(ctx context.Context,instanceID string,serviceName string , hostPort string) error{
	parts := strings.Split(hostPort,":")
	if len(parts) != 2 {
		return errors.New("invalid host:port format")
	}	
	host := parts[0]
	port , _  := strconv.Atoi(parts[1])
	return r.client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID: instanceID,
		Address: host,
		Port: port,
		Name: serviceName,
		Check: &api.AgentServiceCheck{
			CheckID: instanceID,
			TLSSkipVerify: false,
			TTL: "5s",
			Timeout: "5s",
			DeregisterCriticalServiceAfter: "10s",
		},
	})

}

func(r *Registry)Deregister(ctx context.Context,instanceID string , serviceName string) error{
	logrus.WithFields(logrus.Fields{
		"instanceID" : instanceID,
		"serviceName" : serviceName,
	}).Info("deregister service")
	return r.client.Agent().CheckDeregister(instanceID)
}
func(r *Registry)Discover(ctx context.Context,serviceName string)([]string,error){
	logrus.Info("before discovery")
	entries , _ , err := r.client.Health().Service(serviceName,"",true,nil)
	logrus.Infof("after discovery is %v",err)
	if err != nil {
		return nil ,err
	}
	var ips []string
	if len(entries) == 0 {
		panic("service is 0 instance")
	}
	for _ ,e := range entries {
		ips = append(ips, fmt.Sprintf("%s:%d",e.Service.Address ,e.Service.Port))
	}
	logrus.Infof("%s ips is %v",serviceName,ips)
	return ips , nil
}
func(r *Registry)HealthCheck(instanceID string , serviceName string) error {
	return r.client.Agent().UpdateTTL(instanceID,"onlien",api.HealthPassing)
}