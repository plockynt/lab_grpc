package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"local/internal/application"
	"local/internal/infrastructure/persistance/fs"
	"local/internal/shared"
	"local/internal/userinterface/restapigingonic"
)

var (
	svcPort = shared.GetenvOr("SVC_PORT", "8080")
)

// ________________________________ Adapters management ________________________________

func newAdapters(ctx context.Context) (*adapters, error) {
	persistence, err := fs.New(ctx, "/tmp/test.json")
	if err != nil {
		return nil, err
	}

	return &adapters{
		persistence: persistence,
	}, nil
}

type adapters struct {
	persistence application.PersistencePort
}

func (obj *adapters) Close(ctx context.Context) {
	if obj.persistence != nil {
		obj.persistence.Close(ctx)
	}
}

// ________________________________ main ________________________________

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create adapters
	adapters, err := newAdapters(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer adapters.Close(ctx)

	// Create application services
	var projectsSvc = application.NewProjectsSvc(adapters.persistence)

	// Create router
	var router = restapigingonic.NewRouter(projectsSvc)
	// Serve
	startServer(ctx, router)
}

func startServer(ctx context.Context, router http.Handler) error {
	var srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", svcPort),
		Handler: router,
	}

	go shared.ShutdownOnSignal(ctx, srv, 5*time.Second)

	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
