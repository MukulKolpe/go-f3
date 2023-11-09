package test

import (
	"github.com/anorth/f3sim/adversary"
	"github.com/anorth/f3sim/net"
	"github.com/anorth/f3sim/sim"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAbsent(t *testing.T) {
	for i := 0; i < 1000; i++ {
		sm := sim.NewSimulation(&sim.Config{
			HonestCount:  3,
			LatencySeed:  int64(i),
			LatencyMean:  0.100,
			GraniteDelta: 0.200,
		}, net.TraceNone)

		// Adversary has 1/4 of power.
		sm.SetAdversary(adversary.NewAbsent("A", sm.Network), 1)
		a := sm.Base.Extend(sm.CIDGen.Sample())
		sm.ReceiveChains(sim.ChainCount{len(sm.Participants), *a})

		ok := sm.Run()
		if !ok {
			sm.PrintResults()
		}
		require.True(t, ok)
	}
}