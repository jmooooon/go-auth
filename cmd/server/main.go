package main

import (
	"flag"
	"log"
	"net"

	"github.com/BurntSushi/toml"
	"github.com/jmooooon/go-auth/pkg/api"
	"github.com/jmooooon/go-auth/pkg/auth"
	"google.golang.org/grpc"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/auth.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := auth.NewConfig()
	
	_, err := toml.DecodeFile(configPath, config)
	
	if err != nil {
		log.Fatal(err)
	}

	err= auth.NewDB(config.DB)
	
	if err != nil {
		log.Fatal(err)
	}

	defer auth.DB.Close()

	s := grpc.NewServer()
	srv := &auth.GRPCServer{}
	api.RegisterAuthServer(s, srv)
	l,err := net.Listen("tcp", config.Port)

	if err != nil {
		log.Fatal(err)
	}
	
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}