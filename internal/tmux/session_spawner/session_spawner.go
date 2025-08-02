// Package sessionspawner
package sessionspawner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func verifyProjectsDir() []os.DirEntry {
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("User root dir not found! Error:", err)
	}
	entries, err := os.ReadDir(userDir + "/.config/tmuxinator")
	if err != nil {
		log.Fatalln("Config directory not found! Error:", err)
	}
	return entries
}

func GetActiveTmuxSessions() []string {
	cmd := exec.Command("tmux", "list-sessions", "-F", "'#S'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("Error during output generation:", err)
	}
	projNames := strings.Split(strings.ReplaceAll(string(output), "'", ""), "\n")
	return projNames
}

func SpawnSessions() {
	projectEntries := verifyProjectsDir()
	for _, entry := range projectEntries {
		if entry.IsDir() {
			continue
		}
		projConfigName := entry.Name()
		nameSplit := strings.Split(projConfigName, ".")
		if len(nameSplit) < 2 {
			fmt.Println("WARN: Invalid project config file:", projConfigName)
		}
		fmt.Println("Opening project:", nameSplit[0])

		cmd := exec.Command("tmuxinator", "start", nameSplit[0], "--no-attach")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("ERROR: Unable to start the tmuxinator project:", err)
		}
		fmt.Println(string(output))
	}
}
