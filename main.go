package main

import (
	"github.com/armon/go-socks5"
	"github.com/caarlos0/env/v6"
	"log"
	"os"
)

type params struct {
	User            string   `env:"PROXY_USER" envDefault:""`
	Password        string   `env:"PROXY_PASSWORD" envDefault:""`
	Port            string   `env:"PROXY_PORT" envDefault:"1080"`
	AllowedDestFqdn string   `env:"ALLOWED_DEST_FQDN" envDefault:""`
	AllowedIPs      []string `env:"ALLOWED_IPS" envSeparator:"," envDefault:""`
}

func main() {

	// Working with app params
	cfg := params{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	//Initialize socks5 config
	socks5conf := &socks5.Config{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	if cfg.AllowedDestFqdn != "" {
		socks5conf.Rules = PermitDestAddrPattern(cfg.AllowedDestFqdn)
	}

	if cfg.User != "" && cfg.Password != "" {
		socks5conf.Credentials = socks5.StaticCredentials{
			cfg.User: cfg.Password,
		}
	}

	server, err := socks5.New(socks5conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Start listening proxy service on port %s\n", cfg.Port)
	if err := server.ListenAndServe("tcp", ":"+cfg.Port); err != nil {
		log.Fatal(err)
	}
}
