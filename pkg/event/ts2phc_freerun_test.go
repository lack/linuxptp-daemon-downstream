package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTS2PHCFreerunCondition(t *testing.T) {
	const cfgName = "ts2phc.0.config"
	e := &EventHandler{
		data: map[string][]*Data{cfgName: {{
			ProcessName: TS2PHCProcessName,
			Details:     DDetails{{IFace: "ens0", State: PTP_FREERUN}},
		}}},
		clkSyncState:     map[string]*clockSyncState{cfgName: {leadingIFace: "ens0"}},
		LeadingClockData: &LeadingClockParams{},
	}
	e.LeadingClockData.toFreeRunThreshold = 1000
	assert.True(t, e.freeRunCondition(cfgName))
	e.data[cfgName][0].Details[0].State = PTP_LOCKED
	assert.False(t, e.freeRunCondition(cfgName))
}
