package actors

import (
	"fmt"
	"testing"
)

func TestAddPeer(t *testing.T) {
	leader := NewBasicPeer(0)

	leader.AddPeer(NewBasicPeer(1))

	expectedCount := 1
	actualCount := leader.PeerCount()

	if expectedCount != actualCount {
		t.Error(fmt.Printf("expectedCount %d, was %d", expectedCount, actualCount))
	}
}

func TestSend(t *testing.T) {
	leader := NewBasicPeer(0)

	for i := 1; i < 3; i++ {
		leader.AddPeer(NewBasicPeer(i))
	}

	expectedSendResult := true
	actualSendResult := leader.Send(
		Message{
			Entries: []Entry{
				Entry{Command: "doFoo"},
			},
		},
	)

	if expectedSendResult != actualSendResult {
		t.Error(fmt.Printf("expectedSendResult %t, was %t", expectedSendResult, actualSendResult))
	}

	leader = NewBasicPeer(0)
	for i := 1; i < 3; i++ {
		leader.AddPeer(
			&basicPeer{
				id: i,
				log: []Entry{
					Entry{
						Command: "not supposed to be here",
					},
					Entry{
						Command: "not supposed to be here either",
					},
				},
				NextIndexMap: map[int]int{},
				peers:        []Peer{},
			},
		)
	}

	expectedEntry := Entry{Command: "doFoo"}
	actualSendResult = leader.Send(
		Message{
			Entries: []Entry{
				expectedEntry,
			},
		},
	)

	if expectedSendResult != actualSendResult {
		t.Error(fmt.Printf("expectedSendResult %t, was %t", expectedSendResult, actualSendResult))
	}

	expectedPeerLogCount := 1

	for _, p := range leader.Followers() {
		actualPeerLogCount := p.LogCount()

		if expectedPeerLogCount != actualPeerLogCount {
			t.Error(fmt.Printf("expectedPeerLogCount %d, was %d", expectedPeerLogCount, actualPeerLogCount))
		}

		actualPeerEntry := p.Entry(0)

		if expectedEntry != actualPeerEntry {
			t.Error(fmt.Printf("expectedEntry %v, was %v", expectedEntry.Command, actualPeerEntry.Command))
		}
	}
}
