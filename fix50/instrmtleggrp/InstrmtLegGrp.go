package instrmtleggrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentleg"
)

//NoLegs is a repeating group in InstrmtLegGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

//InstrmtLegGrp is a fix50 Component
type InstrmtLegGrp struct {
	//NoLegs is a non-required field for InstrmtLegGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }