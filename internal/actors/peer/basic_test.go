package actors

import (
	"cjvirtucio87/distributed-todo-go/internal/rlog"
	"testing"
)

func TestSend(t *testing.T) {
	leader := NewBasicPeer(0)

	for i := 1; i < 3; i++ {
		leader.AddPeer(NewBasicPeer(i))
	}

	err := leader.Send(
		Message{
			Entries: []rlog.Entry{
				rlog.Entry{Command: "doFoo"},
			},
		},
	)
	
	if err != nil {
		t.Fatal(err)
	}

	leader = NewBasicPeer(0)
	for i := 1; i < 3; i++ {
		leader.AddPeer(
			&basicPeer {
				id: i,
				rlog: rlog.NewBasicLog(
					rlog.WithBackend(
						[]rlog.Entry{
							rlog.Entry{
								Command: "not supposed to be here",
							},
							rlog.Entry{
								Command: "not supposed to be here either",
							},
						},
					),
				),
				NextIndexMap: map[int]int{},
				peers:        []Peer{},
			},
		)
	}

	expectedEntry := rlog.Entry{Command: "doFoo"}
	err = leader.Send(
		Message{
			Entries: []rlog.Entry{
				expectedEntry,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	expectedPeerLogCount := 1
	for _, p := range leader.Followers() {
		actualPeerLogCount := p.LogCount()
		if expectedPeerLogCount != actualPeerLogCount {
			t.Fatalf("expectedPeerLogCount %d, was %d\n", expectedPeerLogCount, actualPeerLogCount)
		}
	}
}