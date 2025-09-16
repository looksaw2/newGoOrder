package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"
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


