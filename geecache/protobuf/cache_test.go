package __

import (
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestRequest(t *testing.T) {
	request := &Request{
		Group: "scores",
		Key:   "Tom",
	}
	data, err := proto.Marshal(request)
	if err != nil {
		t.Fatal("marshaling error: ", err)
	}
	newRequest := &Request{}
	err = proto.Unmarshal(data, newRequest)
	if err != nil {
		t.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if request.GetGroup() != newRequest.GetGroup() || request.GetKey() != newRequest.GetKey() {
		t.Fatal("data mismatch")
	}
}
