package broker

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)
func Connect(user string, password string , host string , port string)(*amqp.Channel ,func() error ,  error){
	address := fmt.Sprintf("amqp://%s:%s@%s:%s",user,password,host,port)
	conn ,err := amqp.Dial(address)
	if err != nil {
		return nil , func() error {return nil} ,err
	}
	ch , err := conn.Channel()
	if err != nil {
		return nil , func() error {return nil} ,err
	}
	err = ch.ExchangeDeclare(EventOrderCreate,"direct",true,false,false,false,nil)
	if err != nil {
		return nil  , func() error {return nil} ,err
	}
	err = ch.ExchangeDeclare(EventOrderPaid,"fanout",true,false,false,false,nil)
	if err != nil {
		return nil , func() error {return nil} ,err
	}
	return ch , ch.Close , nil 
}