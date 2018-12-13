// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

// SigHashType represents hash type bits at the end of a signature.
type SigHashType uint32

// Hash type bits from the end of a signature.
const (
	SigHashOld          SigHashType = 0x0
	SigHashAll          SigHashType = 0x1
	SigHashNone         SigHashType = 0x2
	SigHashSingle       SigHashType = 0x3
	SigHashAnyOneCanPay SigHashType = 0x80

	// sigHashMask defines the number of bits of the hash type which is used
	// to identify which outputs are signed.
	sigHashMask = 0x1f
)

// These are the constants specified for maximums in individual scripts.
const (
	MaxOpsPerScript       = 201 // Max number of non-push operations.
	MaxPubKeysPerMultiSig = 20  // Multisig can't have more sigs than this.
	MaxScriptElementSize  = 520 // Max bytes pushable to the stack.
)
