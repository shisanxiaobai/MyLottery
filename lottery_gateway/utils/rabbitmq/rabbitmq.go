package rabbitmq

import (
	"fmt"
	_ "lottery_gateway/conf"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var (
	Conn *amqp.Connection
	Ch   *amqp.Channel
	err  error
)

func init() {
	//读取配置
	user := viper.GetString("rabbit.user")
	password := viper.GetString("rabbit.password")
	host := viper.GetString("rabbit.host")
	port := viper.GetString("rabbit.port")
	vhost := viper.GetString("rabbit.vhost")
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s%s",
		user, password, host, port, vhost,
	)
	Conn, err = amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}
	Ch, _ = Conn.Channel()

	Ch.QueueDeclare("order_queue", true, false, false, false, nil)
	Ch.ExchangeDeclare("order_exchange", amqp.ExchangeDirect, true, false, false, false, nil)
	Ch.QueueBind("order_queue", "", "order_exchange", false, nil)
}

func Publish(msg string) error {
	err := Ch.Publish("order_exchange", "", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(msg),
		DeliveryMode: amqp.Persistent,
	})
	return err
}

func Close() {
	Ch.Close()
	Conn.Close()
}
