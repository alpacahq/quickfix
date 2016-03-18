package relsymderivsecgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp1/secondarypricelimits"
)

//NoRelatedSym is a repeating group in RelSymDerivSecGrp
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//InstrumentExtension is a non-required component for NoRelatedSym.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//InstrmtLegGrp is a non-required component for NoRelatedSym.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
	//SecondaryPriceLimits is a non-required component for NoRelatedSym.
	SecondaryPriceLimits *secondarypricelimits.SecondaryPriceLimits
	//CorporateAction is a non-required field for NoRelatedSym.
	CorporateAction *string `fix:"292"`
}

//RelSymDerivSecGrp is a fix50sp1 Component
type RelSymDerivSecGrp struct {
	//NoRelatedSym is a non-required field for RelSymDerivSecGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *RelSymDerivSecGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }