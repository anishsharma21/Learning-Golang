package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func fetchurl() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            logError("fetchurl: reading %s: %v", url, err)
            continue
        }
        body, err := io.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            logError("fetchurl: reading %s: %v", url, err)
            continue
        }
        fmt.Printf("%s\n", body)
    }
}

func logError(format string, v ...interface{}) {
    _, file, line, _ := runtime.Caller(1)
    parts := strings.Split(filepath.ToSlash(file), "/")
    shortFile := strings.Join(parts[len(parts)-2:], "/")
    log.Printf("%s:%d:%s", shortFile, line, fmt.Sprintf(format, v...))
}