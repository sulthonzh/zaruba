package watcher

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/state-alchemists/zaruba/dir"
	"github.com/state-alchemists/zaruba/organizer"
)

// Watch projectDir
func Watch(projectDir string, stopChan chan bool, arguments ...string) (err error) {
	projectDir, err = filepath.Abs(projectDir)
	if err != nil {
		return
	}
	// perform organize
	err = organizer.Organize(projectDir, arguments...)
	if err != nil {
		return
	}
	go listen(projectDir, arguments...)
	// wait until stopped
	<-stopChan
	return
}

func listen(projectDir string, arguments ...string) {
	// get allDirs
	allDirs, err := dir.GetAllDirs(projectDir)
	for err != nil {
		log.Printf("[ERROR] Fail to get list of directories: %s. Retrying...", err)
		allDirs, err = dir.GetAllDirs(projectDir)
	}
	// create watcher, don't give up
	watcher, err := fsnotify.NewWatcher()
	for err != nil {
		log.Printf("[ERROR] Fail to create watcher: %s. Retrying...", err)
		watcher, err = fsnotify.NewWatcher()
	}
	defer watcher.Close()
	// add allDirs to watcher
	addDirsToWatcher(watcher, allDirs)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				continue
			}
			log.Printf("[INFO] Detect event: %s", event)
			removeDirsFromWatcher(watcher, allDirs)
			organizer.Organize(projectDir, arguments...)
			allDirs, err = dir.GetAllDirs(projectDir)
			for err != nil {
				allDirs, err = dir.GetAllDirs(projectDir)
			}
			addDirsToWatcher(watcher, allDirs)
		case err, ok := <-watcher.Errors:
			if !ok {
				continue
			}
			log.Printf("[ERROR] Watcher error: %s. Continue to listen...", err)
		}
	}
}

func removeDirsFromWatcher(watcher *fsnotify.Watcher, allDirs []string) {
	for _, dirPath := range allDirs {
		watcher.Remove(dirPath)
	}
}

func addDirsToWatcher(watcher *fsnotify.Watcher, allDirs []string) {
	for _, dirPath := range allDirs {
		watcher.Add(dirPath)
	}
}