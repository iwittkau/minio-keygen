package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

const (
	keyLen    = 18
	secretLen = 36
	pad       = '+'

	devVersion = "dev"

	about = `minio-keygen (%s, %s)
Generates a MinIO access key and secret. You can redirect the output into an .env file.
`
)

func main() {

	flag.Usage = usage
	flag.Parse()
	var (
		rawKey    = make([]byte, keyLen)
		rawSecret = make([]byte, secretLen)
		err       error
	)
	_, err = rand.Read(rawKey)
	if err != nil {
		panic(err)
	}

	_, err = rand.Read(rawSecret)
	if err != nil {
		panic(err)
	}

	enc := base64.URLEncoding.WithPadding(pad)

	fmt.Printf(
		"MINIO_ACCESS_KEY=%s\nMINIO_SECRET_KEY=%s\n",
		strings.ToUpper(enc.EncodeToString(rawKey)),
		enc.EncodeToString(rawSecret),
	)
}

func usage() {
	var (
		info, ok = debug.ReadBuildInfo()
		ver      = devVersion
	)
	if !ok {
		ver = info.Main.Version
	}
	fmt.Printf(about, ver, runtime.Version())
}
