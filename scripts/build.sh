CGO_ENABLED=1 \
go build -o libverifier.so -buildmode=c-shared .