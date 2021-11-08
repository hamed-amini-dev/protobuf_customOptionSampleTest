package heartbeatserver

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/ha-dev/protobuf_customOptionSampleTest/heartbeat_pb"
)

func TestMessage(t *testing.T) {
	var msg heartbeat_pb.HeartBeat

	_, md := descriptor.ForMessage(&msg)

	opts := md.Field[0].GetOptions()
	hea, err := proto.GetExtension(opts, E_heartbeatOptions)
	if err != nil {
		t.Errorf("failed: %v", err)
	} else {
		fmt.Println(hea)
	}
}
