// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

const (
	// MaxStackSize is the maximum combined height of stack and alt stack
	// during execution.
	MaxStackSize = 1000

	// MaxScriptSize is the maximum allowed length of a raw script.
	MaxScriptSize = 10000

	// payToWitnessPubKeyHashDataSize is the size of the witness program's
	// data push for a pay-to-witness-pub-key-hash output.
	payToWitnessPubKeyHashDataSize = 20

	// payToWitnessScriptHashDataSize is the size of the witness program's
	// data push for a pay-to-witness-script-hash output.
	payToWitnessScriptHashDataSize = 32
)
