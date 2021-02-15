package log

import (
	"fyne.io/fyne/v2"
	"os"
	"path/filepath"
)

// storagePath returns the location of the settings storage
func storagePath(a fyne.App, fileName string) string {
	return filepath.Join(storageRoot(a), `files`, fileName)
}

// storageRoot returns the location of the app storage
func storageRoot(a fyne.App) string {
	return filepath.Join(rootConfigDir(), a.UniqueID())
}

func rootConfigDir() string {
	return "/data/data"
	filesDir := os.Getenv("FILESDIR")
	if filesDir == "" {
		//log.Println("FILESDIR env was not set by android native code")
		//return "/data/data" // probably won't work, but we can't make a better guess
	}
	return filesDir
}
