package v2

// Cbor A CBOR encoded bytestring
type Cbor struct {
	Value []byte
}

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
	Decimals uint32
}

// TokenState Token state at the block level
type TokenState struct {
	TokenModuleRef TokenModuleRef
	Decimals       uint32
	TotalSupply    TokenAmount
	ModuleState    Cbor
}

// TokenAccountState Token state at the account level
type TokenAccountState struct {
	Balance     TokenAmount
	ModuleState *Cbor // optional
}

// TokenModuleEvent Single token event originating from a token module
type TokenModuleEvent struct {
	Type    string
	Details Cbor
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
	Details *Cbor // optional
}

// CreatePLT Update payload for creating a new protocol-level token
type CreatePLT struct {
	TokenId                  TokenId
	TokenModule              TokenModuleRef
	Decimals                 uint32
	InitializationParameters Cbor
}

// TokenCreationDetails Details about the creation of a protocol-level token.
type TokenCreationDetails struct {
	CreatePlt CreatePLT
	Events    []TokenEvent
}
