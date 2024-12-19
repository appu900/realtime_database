build:
   go build -o bin/realtimeDb cmd/main.go



run:
   ./bin/realtimeDb

test:
   go test -v ./...  