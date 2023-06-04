package server

import (
	"context"
	"fmt"
	"github.com/kenlabs/permastar/pkg/api/core"
	"github.com/kenlabs/permastar/pkg/api/v1/server/httpserver"
	"github.com/kenlabs/permastar/pkg/option"
	"github.com/kenlabs/permastar/pkg/util/log"
	"github.com/kenlabs/permastar/pkg/util/multiaddress"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var logger = log.NewSubsystemLogger()

type Server struct {
	Core *core.Core
	Opt  *option.DaemonOptions

	HttpServer     *http.Server
	HttpListenAddr string
}

func NewAPIServer(opt *option.DaemonOptions, core *core.Core) (*Server, error) {
	httpListenAddress, err := multiaddress.MultiaddressToNetAddress(opt.ServerAddress.HttpAPIListenAddress)
	if err != nil {
		return nil, err
	}

	s := &Server{
		Core: core,
		Opt:  opt,
		HttpServer: &http.Server{
			Addr:    httpListenAddress,
			Handler: httpserver.NewHttpRouter(core, opt),
		},
		HttpListenAddr: httpListenAddress,
	}

	return s, nil
}

func (s *Server) StartHttpServer() error {
	logger.Infof("http server listening at: %v", s.HttpListenAddr)
	return s.HttpServer.ListenAndServe()
}

func (s *Server) StopHttpServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("stop http server...")
	return s.HttpServer.Shutdown(ctx)
}

func (s *Server) MustStartAllServers() {
	go func() {
		err := s.StartHttpServer()
		if err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("http api server cannot start: %v", err))
		}
	}()
}

func (s *Server) StopAllServers() error {
	g := errgroup.Group{}
	g.Go(func() error {
		return s.StopHttpServer()
	})
	err := g.Wait()
	if err != nil {
		return err
	}

	fmt.Println("Bye, permastar!")
	return nil
}
