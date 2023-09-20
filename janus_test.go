package janus

import (
	"context"
	"testing"
)

func Test_Connect(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	


	
	client, _, err := Connect(ctx,"ws://localhost:8188/janus")
	if err != nil {
		t.Fail()
		return
	}





	mess, err := client.Info(ctx, "")
	if err != nil {
		t.Fail()
		return
	}
	t.Log(mess)

	sess, err := client.Create(ctx,"")
	if err != nil {
		t.Fail()
		return
	}
	t.Log(sess)
	t.Log("connect")
}
