module pelligent.in/grpc-demo

go 1.16

require (
	google.golang.org/grpc v1.37.0
	pelligent.in/grpc-demo/calculatorpb v0.0.0
)

replace pelligent.in/grpc-demo/calculatorpb => ./calculatorpb
