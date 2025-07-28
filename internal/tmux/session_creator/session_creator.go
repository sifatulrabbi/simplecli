package sessioncreator

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

const (
	TmuxinatorConfigDir       = "/.config/tmuxinator/"
	DefaultProjectRootDir     = "/coding/"
	TestTmuxinatorConfigDir   = "/coding/personal-scripts/tmp/"
	TestDefaultProjectRootDir = "/coding/personal-scripts/tmp/"
)

type TmuxSessionCreator struct {
	TmuxinatorConfigLocation string
	NewProjectLocation       string
}

func New() *TmuxSessionCreator {
	return &TmuxSessionCreator{TmuxinatorConfigDir, DefaultProjectRootDir}
}

func (t *TmuxSessionCreator) CreateNewSession(projName string) error {
	if strings.Contains(projName, " ") {
		log.Fatalln("Project name can not contain spaces!")
	}
	projName = strings.ToLower(projName)

	usr, err := user.Current()
	if err != nil {
		log.Fatalln("Unable to get the current user information:", err)
	}
	if usr.HomeDir == "" {
		log.Fatalln("Sorry your home dir is not set!")
	}

	var (
		configDir       = usr.HomeDir + t.TmuxinatorConfigLocation
		txrProjFileName = configDir + projName + ".yml"
		projectDir      = usr.HomeDir + t.NewProjectLocation + projName
	)

	// making sure the tmuxinator config directory exists
	_, err = os.ReadDir(configDir)
	if err != nil {
		fmt.Println("tmuxinator config directory not found, now creating the directory...")
		if err = os.Mkdir(configDir, 0755); err != nil {
			log.Fatalln("Unable to create the config directory:", err)
		}
	}

	// creating the tmuxinator project config file
	newFile, err := os.OpenFile(txrProjFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer newFile.Close()

	// ensuring the new project's root dir
	if _, err := os.ReadDir(projectDir); err != nil {
		fmt.Println("Creating the project's root dir")
		if err = os.Mkdir(projectDir, 0755); err != nil {
			log.Fatalln("Failed to created the project directory:", err)
		}
	}

	// now writing the yaml config file for the new project
	content := fmt.Sprintf("name: %s\n", projName)
	content = fmt.Sprintf("%sroot: %q", content, projectDir)
	content = content + `
windows:
  - editor: nvim .
  - terminal:
  - run:
  - git: lazygit
`
	if _, err = newFile.Write([]byte(content)); err != nil {
		log.Fatalln(err)
	}

	return nil
}
