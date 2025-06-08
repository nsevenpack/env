package env

import (
	"github.com/nsevenpack/env/rootproject"
	"os"
	"path/filepath"
	"testing"
)

func TestGetWithEnvAndEnvTest(t *testing.T) {
	root := rootproject.Root()

	// data pour les fichier .env
	envPath := filepath.Join(root, ".env")
	envTestPath := filepath.Join(root, ".env.test")
	key := "MY_VAR"
	valueBase := "base-value"
	valueTest := "test-value"

	// create files
	err := os.WriteFile(envPath, []byte(key+"="+valueBase+"\n"), 0644)
	if err != nil {
		t.Fatalf("Impossible de créer .env : %v", err)
	}
	defer os.Remove(envPath)

	err = os.WriteFile(envTestPath, []byte(key+"="+valueTest+"\n"), 0644)
	if err != nil {
		t.Fatalf("Impossible de créer .env.test : %v", err)
	}
	defer os.Remove(envTestPath)

	// Set APP_ENV to "test"
	os.Setenv("APP_ENV", "test")
	defer os.Unsetenv("APP_ENV")

	got := Get(key)

	if got != valueTest {
		t.Errorf("env.Get(%q) = %q, attendu %q (valeur surchargée de .env.test)", key, got, valueTest)
	} else {
		t.Logf("env.Get(%q) = %q (OK, valeur surchargée .env.test)", key, got)
	}
}
