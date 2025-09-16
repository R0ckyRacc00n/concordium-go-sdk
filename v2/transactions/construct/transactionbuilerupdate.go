package construct

import v2 "github.com/Concordium/concordium-go-sdk/v2"

type transactionBuilderUpd struct {
	header  *v2.UpdateInstructionHeader
	payload *v2.UpdateInstructionPayload
	encoded *v2.RawPayload
}

// newTransactionBuilder is a constructor for transactionBuilder.
func newTransactionBuilderUpd(effectiveTime v2.TransactionTime, nonce v2.SequenceNumber,
	expiry v2.TransactionTime, payload *v2.UpdateInstructionPayload) *transactionBuilderUpd {
	encoded := payload.Payload.Encode()
	payloadSize := encoded.Size()
	header := &v2.UpdateInstructionHeader{
		SequenceNumber: (*v2.UpdateSequenceNumber)(&nonce),
		EffectiveTime:  &effectiveTime,
		Timeout:        &expiry,
		PayloadSize:    &payloadSize,
	}
	return &transactionBuilderUpd{
		header:  header,
		payload: payload,
		encoded: encoded,
	}
}

// size returns size of built transaction.
func (transactionBuilderUpd *transactionBuilderUpd) size() uint64 {
	return v2.UpdateInstructionHeaderSize + uint64(transactionBuilderUpd.header.PayloadSize.Value)
}

// construct builds PreAccountTransaction with updated energy amount by transmitted counting function.
func (transactionBuilderUpd *transactionBuilderUpd) constructUpd() *v2.PreUpdateInstruction {
	hashToSign, _ := v2.ComputeUpdateSignHash(transactionBuilderUpd.header,
		&v2.UpdateInstructionPayload{Payload: transactionBuilderUpd.encoded})

	return &v2.PreUpdateInstruction{
		Header:     transactionBuilderUpd.header,
		Payload:    transactionBuilderUpd.payload,
		Encoded:    transactionBuilderUpd.encoded,
		HashToSign: hashToSign,
	}
}

func makeTransactionUpd(effectiveTime v2.TransactionTime, nonce v2.SequenceNumber,
	expiry v2.TransactionTime, payload *v2.UpdateInstructionPayload) *v2.PreUpdateInstruction {
	builder := newTransactionBuilderUpd(effectiveTime, nonce, expiry, payload)

	return builder.constructUpd()
}
