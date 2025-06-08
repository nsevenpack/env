package env

import (
	"github.com/nsevenpack/env/rootproject"
	"os"
	"path/filepath"
	"testing"
)

func TestSetEnvFileNameAndGet(t *testing.T) {
	// Arrange : crée un fichier .env.test temporaire à la racine du projet
	root := rootproject.Root()
	envPath := filepath.Join(root, ".env.test")
	key := "MY_VAR"
	value := "test-value"
	err := os.WriteFile(envPath, []byte(key+"="+value+"\n"), 0644)
	if err != nil {
		t.Fatalf("Impossible de créer le fichier test .env : %v", err)
	}
	defer os.Remove(envPath) // Cleanup

	envPath2 := filepath.Join(root, ".env.dev")
	key2 := "MY_VAR_2"
	value2 := "test-value_2"
	err2 := os.WriteFile(envPath2, []byte(key2+"="+value2+"\n"), 0644)
	if err2 != nil {
		t.Fatalf("Impossible de créer le fichier test .env : %v", err2)
	}
	defer os.Remove(envPath2) // Cleanup

	SetEnvFiles(".env.test", ".env.dev") // Définit les fichiers à charger
	got := Get(key)
	got2 := Get(key2)

	// Assert premie fichier
	if got != value {
		t.Errorf("env.Get(%q) = %q, attendu %q", key, got, value)
	}
	if got == value {
		t.Logf("env.Get(%q) = %q, attendu %q", key, got, value)
	}

	// Assert deuxieme fichier
	if got2 != value2 {
		t.Errorf("env.Get(%q) = %q, attendu2 %q", key2, got2, value2)
	}
	if got2 == value2 {
		t.Logf("env.Get(%q) = %q, attendu2 %q", key2, got2, value2)
	}
}
