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

	envTemplate = "MINIO_ACCESS_KEY=%s\nMINIO_SECRET_KEY=%s\n"

	about = `minio-keygen (%s, %s)
Generates a MinIO access key and secret.You can redirect the output into an .env file.
`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	rawKey := make([]byte, keyLen)
	if _, err := rand.Read(rawKey); err != nil {
		panic(err)
	}

	rawSecret := make([]byte, secretLen)
	if _, err := rand.Read(rawSecret); err != nil {
		panic(err)
	}

	enc := base64.URLEncoding.WithPadding(pad).EncodeToString
	fmt.Printf(envTemplate, strings.ToUpper(enc(rawKey)), enc(rawSecret))
}

func usage() {
	ver := devVersion
	if info, ok := debug.ReadBuildInfo(); ok {
		ver = info.Main.Version
	}
	fmt.Printf(about, ver, runtime.Version())
}
