package tests

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Workspace struct {
	entry     fs.DirEntry
	directory string
}

func discoverWorkspaces() (workspaces []Workspace, err error) {
	err = filepath.WalkDir("../", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// TODO: Corrigir isso antes de enviar o código para a master.
		if d.Name() == "Dockerfile" && strings.Contains(path, "jeffotoni/grow.standard.libray") {
			directory := strings.TrimSuffix(path, "Dockerfile")

			workspaces = append(workspaces, Workspace{
				entry:     d,
				directory: directory,
			})
		}
		return nil
	})
	return
}
