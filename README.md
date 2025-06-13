# Go-MUD

A text-based MUD written in Go. Features multiplayer with WebRTC and Sockets.

# Running without executable

```sh
go run main.go engine.go
```

![Alt Text](/docs/showcase.png)

# Building Go-MUD

## Building with goreleaser

```sh
goreleaser release --snapshot --clean
```

Check the dist folder for the executables.

## Installing goreleaser

> Docs: https://goreleaser.com/install/

npm:
```sh
npm i -g @goreleaser/goreleaser
```

chocolatey:
```sh
choco install goreleaser
```

winget:
```sh
winget install goreleaser
```

apt:
```sh
sudo apt install goreleaser
```

go install:
```sh
go install github.com/goreleaser/goreleaser/v2@latest
```
