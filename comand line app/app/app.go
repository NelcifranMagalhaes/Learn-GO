package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar will return the command line app to run
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line app"
	app.Usage = "Search Ips and service name in net"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search Ips",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "livechart.me",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
