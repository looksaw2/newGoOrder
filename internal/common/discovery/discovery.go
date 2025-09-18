package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/looksaw/go-orderv2/common/discovery/consul"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)


type Registry interface {
	//注册服务
	Register(ctx context.Context,instanceID string,serviceName string , hostPort string) error
	//注销服务
	Deregister(ctx context.Context,instanceID string , serviceName string) error
	//服务发现
	Discover(ctx context.Context,serviceName string)([]string,error)
	//健康检查
	HealthCheck(instanceID string , serviceName string) error 
}


//生成InstanceID
func GenerateInstanceID(serviceName string) string {
	x := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	return fmt.Sprintf("%s-%d",serviceName,x)
}

//得到服务地址
func GetServiceAddr(ctx context.Context , serviceName string)(string ,error){
	logrus.Infof("consul addr is %s",viper.GetString("consul.addr"))
	registry ,err := consul.New(viper.GetString("consul.addr"))
	if err != nil {
		return "" , err
	}
	addrs , err := registry.Discover(ctx,serviceName)
	logrus.Infof("%s dicovery addr is %v",serviceName , addrs)
	if err != nil {
		return "",err
	}
	if len(addrs) == 0 {
		return "",fmt.Errorf("got empty %s addr from consul",serviceName)
	}
	i := rand.Intn(len(addrs))
	logrus.Infof("Discovered %d instance of %s ,addrs=%v",len(addrs),serviceName,addrs)
	return addrs[i] , nil
}