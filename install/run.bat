cd kraken

git pull

set GO111MODULE=on

go build ./...

go test ./...

go test -v orderbook_test.go
