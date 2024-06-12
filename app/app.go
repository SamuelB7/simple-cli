package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Simple CLI App"
	app.Usage = "Searches for IPs and Server Names on the internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search for IPs address on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchIps,
		},
		{
			Name:  "serverName",
			Usage: "Searches for server name on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
