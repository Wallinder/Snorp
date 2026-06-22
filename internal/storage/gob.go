package storage

import (
	"encoding/gob"
	"os"
)

func saveGob(path string, object interface{}) error {
	file, err := os.Create(path + ".gob")
	if err != nil {
		return err
	}
	defer file.Close()

	return gob.NewEncoder(file).Encode(object)
}

func readGob(path string, object interface{}) error {
	file, err := os.Open(path + ".gob")
	if err != nil {
		return err
	}
	defer file.Close()

	return gob.NewDecoder(file).Decode(object)
}

func deleteGob(path string) error {
	return os.Remove(path + ".gob")
}
