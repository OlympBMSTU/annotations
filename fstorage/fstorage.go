package fstorage

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/OlympBMSTU/annotations/config"
	"github.com/OlympBMSTU/annotations/fstorage/result"
)

func FileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ComputeName(filename string) string {
	digest := md5.New()
	digest.Write([]byte(filename))
	hashBytes := digest.Sum(nil)

	z := new(big.Int)
	z.SetBytes(hashBytes)

	// see how correctly do this and there
	newPath := fmt.Sprintf("%x", z)
	newPath = newPath[:2] + "/" + newPath[2:4] + "/" + newPath[4:]
	return newPath
}

// todo refactor close open work eith strings
func WriteFile(fileHdr *multipart.FileHeader, args ...interface{}) result.FSResult {
	if fileHdr == nil {
		log.Print("No file sended")
		return result.ErrorResult(errors.New("No file presented"))
	}
	conf, _ := config.GetConfigInstance()


	name := ""
	if len(args) > 0 {
		name = string(args[0].(int))
	}

	ext := filepath.Ext(fileHdr.Filename)

	staticPath := conf.GetFileStorageName() + "/"
	filePathWithExt := ""
	newDirsPath := staticPath
	if len(name) > 0 {
		filePathWithExt = staticPath + name + ext
	} else {
		newNamePart := ComputeName(fileHdr.Filename)
		newDirsPath += newNamePart[:6]

		filePathWithExt := staticPath + newNamePart + ext
		idx := 1
		for {
			if FileExist(filePathWithExt) {
			newNamePart += strconv.Itoa(idx)
			filePathWithExt = staticPath + newNamePart + ext
			} else {
				break
			}
			idx++
		}
		name = newNamePart + ext
	}

	inFile, err := fileHdr.Open()
	defer inFile.Close()

	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}

	err = os.MkdirAll(newDirsPath, 0777)
	if err != nil {
		log.Println(err.Error())
		// clear dirs
		return result.ErrorResult(err)
	}

	f, err := os.Create(filePathWithExt)
	defer f.Close() // ? is it
	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}
	_, err = io.Copy(f, inFile)
	if err != nil {
		log.Println(err.Error())
		return result.ErrorResult(err)
	}
	return result.OkResult(name)
}
