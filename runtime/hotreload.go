package runtime

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// WatchFiles monitors .fg files for changes and calls reload func.
func WatchFiles(dir string, onChange func()) {
	stamps := make(map[string]time.Time)

	// Initial scan
	scanFG(dir, stamps)

	go func() {
		for {
			time.Sleep(1 * time.Second)

			changed := false
			filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if err != nil || info.IsDir() {
					return nil
				}
				if filepath.Ext(path) != ".fg" {
					return nil
				}

				mod := info.ModTime()
				if prev, ok := stamps[path]; ok {
					if mod.After(prev) {
						fmt.Printf("[flang] Arquivo modificado: %s\n", filepath.Base(path))
						changed = true
					}
				} else {
					// New file
					changed = true
				}
				stamps[path] = mod
				return nil
			})

			if changed {
				fmt.Println("[flang] Recarregando...")
				onChange()
			}
		}
	}()
}

func scanFG(dir string, stamps map[string]time.Time) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".fg" {
			stamps[path] = info.ModTime()
		}
		return nil
	})
}
