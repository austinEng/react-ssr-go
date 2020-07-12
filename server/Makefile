snapshot:
	touch ./src/client_snap/build.rs
	cargo build -vv --release --manifest-path ./src/client_snap/Cargo.toml --out-dir ./src/client_snap/gen/ -Z unstable-options

generate:
	GOPATH=$(realpath ./) go generate src/app/main.go

start:
	GOPATH=$(realpath ./) go run src/app/main.go ../client/dist/

build:
	GOPATH=$(realpath ./) go build src/app/main.go
