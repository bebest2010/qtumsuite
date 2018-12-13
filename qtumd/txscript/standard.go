// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

import (
	"fmt"

	qtumutil "github.com/bebest2010/qtumsuite/qtumutil"
)

// MultiSigScript returns a valid script for a multisignature redemption where
// nrequired of the keys in pubkeys are required to have signed the transaction
// for success.  An Error with the error code ErrTooManyRequiredSigs will be
// returned if nrequired is larger than the number of keys provided.
func MultiSigScript(pubkeys []*qtumutil.AddressPubKey, nrequired int) ([]byte, error) {
	if len(pubkeys) < nrequired {
		str := fmt.Sprintf("unable to generate multisig script with "+
			"%d required signatures when there are only %d public "+
			"keys available", nrequired, len(pubkeys))
		return nil, scriptError(ErrTooManyRequiredSigs, str)
	}

	builder := NewScriptBuilder().AddInt64(int64(nrequired))
	for _, key := range pubkeys {
		builder.AddData(key.ScriptAddress())
	}
	builder.AddInt64(int64(len(pubkeys)))
	builder.AddOp(OP_CHECKMULTISIG)

	return builder.Script()
}
