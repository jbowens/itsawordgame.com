package app

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

// TODO(jackson): Move configuration data into a file.
const (
	publicListenAddress   = ":8085"
	internalListenAddress = ":8086"
)

// App encapsulates the entire itsawordgame.com server.
type App struct {
	gamekeeper
	templates      *template.Template
	publicServer   http.Server
	publicMux      *http.ServeMux
	internalServer http.Server
	internalMux    *http.ServeMux
	upgrader       websocket.Upgrader
	stopped        sync.WaitGroup
}

// Start starts up the server.
func (a *App) Start() {
	a.gamekeeper.init()

	// Load all the html templates
	a.templates = template.Must(template.ParseFiles("static/index.html"))

	// Initialize the websocket upgrader
	a.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Initialize the http servers.
	a.publicServer = http.Server{
		Addr:    publicListenAddress,
		Handler: http.HandlerFunc(a.servePublicHTTP),
	}
	a.internalServer = http.Server{
		Addr:    internalListenAddress,
		Handler: http.HandlerFunc(a.serveInternalHTTP),
	}
	a.publicMux = http.NewServeMux()
	a.internalMux = http.NewServeMux()

	// Setup the routes.
	a.publicMux.HandleFunc("/connect", a.websocketUpgradeRoute)
	a.publicMux.HandleFunc("/time", a.time)
	a.publicMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	a.publicMux.HandleFunc("/", a.index)

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

func (a *App) websocketUpgradeRoute(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(rw, "Method not allowed", 405)
		return
	}

	ws, err := a.upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Infof("Initializing a new client from host %s", req.RemoteAddr)
	a.gamekeeper.ConnectingClients <- newClient(req.RemoteAddr, ws, a.incomingMessages, a.gamekeeper.DisconnectingClients)
}

func (a *App) time(rw http.ResponseWriter, req *http.Request) {
	response := struct {
		ServerTime time.Time `json:"server_time"`
	}{
		ServerTime: time.Now(),
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Errorf("Error marshalling json: %s", err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)
}

func (a *App) index(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFoundHandler().ServeHTTP(rw, req)
		return
	}

	a.templates.Execute(rw, nil)
}
