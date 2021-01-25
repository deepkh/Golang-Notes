module main

go 1.15

replace example.com/04_grpcftp => ../protos

require (
	example.com/04_grpcftp v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.35.0 // indirect
)
