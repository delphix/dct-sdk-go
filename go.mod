module github.com/delphix/dct-sdk-go

go 1.17

require golang.org/x/oauth2 v0.0.0-20220309155454-6242fa91716a

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

retract (
	v1.1.1 // Contains retractions only.
	v1.1.1-beta3 // Published for testing
	v1.1.0 // Published for testing
	v1.0.0-beta2 // Published for testing
	v1.0.0-beta // Published for testing
	v1.0.0-b1 // Published for testing
)
