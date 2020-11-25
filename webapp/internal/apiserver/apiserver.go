package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/Ilya001/Notes/internal/note"
	"github.com/gorilla/mux"
	"github.com/ivahaev/go-logger"
)

// APIserver ...
type APIserver struct {
	config *Config
	router *mux.Router
	note   *note.Note
}

// New ...
func New(apiServerConfig *Config, note *note.Note) *APIserver {
	return &APIserver{
		config: apiServerConfig,
		router: mux.NewRouter(),
		note:   note,
	}
}

// Start ...
func (server *APIserver) Start() error {
	server.configRouter()

	logger.Info("Server start in ", server.config.BindAddr)
	return http.ListenAndServe(server.config.BindAddr, server.router)
}

// confRouter ...
func (server *APIserver) configRouter() {
	server.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(server.config.StaticDir))))
	server.router.HandleFunc("/", server.IndexPage())
	server.router.HandleFunc("/api/notes/", server.GetAllNotes()).Methods("GET")
	server.router.HandleFunc("/api/notes/", server.CreateNote()).Methods("POST")
	server.router.HandleFunc("/api/notes/{id}", server.GetOneNote()).Methods("GET")
	server.router.HandleFunc("/api/notes/{id}", server.UpdateNote()).Methods("POST")
	server.router.HandleFunc("/api/notes/{id}", server.DeleteOneNote()).Methods("DELETE")
}

// WEBLogger ...
func (server *APIserver) WEBLogger(r *http.Request) {
	logger.Notice(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
}

// IndexPage ...
func (server *APIserver) IndexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadFile("public/templates/index.html")
		server.WEBLogger(r)
		if err != nil {
			logger.Crit(err)
		}
		io.WriteString(w, string(body))
	}
}

// GetAllNotes ...
func (server *APIserver) GetAllNotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		server.WEBLogger(r)
		allNotes, err := server.note.GetAllNotes(10)
		if err != nil {
			logger.Crit(err)
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(allNotes)
	}
}

// GetOneNote ...
func (server *APIserver) GetOneNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		server.WEBLogger(r)
		requestData := mux.Vars(r)["id"]
		id, err := strconv.Atoi(requestData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err)
		}
		note, err := server.note.GetOneNote(id)
		if err != nil {
			logger.Crit(err)
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(note)
	}
}

// CreateNote ...
func (server *APIserver) CreateNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		server.WEBLogger(r)
		var not note.NotStruct
		err := json.NewDecoder(r.Body).Decode(&not)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err)
		}
		err = server.note.CreateNote(&not)
		if err != nil {
			logger.Crit(err)
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(not)
	}
}

// UpdateNote ...
func (server *APIserver) UpdateNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		server.WEBLogger(r)
		var not note.NotStruct
		requestData := mux.Vars(r)["id"]
		id, err := strconv.Atoi(requestData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err)
		}
		not.ID = uint64(id)
		err = json.NewDecoder(r.Body).Decode(&not)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err)
		}
		err = server.note.UpdateNote(&not)
		if err != nil {
			logger.Crit(err)
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(not)
	}
}

// DeleteOneNote ...
func (server *APIserver) DeleteOneNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		server.WEBLogger(r)
		requestData := mux.Vars(r)["id"]
		id, err := strconv.Atoi(requestData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Error(err)
		}
		err = server.note.DeleteOneNote(id)
		if err != nil {
			logger.Crit(err)
		}
		allNotes, err := server.note.GetAllNotes(10)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logger.Crit(err)
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(allNotes)
	}
}
