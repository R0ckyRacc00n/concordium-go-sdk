package send

import (
	v2 "github.com/Concordium/concordium-go-sdk/v2"
	"github.com/Concordium/concordium-go-sdk/v2/transactions/construct"
)

// TokenUpdateOperation
// / Construct and sign a protocol level token update transaction consisting
// / of the token update operations encoded in the given CBOR.
func TokenUpdateOperation(signer v2.UpdateSigner, nonce v2.SequenceNumber, effectiveTime v2.TransactionTime,
	expiry v2.TransactionTime, tokenId v2.TokenID, operations v2.TokenOperations) (*v2.UpdateInstruction, error) {
	return construct.TokenUpdateOperation(nonce, effectiveTime, expiry, tokenId, operations).Sign(signer)
}
