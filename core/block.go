package core

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/core/pb"
	"github.com/medibloc/go-medibloc/crypto/signature"
	"github.com/medibloc/go-medibloc/crypto/signature/algorithm"
	"github.com/medibloc/go-medibloc/storage"
	byteutils "github.com/medibloc/go-medibloc/util/bytes"
	"github.com/medibloc/go-medibloc/util/logging"
	"golang.org/x/crypto/sha3"
)

// BlockHeader is block header
type BlockHeader struct {
	hash       common.Hash
	parentHash common.Hash

	accsRoot  common.Hash
	txsRoot   common.Hash
	usageRoot common.Hash

	coinbase  common.Address
	timestamp int64
	chainID   uint32

	alg  algorithm.Algorithm
	sign []byte
}

// ToProto converts BlockHeader to corepb.BlockHeader
func (b *BlockHeader) ToProto() (proto.Message, error) {
	return &corepb.BlockHeader{
		Hash:       b.hash.Bytes(),
		ParentHash: b.parentHash.Bytes(),
		AccsRoot:   b.accsRoot.Bytes(),
		TxsRoot:    b.txsRoot.Bytes(),
		UsageRoot:  b.usageRoot.Bytes(),
		Coinbase:   b.coinbase.Bytes(),
		Timestamp:  b.timestamp,
		ChainId:    b.chainID,
		Alg:        uint32(b.alg),
		Sign:       b.sign,
	}, nil
}

// FromProto converts corepb.BlockHeader to BlockHeader
func (b *BlockHeader) FromProto(msg proto.Message) error {
	if msg, ok := msg.(*corepb.BlockHeader); ok {
		b.hash = common.BytesToHash(msg.Hash)
		b.parentHash = common.BytesToHash(msg.ParentHash)
		b.accsRoot = common.BytesToHash(msg.AccsRoot)
		b.txsRoot = common.BytesToHash(msg.TxsRoot)
		b.usageRoot = common.BytesToHash(msg.UsageRoot)
		b.coinbase = common.BytesToAddress(msg.Coinbase)
		b.timestamp = msg.Timestamp
		b.chainID = msg.ChainId
		b.alg = algorithm.Algorithm(msg.Alg)
		b.sign = msg.Sign
		return nil
	}
	return ErrInvalidProtoToBlockHeader
}

// Block represents a block
type Block struct {
	header       *BlockHeader
	transactions Transactions

	storage storage.Storage
	state   *BlockState

	sealed bool
	height uint64
}

// ToProto converts Block to corepb.Block
func (block *Block) ToProto() (proto.Message, error) {
	header, err := block.header.ToProto()
	if err != nil {
		return nil, err
	}
	if header, ok := header.(*corepb.BlockHeader); ok {
		txs := make([]*corepb.Transaction, len(block.transactions))
		for idx, v := range block.transactions {
			tx, err := v.ToProto()
			if err != nil {
				return nil, err
			}
			if tx, ok := tx.(*corepb.Transaction); ok {
				txs[idx] = tx
			} else {
				return nil, ErrCannotConvertTransaction
			}
		}
		return &corepb.Block{
			Header:       header,
			Transactions: txs,
			Height:       block.height,
		}, nil
	}
	return nil, ErrInvalidBlockToProto
}

// FromProto converts corepb.Block to Block
func (block *Block) FromProto(msg proto.Message) error {
	if msg, ok := msg.(*corepb.Block); ok {
		block.header = new(BlockHeader)
		if err := block.header.FromProto(msg.Header); err != nil {
			return err
		}

		block.transactions = make(Transactions, len(msg.Transactions))
		for idx, v := range msg.Transactions {
			tx := new(Transaction)
			if err := tx.FromProto(v); err != nil {
				return err
			}
			block.transactions[idx] = tx
		}
		block.height = msg.Height
		return nil
	}
	return ErrInvalidProtoToBlock
}

// NewBlock initialize new block pointing parent
func NewBlock(chainID uint32, coinbase common.Address, parent *Block) (*Block, error) {
	state, err := parent.state.Clone()
	if err != nil {
		return nil, err
	}

	block := &Block{
		header: &BlockHeader{
			parentHash: parent.Hash(),
			coinbase:   coinbase,
			timestamp:  time.Now().Unix(),
			chainID:    chainID,
		},
		transactions: make(Transactions, 0),
		storage:      parent.storage,
		state:        state,
		height:       parent.height + 1,
		sealed:       false,
	}

	return block, nil
}

// ChainID returns chain id
func (block *Block) ChainID() uint32 {
	return block.header.chainID
}

// Coinbase returns coinbase address
func (block *Block) Coinbase() common.Address {
	return block.header.coinbase
}

// Alg returns sign altorithm
func (block *Block) Alg() algorithm.Algorithm {
	return block.header.alg
}

// Signature returns block signature
func (block *Block) Signature() []byte {
	return block.header.sign
}

// Timestamp returns timestamp
func (block *Block) Timestamp() int64 {
	return block.header.timestamp
}

// SetTimestamp set block timestamp
func (block *Block) SetTimestamp(timestamp int64) error {
	if block.sealed {
		return ErrBlockAlreadySealed
	}
	block.header.timestamp = timestamp
	return nil
}

// Hash returns block hash
func (block *Block) Hash() common.Hash {
	return block.header.hash
}

