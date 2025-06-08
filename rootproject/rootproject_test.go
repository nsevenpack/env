package rootproject

import (
	"path/filepath"
	"strings"
	"testing"
)

func Test_Root(t *testing.T) {
	r := Root()

	if !filepath.IsAbs(r) {
		t.Errorf("Root() doit retourner un chemin absolu, obtenu: %s", r)
	}

	if !strings.Contains(r, "app") {
		t.Errorf("Root() doit contenir le chemin du projet, obtenu: %s", r)
	}

	if strings.Contains(r, "app") {
		t.Logf("Root() retourne la racine du projet: %s", r)
	}
}

func Test_PathWithoutArg(t *testing.T) {
	r := Root()
	p := Path()

	if p != r {
		t.Errorf("Path() sans argument doit retourner la racine, attendu %s, obtenu %s", r, p)
	}

	if p == r {
		t.Logf("Path() sans argument retourne la racine du projet: %s", p)
	}
}

func Test_PathWithArgs(t *testing.T) {
	r := Root()
	p := Path("internal", "config", "dev.env")

	expected := filepath.Join(r, "internal", "config", "dev.env")

	if p != expected {
		t.Errorf("Path(\"internal\", \"config\", \"dev.env\") attendu %s, obtenu %s", expected, p)
	}

	if p == expected {
		t.Logf("Path(\"internal\", \"config\", \"dev.env\") attendu %s, obtenu %s", expected, p)
	}
}
