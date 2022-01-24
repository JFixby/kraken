cd kraken

git pull

GO111MODULE=on

go build ./...

go test ./...

go test -v orderbook_test.go
