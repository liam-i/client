package engine

import (
	"testing"

	keybase1 "github.com/keybase/client/protocol/go"
)

func TestSearch(t *testing.T) {
	tc := SetupEngineTest(t, "btc")
	defer tc.Cleanup()

	ctx := &Context{
		LogUI: G.UI.GetLogUI(),
	}
	// This twitter handle is used by t_alice and t_charlie.
	e := NewSearchEngine(SearchEngineArgs{
		Query: "tacovontaco",
		// Asking for a lot of results hacks around the case where so many test
		// users have been created with these identities that the search reply
		// leaves out the ones we wanted.
		NumWanted: 100,
	}, tc.G)
	err := RunEngine(e, ctx)
	if err != nil {
		t.Fatal(err)
	}
	results := e.GetResults()

	if len(results) < 2 {
		// The test DB could contain other random test users, namely Max's ;)
		t.Fatalf("Expected at least 2 search results for 'tacovontaco'. Got %d.", len(results))
	}

	var alice keybase1.UserSummary
	var charlie keybase1.UserSummary
	for _, summary := range results {
		if summary.Username == "t_alice" {
			alice = summary
		} else if summary.Username == "t_charlie" {
			charlie = summary
		}
	}

	if alice.Username == "" {
		t.Fatal("Failed to find t_alice.")
	}
	if charlie.Username == "" {
		t.Fatal("Failed to find t_charlie.")
	}

	if len(alice.Proofs.Social) != 2 {
		t.Fatalf("Expected 2 proofs for t_alice, got %d.", len(alice.Proofs.Social))
	}
	if len(charlie.Proofs.Social) != 2 {
		t.Fatalf("Expected 2 proofs for t_charlie, got %d.", len(charlie.Proofs.Social))
	}
}
