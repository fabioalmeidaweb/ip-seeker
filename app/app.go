package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns a new instance of the cli.App struct.
//
// It creates a new cli.App instance, sets the Name and Usage fields,
// and initializes the flags and commands for IP and server information
// seeking from the Internet.
//
// Returns:
// - A pointer to a cli.App struct.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IP Seeker"
	app.Usage = "Seek IP and server information from the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Seek IP information from the Internet",
			Flags:  flags,
			Action: seekIps,
		},
		{
			Name:   "server",
			Usage:  "Seek server information from the Internet",
			Flags:  flags,
			Action: seekServers,
		},
	}

	return app
}

// seekIps is a Go function that looks up the IP addresses associated with a given host.
//
// It takes a single parameter:
// - c: a pointer to a cli.Context object
func seekIps(c *cli.Context) error {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip.String())
	}

	return nil
}

// seekServers is a Go function that performs a DNS lookup for the given host and prints the name servers associated with it.
//
// It takes a *cli.Context parameter named "c" which is used to retrieve the value of the "host" flag.
func seekServers(c *cli.Context) error {
	host := c.String("host")

	servers, err := net.LookupNS(host) // Name Server

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

	return nil
}
