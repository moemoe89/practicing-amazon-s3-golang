package helpers

import (
	"os"
)

func DeleteFile(path string) (error){
	var err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}