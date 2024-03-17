module github.com/rsturla/platform/services/greeter

go 1.21

require (
	github.com/rsturla/platform-contracts/gen/go v0.0.0-20240220223913-1c35ebf82913
	google.golang.org/grpc v1.61.0
	github.com/rsturla/platform/libs/grpc-helpers v0.0.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace (
	"github.com/rsturla/platform/libs/grpc-helpers" => "../../libs/grpc-helpers"
)
