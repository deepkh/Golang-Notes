module main

go 1.15

replace addressbook => ../addressbook

require (
	addressbook v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.1
	google.golang.org/protobuf v1.25.0 // indirect
)
