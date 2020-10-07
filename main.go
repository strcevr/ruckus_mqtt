package main

import (
	"./proto"
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var args struct {
	Username string `arg:"-u,--username" help:"provide a username" default:"testuser"`
	Password string `arg:"-p" help:"provide a password", default:""`
	Host     string `arg:"-H, required" help:"mqtt host to connect to"`
	Port     int    `arg:"required,-P" help:"network port to connect"`
	Topic    string `arg:"-t" help:"mqtt topic to subscribe to" default:"sci-topic"`
	ClientID string `arg:"--id" help:"id to use for this client" default:"test-client"`
	Verbose  bool   `arg:"-v,--verbose" help:"verbosity level"`
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	m := proto.SciMessage{}
	err := m.Unmarshal(msg.Payload())
	if err != nil {
		log.Println(err)
	} else {
		b, err := json.Marshal(m)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Payload: ", string(b))
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}()
	arg.MustParse(&args)
	if args.Port == 0 {
		args.Port = 1884
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", args.Host, args.Port))
	opts.SetClientID(args.ClientID)
	opts.SetUsername(args.Username)
	opts.SetPassword(args.Password)
	opts.SetDefaultPublishHandler(f)
	topic := args.Topic

	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Println("Connected to server")
	}
	<-c
}
