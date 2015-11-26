package app

import (
	"net/http"
	"sync"

	log "github.com/Sirupsen/logrus"
)

// TODO(jackson): Move configuration data into a file.
const (
	publicListenAddress   = ":8080"
	internalListenAddress = ":8081"
)

// App encapsulates the entire itsawordgame.com server.
type App struct {
	publicServer   http.Server
	publicMux      http.ServeMux
	internalServer http.Server
	internalMux    http.ServeMux
	stopped        sync.WaitGroup
}

// Start starts up the server.
func (a *App) Start() {
	// Initialize the http servers.
	a.publicServer = http.Server{
		Addr:    publicListenAddress,
		Handler: http.HandlerFunc(a.servePublicHTTP),
	}
	a.internalServer = http.Server{
		Addr:    internalListenAddress,
		Handler: http.HandlerFunc(a.serveInternalHTTP),
	}

	go a.listenAndServe(&a.publicServer)
	go a.listenAndServe(&a.internalServer)

	a.stopped.Add(1)
}

func (a *App) Stop() {
	a.stopped.Done()
}

func (a *App) Wait() {
	a.stopped.Wait()
}

func (a *App) listenAndServe(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		log.Errorf("Error serving http traffic on `%s`: %s", server.Addr, err)
	}
}

func (a *App) servePublicHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Infof("[public] %s — %s %s", req.RemoteAddr, req.Method, req.URL.Path)
	a.publicMux.ServeHTTP(rw, req)
}

func (a *App) serveInternalHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Infof("[internal] %s — %s %s", req.RemoteAddr, req.Method, req.URL.Path)
	a.internalMux.ServeHTTP(rw, req)
}