// ParentHash returns hash of parent block
func (block *Block) ParentHash() common.Hash {
	return block.header.parentHash
}

// State returns block state
func (block *Block) State() *BlockState {
	return block.state
}

// AccountsRoot returns root hash of accounts trie
func (block *Block) AccountsRoot() common.Hash {
	return block.header.accsRoot
}

// TransactionsRoot returns root hash of txs trie
func (block *Block) TransactionsRoot() common.Hash {
	return block.header.txsRoot
}

// Height returns height
func (block *Block) Height() uint64 {
	return block.height
}

// Transactions returns txs in block
func (block *Block) Transactions() Transactions {
	return block.transactions
}

// SetTransactions sets transactions TO BE REMOVED: For test without block pool
func (block *Block) SetTransactions(txs Transactions) error {
	block.transactions = txs
	return nil
}

// Storage returns storage used by block
func (block *Block) Storage() storage.Storage {
	return block.storage
}

// Sealed returns sealed
func (block *Block) Sealed() bool {
	return block.sealed
}

// Seal writes state root hashes and block hash in block header
func (block *Block) Seal() error {
	if block.sealed {
		return ErrBlockAlreadySealed
	}

	block.header.accsRoot = block.state.AccountsRoot()
	block.header.txsRoot = block.state.TransactionsRoot()

	var err error
	block.header.hash, err = HashBlock(block)
	if err != nil {
		return err
	}
	block.sealed = true
	return nil
}

// HashBlock returns hash of block
func HashBlock(block *Block) (common.Hash, error) {
	if block == nil {
		return common.Hash{}, ErrNilArgument
	}

	hasher := sha3.New256()

	hasher.Write(block.ParentHash().Bytes())
	hasher.Write(block.header.coinbase.Bytes())
	hasher.Write(byteutils.FromInt64(block.Timestamp()))
	hasher.Write(byteutils.FromUint32(block.ChainID()))

	for _, tx := range block.transactions {
		hasher.Write(tx.Hash().Bytes())
	}

	return common.BytesToHash(hasher.Sum(nil)), nil
}

// ExecuteTransaction on given block state
func (block *Block) ExecuteTransaction(tx *Transaction) error {
	return block.state.ExecuteTx(tx)
}

// VerifyExecution executes txs in block and verify root hashes using block header
func (block *Block) VerifyExecution() error {
	block.BeginBatch()

	if err := block.ExecuteAll(); err != nil {
		block.RollBack()
		return err
	}

	if err := block.VerifyState(); err != nil {
		block.RollBack()
		return err
	}

	block.Commit()

	return nil
}

// ExecuteAll executes all txs in block
func (block *Block) ExecuteAll() error {
	for _, tx := range block.transactions {
		if err := block.ExecuteTransaction(tx); err != nil {
			return err
		}

		if err := block.AcceptTransaction(tx); err != nil {
			return err
		}
	}

	return nil
}

// AcceptTransaction adds tx in block state
func (block *Block) AcceptTransaction(tx *Transaction) error {
	return block.state.AcceptTransaction(tx, block.Timestamp())
}

// VerifyState verifies block states comparing with root hashes in header
func (block *Block) VerifyState() error {
	if !byteutils.Equal(block.state.AccountsRoot().Bytes(), block.AccountsRoot().Bytes()) {
		return ErrInvalidBlockAccountsRoot
	}
	if !byteutils.Equal(block.state.TransactionsRoot().Bytes(), block.TransactionsRoot().Bytes()) {
		return ErrInvalidBlockTxsRoot
	}
	return nil
}

// SignThis sets signature info in block
func (block *Block) SignThis(signer signature.Signature) error {
	if !block.Sealed() {
		return ErrBlockNotSealed
	}

	sig, err := signer.Sign(block.header.hash.Bytes())
	if err != nil {
		return err
	}
	block.header.alg = signer.Algorithm()
	block.header.sign = sig
	return nil
}

// VerifyIntegrity verifies if block signature is valid
func (block *Block) VerifyIntegrity() error {
	for _, tx := range block.transactions {
		if err := tx.VerifyIntegrity(block.header.chainID); err != nil {
			return err
		}
	}

	wantedHash, err := HashBlock(block)
	if err != nil {
		return err
	}
	if !wantedHash.Equals(block.header.hash) {
		return ErrInvalidBlockHash
	}

	// TODO: Verify according to consensus algorithm

	return nil
}

// BeginBatch makes block state update possible
func (block *Block) BeginBatch() error {
	if err := block.state.BeginBatch(); err != nil {
		return err
	}
	return nil
}

// RollBack rolls back block state batch updates
func (block *Block) RollBack() error {
	if err := block.state.RollBack(); err != nil {
		return err
	}
	return nil
}

// Commit saves batch updates to storage
func (block *Block) Commit() error {
	if err := block.state.Commit(); err != nil {
		return err
	}
	return nil
}

func bytesToBlock(bytes []byte) (*Block, error) {
	pb := new(corepb.Block)
	if err := proto.Unmarshal(bytes, pb); err != nil {
		logging.Debug("") // TODO
		return nil, err
	}
	block := new(Block)
	if err := block.FromProto(pb); err != nil {
		logging.Debug("") // TODO
		return nil, err
	}
	return block, nil
}
