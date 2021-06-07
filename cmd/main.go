package main

import (
	"context"
	"github.com/easy-project-templete/internal/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	pkgLog "github.com/easy-project-templete/pkg/log"
	"github.com/easy-project-templete/config"

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
	_, err = service.New()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ctx, _ := context.WithCancel(context.Background())

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