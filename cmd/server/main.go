package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/thatmounaim/dzforex.microservice/internal/exchange"
	"github.com/thatmounaim/dzforex.microservice/internal/storage"
	"github.com/thatmounaim/dzforex.microservice/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	sc := &exchange.Scrapper{
		Endpoint:   "http://www.forexalgerie.com/connect/updateExchange.php",
		Passkey:    "afous",
		Passphrase: "moh!12!",
	}

	ds := storage.NewMemoryStore()
	l := hclog.Default()

	svc := exchange.NewExchangeService(sc, ds, l)
	svc.UpdateData()
	srv := grpc.NewServer()
	proto.RegisterDzForexServer(srv, svc)
	reflection.Register(srv)
	listener, err := net.Listen("tcp", ":4444") // Hardcoded lisenAddr i know

	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}

	err = srv.Serve(listener)

	if err != nil {
		l.Error(err.Error())
		os.Exit(1)
	}

}
