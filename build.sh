#!/usr/bin/env bash

# Gofs build script for Unix-like system

function die() {
  echo -e "\033[01;31;01m $*\033[01;00;00m"
  exit 1
}

function checkDependencies() {
  if [ $# -lt 1 ]; then
    return
  fi

  for arg in $@ ; do
    which ${arg} >/dev/null 2>&1 || die "No '${arg}' command could be found in your PATH."
  done
}

# ---*--- main ---*---
checkDependencies go
go build -o bin/gofs src/*.go
