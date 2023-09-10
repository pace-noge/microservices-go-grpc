package main

import (
	"time"

	userpb "github.com/pace-noge/microservices-go-rpc/user/userpb/user.pb.go"
)

var (
	timeout = time.Second
)

type server struct {
	userpb.UnimplementedUserserviceServer
}
