package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Generate returns a new cli.App with our desired configuration
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Who Am I Application"
	app.Usage = "Search IPs and Domains to find out who they are"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "ips",
			Aliases: []string{"i"},
			Usage:   "Search for IPs information",
			Action:  searchIPs,
		},
		{
			Name:    "domains",
			Aliases: []string{"d"},
			Usage:   "Search for Domains information",
			Action:  searchDomains,
		},
	}
	return app
}

func searchIPs(c *cli.Context) {
	if c.NArg() != 1 {
		fmt.Println("Usage: whoami ips <host>. You must provide a host to search IPs, e.g.: google.com")
		return
	}

	host := c.Args().Get(0)
	ips, error := net.LookupIP(host)
	if error != nil {
		fmt.Println(error)
	} else {
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}
}

func searchDomains(c *cli.Context) {
	if c.NArg() != 1 {
		fmt.Println("Usage: whoami domains <ip_address>. You must provide a host to search IPs, e.g.: 1.1.1.1")
		return
	}

	host := c.Args().Get(0)
	domains, error := net.LookupAddr(host)
	if error != nil {
		fmt.Println(error)
	} else {
		for _, domain := range domains {
			fmt.Println(domain)
		}
	}
}
