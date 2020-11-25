package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/Ilya001/Notes/internal/apiserver"
	"github.com/Ilya001/Notes/internal/note"
	"github.com/ivahaev/go-logger"
)

var (
	apiServerconfigPath   string
	tarantoolDbconfigPath string
)

func init() {
	flag.StringVar(&apiServerconfigPath, "configApi", "configs/apiserver.toml", "Path to Apiserver config file")
	flag.StringVar(&tarantoolDbconfigPath, "configDb", "configs/note.toml", "Path to Tarantool config file")
}

func main() {
	// Parse Flasgs
	flag.Parse()

	// Settings and decode configFile.toml
	configAPIServer := apiserver.NewConfig()
	_, err := toml.DecodeFile(apiServerconfigPath, &configAPIServer)
	if err != nil {
		logger.Crit(err)
		os.Exit(1)
	}

	configNoteDb := note.NewConfig()
	_, err = toml.DecodeFile(tarantoolDbconfigPath, &configNoteDb)
	if err != nil {
		logger.Crit(err)
		os.Exit(1)
	}

	// CreateNoteObj
	note := note.New(configNoteDb)

	// Create Sever and Run
	server := apiserver.New(configAPIServer, note)
	err = server.Start()
	if err := server.Start(); err != nil {
		logger.Crit(err)
		os.Exit(1)
	}
}
