package assembler

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *assembler) ReleaseResourcesAPI(e *echo.Echo) {
	a.TerminateSignal()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	releasers := []func(context.Context){}

	select {
	case <-ctx.Done():
		log.Println("release resources timeout")
	case <-a.release(ctx, releasers...):
		log.Println("all resources released")
	}

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	log.Println("Exited gracefully...")
}

func (a *assembler) release(
	ctx context.Context,
	releaseResources ...func(context.Context),
) <-chan struct{} {
	wg := new(sync.WaitGroup)
	released := make(chan struct{}, 1)

	wg.Add(len(releaseResources))
	for i := range releaseResources {
		releaseResource := releaseResources[i]
		go func(ctx context.Context) {
			defer wg.Done()
			releaseResource(ctx)
		}(ctx)
	}

	go func() {
		wg.Wait()
		released <- struct{}{}
	}()

	return released
}
