package path

import (
	"os"
	"test/src/tools/env"
	"test/src/tools/log"
)

func MkDir(dir string) error {
	err := os.Mkdir(dir, os.ModeDir)

	if os.IsExist(err) || err == nil {
		return nil
	}

	return err
}

func StorageDir() string {
	dir, err := os.UserCacheDir()

	if err != nil {
		log.Fatal(err)
	}

	dir = dir + "\\" + env.Get("APP_NAME")

	err = MkDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	return dir
}
