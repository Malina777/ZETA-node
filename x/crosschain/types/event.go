package types

const (
	// event key
	SubTypeKey = "SubTypeKey"
	CctxIndex  = "CctxIndex"

	Sender        = "Sender"
	SenderChain   = "SenderChain"
	TxOrigin      = "TxOrigin"
	InTxHash      = "InTxObservedHash"
	InBlockHeight = "InTxObservedBlockHeight"

	Receiver      = "Receiver"
	ReceiverChain = "ReceiverChain"
	OutTxHash     = "OutTxObservedHash"

	ZetaMint         = "ZetaMint"
	Amount           = "Amount"
	Asset            = "Asset"
	OutTXVotingChain = "OutTxVotingChain"
	OutBoundChain    = "OutBoundChain"
	OldStatus        = "OldStatus"
	NewStatus        = "NewStatus"
	StatusMessage    = "StatusMessage"
	RelayedMessage   = "RelayedMessage"
	Identifiers      = "LogIdentifiers"

	BallotIdentifier       = "BallotIdentifier"
	CCTXIndex              = "CCTXIndex"
	BallotObservationHash  = "BallotObservationHash"
	BallotObservationChain = "BallotObservationChain"
	BallotVotes            = "BallotVotes"
	BallotObservationType  = "BallotObservationType"
)

const (
	OutboundTxSuccessful = "crosschain/OutboundTxSuccessful"
	OutboundTxFailed     = "crosschain/OutboundTxFailed"
	CctxCreated          = "crosschain/CctxCreated"
	ZrcWithdrawCreated   = "crosschain/ZrcWithdrawCreated"
	BallotCreated        = "crosschain/BallotCreated"
	ZetaWithdrawCreated  = "crosschain/ZetaWithdrawCreated"
	InboundFinalized     = "crosschain/InboundFinalized"
	InboundFailed        = "crosschain/InboundFailed"
	StatusChanged        = "crosschain/StatusChanged"
	CctxFinalized        = "crosschain/CctxFinalized"
	CctxScrubbed         = "crosschain/CCTXScrubbed"
)
