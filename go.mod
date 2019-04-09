module go-grpc

go 1.12

require (
	github.com/golang/protobuf v1.3.1
	github.com/satori/go.uuid v1.2.0
	google.golang.org/grpc v1.16.0
)

replace (
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
