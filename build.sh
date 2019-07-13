#!/bin/sh

set -e

# Build for Linux amd64
GO111MODULE=on \
GOOS=linux \
GOARCH=arm64 \
  go build \
    -o bin/pinatapinner_linux_amd64

# Build for Linux 386
GO111MODULE=on \
GOOS=linux \
GOARCH=386 \
  go build \
    -o bin/pinatapinner_linux_386

# Build for Linux arm (v6)
GO111MODULE=on \
GOOS=linux \
GOARCH=arm \
GOARM=6 \
  go build \
    -o bin/pinatapinner_linux_arm

# Build for Linux arm64
GO111MODULE=on \
GOOS=linux \
GOARCH=arm64 \
  go build \
    -o bin/pinatapinner_linux_arm64

# Build for macOS
GO111MODULE=on \
GOOS=darwin \
GOARCH=amd64 \
  go build \
    -o bin/pinatapinner_macos

# Build for Windows (64-bit)
GO111MODULE=on \
GOOS=windows \
GOARCH=amd64 \
  go build \
    -o bin/pinatapinner_win64.exe
