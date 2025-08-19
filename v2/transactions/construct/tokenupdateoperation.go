package construct

import (
	v2 "github.com/Concordium/concordium-go-sdk/v2"
	"github.com/fxamacker/cbor/v2"
)

func TokenUpdateOperation(numSigs uint32, sender v2.AccountAddress, nonce v2.SequenceNumber, expiry v2.TransactionTime,
	tokenId v2.TokenID, operations v2.TokenOperations,
) *v2.PreAccountTransaction {
	energy := TokenOperationsTxnEnergy(operations)
	addEnergy := v2.AddEnergy{Energy: energy, NumSigs: numSigs}
	givenEnergy := &v2.GivenEnergy{Energy: &addEnergy}

	rawOps, err := MarshalTokenOperationsToCBORArray(operations)
	if err != nil {
		panic("failed to marshal operations: " + err.Error())
	}

	payload := &v2.AccountTransactionPayload{Payload: &v2.TokenUpdate{Payload: &v2.TokenOperationsPayload{TokenId: tokenId,
		Operations: rawOps}}}

	return makeTransaction(sender, nonce, expiry, givenEnergy, payload)
}

func TokenOperationsTxnEnergy(ops v2.TokenOperations) v2.Energy {
	total := v2.PltOperationsTransactions
	for _, op := range ops {
		total += op.TxnEnergy().Value
	}
	return v2.Energy{Value: total}
}

func MarshalTokenOperationsToCBORArray(ops v2.TokenOperations) (v2.RawCBOR, error) {
	var encodedItems [][]byte
	for _, op := range ops {
		b, err := op.EncodeCBOR()
		if err != nil {
			return v2.RawCBOR{}, err
		}
		encodedItems = append(encodedItems, b)
	}

	finalCBOR, err := cbor.Marshal(encodedItems)
	if err != nil {
		return v2.RawCBOR{}, err
	}

	return v2.RawCBOR{Bytes: finalCBOR}, nil
}
