package protobuf

import (
	"log"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestMarshalAndUnmarshal(t *testing.T) {
	test := &Test{
		Name:   proto.String("Hello"),
		Gender: proto.String("M"),
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Printf("Protobuf Error: %s", err.Error())
	}
	log.Printf("Marshal Result: %v", data)

	rtest := new(Test)
	err = proto.Unmarshal(data, rtest)
	if err != nil {
		log.Printf("Protobuf Error: %s", err.Error())
	}
	if *rtest.Name != "Hello" || *rtest.Gender != "M" {
		t.Fail()
	}
}
