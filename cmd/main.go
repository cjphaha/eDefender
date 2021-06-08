package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cjphaha/eDefender/config"
	"github.com/cjphaha/eDefender/internal/server"
	"github.com/cjphaha/eDefender/internal/service"
	pkgLog "github.com/cjphaha/eDefender/pkg/log"
	_ "github.com/cjphaha/eDefender/plugin/goPlugins"

	log "github.com/sirupsen/logrus"
)

func main() {
	// new log
	err := pkgLog.InitLog(config.AppConfig.Log)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// new service
	srv, err := service.New(config.AppConfig.Service)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// new server
	ctx, cancel := context.WithCancel(context.Background())
	s, err := server.New(ctx, cancel, srv, config.AppConfig.Server)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// start program
	go s.Start()
	// welecome
	fmt.Println(config.AppConfig.Base.Welecome)
	initSignal(ctx)
}

func initSignal(ctx context.Context) {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		log.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			log.Printf("app need reload")
		default:
			time.AfterFunc(time.Duration(time.Second*5), func() {
				log.Warn("app exit now by force...")
				os.Exit(1)
			})
			ctx.Done()
			// The program exits normally or timeout forcibly exits.
			log.Warnf("app exit now...")
			return
		}
	}
}