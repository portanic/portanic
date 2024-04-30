package plugins

import (
	"fmt"
	"os"
	"path/filepath"
)

func DiscoverPlugins(pluginDirs []string) []string {
	var plugins []string
	for _, pluginDir := range pluginDirs {
		files, err := os.ReadDir(pluginDir)
		if err != nil {
			fmt.Println("Failed to read directory:", pluginDir, err.Error())
			continue
		}

		for _, file := range files {
			if file.Type().IsRegular() {
				// Assuming the executable bit is set for plugins
				if file.Type().Perm()&0111 != 0 {
					pluginPath := filepath.Join(pluginDir, file.Name())
					plugins = append(plugins, pluginPath)
				}
			}
		}
	}
	return plugins
}
