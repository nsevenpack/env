package env

import (
	"github.com/joho/godotenv"
	"github.com/nsevenpack/env/rootproject"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once

func loadEnvOnce() {
	root := rootproject.Root()
	baseEnv := filepath.Join(root, ".env")

	// .env de base
	_ = godotenv.Load(baseEnv)

	// surcharger avec .env spécifique en fonction de APP_ENV de l'os
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "" && appEnv != "dev" {
		specificEnv := filepath.Join(root, ".env."+appEnv)
		_ = godotenv.Overload(specificEnv)
	}
}

func Get(key string) string {
	once.Do(loadEnvOnce)
	return os.Getenv(key)
}
