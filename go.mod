module github.com/delphix/dct-sdk-go

go 1.17

retract (
	v1.0.0-b1 // Published for testing
	v0.22.321 // Published for testing
)

require golang.org/x/oauth2 v0.0.0-20220309155454-6242fa91716a

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
