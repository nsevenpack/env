package env

import (
	"github.com/joho/godotenv"
	"github.com/nsevenpack/env/rootproject"
	"os"
	"path/filepath"
	"sync"
)

var (
	once         sync.Once
	envFileNames = []string{".env"} // liste de fichiers à loader // par defaut .env
)

func SetEnvFiles(files ...string) {
	envFileNames = files
}

func loadEnvOnce() {
	for _, f := range envFileNames {
		envPath := filepath.Join(rootproject.Root(), f)
		_ = godotenv.Overload(envPath) // Overload fusionne, Load charge une seule fois
	}
}

func Get(key string) string {
	once.Do(loadEnvOnce)
	return os.Getenv(key)
}
