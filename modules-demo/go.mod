module pelligent.in/modules-demo

go 1.16

require (
	google.golang.org/grpc v1.37.0
	pelligent.in/hellopb v0.0.0
)

replace pelligent.in/hellopb => ./demo/hellopb
