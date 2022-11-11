package dao

import (
	"fmt"

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
	Ch, err = Conn.Channel()
	if err != nil {
		panic(err)
	}
	Ch.Qos(1, 0, false)
}

func Close() {
	Ch.Close()
	Conn.Close()
}
