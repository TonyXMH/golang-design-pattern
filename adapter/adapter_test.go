package adapter

import "testing"

var except = "adaptee method"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	ret := adapter.Request()
	if ret != except {
		t.Fatalf("expect: %s, actual: %s", except, ret)
	}
}
