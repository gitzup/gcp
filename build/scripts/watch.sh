#!/usr/bin/env bash

INOTIFY=$(which inotifywait 2>/dev/null)
[[ -z "${INOTIFY}" ]] && echo "Install inotify-tools first." >&2 && exit 1

make docker

while true; do
    EVENT=$(inotifywait -e create,modify,delete -r -q ./main.go ./Makefile ./build/Dockerfile ./api ./cmd ./internal)
    echo >&2
    echo "=========================================" >&2
    echo "CHANGE DETECTED: ${EVENT}" >&2
    echo "=========================================" >&2
    echo >&2
    make docker
done
