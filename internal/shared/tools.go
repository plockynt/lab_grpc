package shared

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var EnvType = GetenvOr("ENV", "dev")

func GetenvOr(name string, defaultV string) string {
	var ret = os.Getenv(name)
	if ret == "" {
		ret = defaultV
	}
	return ret
}

func ShutdownOnSignal(ctx context.Context, srv *http.Server, timeToShutdown time.Duration) {
	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(ctx, timeToShutdown)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	} else {
		fmt.Println("Server exiting")
	}
}
