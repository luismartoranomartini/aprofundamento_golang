package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {

	app := cli.NewApp()

	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Buscar IP's e nomes de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "nexxum.co",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca Ip's de endereços na internet",
			Flags:  flags,
			Action: buscaIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores",
			Flags:  flags,
			Action: buscaServidores,
		},
	}

	return app
}

func buscaIPs(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host) // name server
	if err != nil {
		log.Fatal(err)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
