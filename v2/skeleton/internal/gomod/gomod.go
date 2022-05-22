package gomod

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/mod/modfile"
)

func ModFile(dir string) (string, error) {
	var stdout bytes.Buffer
	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Dir = dir
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("can not get the parent module with %s: %w", dir, err)
	}

	return strings.TrimSpace(stdout.String()), nil
}

func ParentModule(dir string) (string, error) {
	gomodfile, err := ModFile(dir)
	if err != nil {
		return "", err
	}

	moddata, err := os.ReadFile(gomodfile)
	if err != nil {
		return "", fmt.Errorf("cat not read the go.mod of the parent module: %w", err)
	}

	gomod, err := modfile.Parse(gomodfile, moddata, nil)
	if err != nil {
		return "", fmt.Errorf("cat parse the go.mod of the parent module: %w", err)
	}

	return gomod.Module.Mod.Path, nil
}
