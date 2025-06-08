# env

[![Go Reference](https://pkg.go.dev/badge/github.com/nsevenpack/env.svg)](https://pkg.go.dev/github.com/nsevenpack/env)
[![Version](https://img.shields.io/github/v/tag/nsevenpack/env?label=version&sort=semver)](https://github.com/nsevenpack/env/releases)
[![CI](https://github.com/nsevenpack/env/actions/workflows/release.yml/badge.svg)](https://github.com/nsevenpack/env/actions/workflows/release.yml)
[![License](https://img.shields.io/github/license/nsevenpack/env)](https://github.com/nsevenpack/env/blob/main/LICENSE)

Package de configuration d'environnement simple pour Go, pour recupérer les variables d'environnement et les valeurs.

## Installation

```bash
# installe la derniere version 1.x.x
go get github.com/nsevenpack/env@latest

# liste les versions disponibles pour 1.x.x
go list -m -versions github.com/nsevenpack/env

# installe une version précise
go get github.com/nsevenpack/env@v1.1.0
```

## Fonctionnalités

- Initialisation des variables d'environnement à partir d'un fichier `.env`.
ou autres fichier `.env` il est possible de les surcharger.

## Utilisation
```go
package main
import (
    "github.com/nsevenpack/env"
)

func main() {
	env.SetEnvFileName(".env.dev")
	val := env.Get("MON_SECRET")
	// ...
}

// avec plusieurs fichiers env
func main() {
	env.SetEnvFileName(".env.dev", ".env.test", etc ...)
	val := env.Get("MON_SECRET")
	// ...
}

```
