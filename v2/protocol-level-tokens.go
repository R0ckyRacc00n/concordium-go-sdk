package v2

import (
	"context"
	"errors"
	"fmt"
	"github.com/Concordium/concordium-go-sdk/v2/pb"
	"io"
)

// TokenId Token ID: a unique symbol and identifier of a protocol level token.
type TokenId struct {
	// Between 1 and 128 characters, consisting of a-z, A-Z, 0-9, `.`, `%` and `-`.
	Value string
}

// TokenModuleRef A token module reference. Always 32 bytes.
type TokenModuleRef struct {
	Value []byte
}

// TokenAmount PLT amount representation. Actual amount = value * 10^(-decimals).
type TokenAmount struct {
	Value    uint64
	Decimals uint8
}

// TokenState Token state at the block level
type TokenState struct {
	TokenModuleRef TokenModuleRef
	Decimals       uint8
	TotalSupply    TokenAmount
	ModuleState    RawCBOR
}

// TokenAccountState Token state at the account level
type TokenAccountState struct {
	Balance     TokenAmount
	ModuleState *RawCBOR // optional
}

// TokenModuleEvent Single token event originating from a token module
type TokenModuleEvent struct {
	Type    string
	Details RawCBOR
}

// TokenHolder A token holder entity
type TokenHolder struct {
	// Currently only account supported
	Account *AccountAddress
}

// TokenTransferEvent An event emitted when a transfer of tokens is performed.
type TokenTransferEvent struct {
	From   TokenHolder
	To     TokenHolder
	Amount TokenAmount
	Memo   *Memo // optional
}

// TokenSupplyUpdateEvent An event emitted when the token supply is updated (mint/burn)
type TokenSupplyUpdateEvent struct {
	Target TokenHolder
	Amount TokenAmount
}

// TokenEvent Token event originating from token transactions.
type TokenEvent struct {
	TokenId       TokenId
	ModuleEvent   *TokenModuleEvent
	TransferEvent *TokenTransferEvent
	MintEvent     *TokenSupplyUpdateEvent
	BurnEvent     *TokenSupplyUpdateEvent
}

// TokenEffect Token events originating from token transactions.
type TokenEffect struct {
	Events []TokenEvent
}

// TokenModuleRejectReason Details provided by the token module in the event of rejecting a transaction.
type TokenModuleRejectReason struct {
	TokenId TokenId
	Type    string
	Details *RawCBOR // optional
}

// TokenCreationDetails Details about the creation of a protocol-level token.
type TokenCreationDetails struct {
	CreatePlt CreatePLT
	Events    []TokenEvent
}

// Convert protobuf TokenId → SDK TokenId
func convertTokenIdFromPB(pbToken *pb.TokenId) *TokenId {
	if pbToken == nil {
		return nil
	}
	return &TokenId{pbToken.Value}
}

// Convert SDK TokenId → protobuf
func convertTokenIdToPB(token *TokenId) *pb.TokenId {
	if token == nil {
		return nil
	}
	return &pb.TokenId{Value: token.Value}
}

func (c *Client) GetPLTList(ctx context.Context, blockHash BlockHashInputBest) ([]TokenId, error) {
	stream, err := c.GrpcClient.GetTokenList(ctx, convertBlockHashInput(blockHash))
	if err != nil {
		return nil, fmt.Errorf("failed to call GetTokenList: %w", err)
	}

	var tokens []TokenId
	for {
		token, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading token stream: %w", err)
		}
		tokens = append(tokens, *convertTokenIdFromPB(token))
	}

	return tokens, nil
}

func (c *Client) GetTokenInfo(ctx context.Context, blockHash BlockHashInput, tokenId *TokenId) (*pb.TokenInfo, error) {
	req := &pb.TokenInfoRequest{
		BlockHash: convertBlockHashInput(blockHash),
		TokenId:   convertTokenIdToPB(tokenId),
	}

	info, err := c.GrpcClient.GetTokenInfo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get token info: %w", err)
	}
	return info, nil
}
