package utils

import (
	"encoding/json"
	"fmt"
	"runtime"
	"testing"

	"github.com/thrasher-corp/gocryptotrader/exchanges/binance"
)

func TestAdjustGoMaxProcs(t *testing.T) {
	// Test default settings
	curr := runtime.GOMAXPROCS(-1)
	numCPUs := runtime.NumCPU()

	// This func both checks for an error of AdjustGoMaxProcs, plus
	// ensures that the value it sets is the one that is expected
	checker := func(setting, expected int) error {
		if err := AdjustGoMaxProcs(setting); err != nil {
			return err
		}
		if i := runtime.GOMAXPROCS(expected); i != expected {
			return fmt.Errorf("expected %d, got %d", expected, i)
		}
		return nil
	}

	tester := []struct {
		Setting  int
		Expected int
	}{
		{
			// Test setting to current runtime val
			Setting:  curr,
			Expected: curr,
		},
		{
			// Test setting to num of logical CPUs
			Setting:  numCPUs,
			Expected: numCPUs,
		},
		{
			// Test crazy value and make sure it defaults to numCPUs
			Setting:  1000,
			Expected: numCPUs,
		},
		{
			// Test another crazy value and make sure it defaults to numCPUs
			Setting:  -1,
			Expected: numCPUs,
		},
	}

	for x := range tester {
		if err := checker(tester[x].Setting, tester[x].Expected); err != nil {
			t.Errorf("%d failed. %s", x, err)
		}
	}
}

func TestRound(t *testing.T) {
	d := []byte("[\n    {\n        \"id\": \"d6981a9a9c78426f959f644769a619fa\",\n        \"amount\": \"70.832\",\n        \"transactionFee\": \"0.018\",\n        \"coin\": \"NEAR\",\n        \"status\": 6,\n        \"address\": \"deltabottest.near\",\n        \"txId\": \"CMiNShYtkEdvuVJyje8vjCUfzi2FzwSnwedJkzpyVYBE\",\n        \"applyTime\": \"2025-11-28 09:45:08\",\n        \"network\": \"NEAR\",\n        \"transferType\": 0,\n        \"info\": \"9c484fa5d2d069569ba063fc555c34e621ccd88fdbb0295fc79bad232621c5c1,106259909335036\",\n        \"confirmNo\": 20,\n        \"walletType\": 0,\n        \"txKey\": \"\",\n        \"completeTime\": \"2025-11-28 09:46:14\"\n    }\n]")
	var l []binance.WithdrawStatusResponse
	err := json.Unmarshal(d, &l)
	if err != nil {
		fmt.Println(err)
		return
	}
}
