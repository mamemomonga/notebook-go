module github.com/mamemomonga/notebook-go/api/docker

go 1.12

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 // indirect
	google.golang.org/grpc v1.21.1 // indirect
)

replace github.com/docker/docker v1.13.1 => github.com/docker/docker v0.7.3-0.20190626170757-a50869f077ea
