#!/bin/bash

env GOOS=js GOARCH=wasm go build -o snake.wasm github.com/wpcasarin/ebiten-snake