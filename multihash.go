package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash/crc32"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	args := os.Args[1:]
	for _, path := range args {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			GetFileInfo(removeTrailingPathSeparator(path))
		}(path)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\n# %s runtime\n", elapsed)
}

func GetFileInfo(src string) {
	file, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	fileInfo, err := os.Lstat(src)
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		return
	}

	if fileInfo.IsDir() {
		entries, err := file.ReadDir(0)
		if err != nil {
			log.Fatal(err)
		}
		for _, e := range entries {
			GetFileInfo(fmt.Sprintf("%s/%s", src, e.Name()))
		}
	} else {
		readFile, err := os.ReadFile(src)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n# %s\n", src)
		fmt.Printf("size:%d mtime:%s  %s\n", fileInfo.Size(), fileInfo.ModTime().Format("15:04:05 2006-02-01"), src)
		fmt.Printf("CRC32:%x  %s\n", crc32sum(readFile), src)
		fmt.Printf("MD5:%x  %s\n", md5sum(readFile), src)
		fmt.Printf("SHA1:%x  %s\n", sha1sum(readFile), src)
		fmt.Printf("SHA256:%x  %s\n", sha256sum(readFile), src)
		fmt.Printf("SHA512:%x  %s\n", sha512sum(readFile), src)
	}
}

func crc32sum(s []byte) uint32 {
	crc32q := crc32.MakeTable(crc32.IEEE)
	return crc32.Checksum(s, crc32q)
}

func md5sum(s []byte) []byte {
	md5 := md5.New()
	md5.Write(s)
	return md5.Sum(nil)
}

func sha1sum(s []byte) []byte {
	sha1 := sha1.New()
	sha1.Write(s)
	return sha1.Sum(nil)
}

func sha256sum(s []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(s)
	return sha256.Sum(nil)
}

func sha512sum(s []byte) []byte {
	sha512 := sha512.New()
	sha512.Write(s)
	return sha512.Sum(nil)
}

func removeTrailingPathSeparator(p string) string {
	separator := fmt.Sprintf("%c", os.PathSeparator)
	return strings.TrimSuffix(p, separator)
}
