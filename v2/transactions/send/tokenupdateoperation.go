package send

import (
	v2 "github.com/Concordium/concordium-go-sdk/v2"
	"github.com/Concordium/concordium-go-sdk/v2/transactions/construct"
)

// TokenUpdateOperation
// / Construct and sign a protocol level token update transaction consisting
// / of the token update operations encoded in the given CBOR.
func TokenUpdateOperation(signer v2.ExactSizeTransactionSigner, sender v2.AccountAddress, nonce v2.SequenceNumber,
	expiry v2.TransactionTime, tokenId v2.TokenID, operations v2.TokenOperations) (*v2.AccountTransaction, error) {
	return construct.TokenUpdateOperation(signer.NumberOfKeys(), sender, nonce, expiry, tokenId, operations).Sign(signer)
}
