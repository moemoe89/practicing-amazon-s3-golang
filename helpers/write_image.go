package helpers

import (
	"encoding/base64"
	"os"
	"strconv"
	"time"
)

func WriteImage(encodedImage string, directoryName string) (string,error) {

	var i int
	fileExtension := ".png"
	fileName := "IMAGE-"+time.Now().Format("20060102150405")

	data,err := base64.StdEncoding.DecodeString(encodedImage)
	if err != nil {
		return "",err
	}

	for i = 0;; {
		i++
		fileName = fileName+strconv.Itoa(i)
		if _, err := os.Stat(directoryName+"/"+fileName+fileExtension); os.IsNotExist(err) {
			break
		}
	}

	file,err := os.Create(directoryName+"/"+fileName+fileExtension)
	if err != nil {
		return "",err
	}

	defer file.Close()

	_,err = file.Write(data)
	if err != nil {
		return "",err
	}

	return fileName+fileExtension,nil

}