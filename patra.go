package patra

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/teacat/patra/balancer"
	"github.com/teacat/patra/model"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var flags = []cli.Flag{
	cli.IntFlag{
		EnvVar: "PATRA_PORT",
		Name:   "port",
		Value:  40051,
	},
	cli.IntFlag{
		EnvVar: "PATRA_INTERVAL",
		Name:   "interval",
		Value:  30,
	},
	cli.StringFlag{
		EnvVar: "PATRA_CONSUL_ADDR",
		Name:   "consulAddress",
		Value:  "",
	},
	cli.StringFlag{
		EnvVar: "PATRA_BALANCE_TYPE",
		Name:   "balanceType",
		Value:  "least-usage",
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "patra"
	app.Version = "1.0.0"
	app.Action = startup
	app.Flags = flags
	app.Run(os.Args)
}

func startup(c *cli.Context) {
	var b balancer.Balancer
	var err error

	//
	if c.String("consulAddress") == "" {
		b, err = balancer.NewDefault()
	}
	//
	if c.String("consulAddress") != "" {
		conf := api.DefaultConfig()
		conf.Address = c.String("consulAddress")
		b, err = balancer.NewConsul(conf)
	}
	if err != nil {
		log.Fatalln(err)
	}

	var typ balancer.BalanceType
	switch c.String("balanceType") {
	case "least-usage":
		typ = balancer.TypeLeastUsage
		break
	case "random":
		typ = balancer.TypeRandom
		break
	case "least-connection":
		typ = balancer.TypeLeastConnection
		break
	default:
		typ = balancer.TypeRoundRobin
		break
	}

	conf := &Config{
		Port:     c.Int("port"),
		Interval: c.Int("interval"),
		Type:     typ,
	}
	if err := New(conf, b); err != nil {
		log.Fatalln(err)
	}
}

// Config
type Config struct {
	// Port
	Port int
	// Interval
	Interval int
	//
	Type balancer.BalanceType
}

// New
func New(c *Config, b balancer.Balancer) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		return error
	}

	s := grpc.NewServer()
	model.RegisterRegistryServer(s, &server{
		balancer: b,
	})

	reflection.Register(s)

	return s.Serve(lis)
}
