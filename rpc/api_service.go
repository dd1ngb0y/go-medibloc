// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package rpc

import (
	"encoding/hex"

	"github.com/medibloc/go-medibloc/consensus/dpos"

	"github.com/medibloc/go-medibloc/util/math"

	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/common/trie"
	"github.com/medibloc/go-medibloc/core"
	"github.com/medibloc/go-medibloc/core/pb"
	"github.com/medibloc/go-medibloc/rpc/pb"
	"github.com/medibloc/go-medibloc/util"
	"github.com/medibloc/go-medibloc/util/byteutils"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// APIService is blockchain api rpc service.
type APIService struct {
	bm *core.BlockManager
	tm *core.TransactionManager
	ee *core.EventEmitter
}

func newAPIService(bm *core.BlockManager, tm *core.TransactionManager, ee *core.EventEmitter) *APIService {
	return &APIService{
		bm: bm,
		tm: tm,
		ee: ee,
	}
}

// GetAccount handles GetAccount rpc.
func (s *APIService) GetAccount(ctx context.Context, req *rpcpb.GetAccountRequest) (*rpcpb.GetAccountResponse, error) {
	var block *core.Block
	var err error

	// XOR
	if !(math.XOR(len(req.Address) != 66, req.Alias == "") && math.XOR(req.Type == "", req.Height == 0)) {
		return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidRequest)
	}

	if req.Type != "" {
		switch req.GetType() {
		case GENESIS:
			block, err = s.bm.BlockByHeight(1)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, ErrMsgInternalError)
			}
		case CONFIRMED:
			block = s.bm.LIB()
		case TAIL:
			block = s.bm.TailBlock()
		default:
			return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidBlockType)
		}
	} else {
		block, err = s.bm.BlockByHeight(req.GetHeight())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidBlockHeight)
		}
	}

	if block == nil {
		return nil, status.Error(codes.InvalidArgument, ErrMsgInternalError)
	}

	var addr string
	if common.IsHexAddress(req.Address) {
		addr = req.Address
	} else {
		alisAcc, err := block.State().AccState().GetAliasAccount(req.Alias)
		if err == trie.ErrNotFound {
			return nil, status.Error(codes.NotFound, ErrMsgAliasNotFound)
		}
		if err != nil {
			return nil, status.Error(codes.Internal, ErrMsgInternalError)
		}
		addr = alisAcc.Account.String()
	}

	acc, err := block.State().GetAccount(common.HexToAddress(addr))
	if err != nil && err != trie.ErrNotFound {
		return nil, status.Error(codes.Internal, ErrMsgInternalError)
	}
	return coreAccount2rpcAccount(acc, block.Timestamp(), addr)
}

// GetBlock returns block
func (s *APIService) GetBlock(ctx context.Context, req *rpcpb.GetBlockRequest) (*rpcpb.GetBlockResponse, error) {
	var block *core.Block
	var err error

	if !math.TernaryXOR(len(req.Hash) == 64, req.Type != "", req.Height != 0) {
		return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidRequest)
	}

	if len(req.Hash) == 64 {
		block = s.bm.BlockByHash(byteutils.FromHex(req.Hash))
		if block == nil {
			return nil, status.Error(codes.Internal, ErrMsgBlockNotFound)
		}
		return coreBlock2rpcBlock(block, false)
	}

	switch req.Type {
	case GENESIS:
		block, err = s.bm.BlockByHeight(1)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, ErrMsgInternalError)
		}
	case CONFIRMED:
		block = s.bm.LIB()
	case TAIL:
		block = s.bm.TailBlock()
	default:
		block, err = s.bm.BlockByHeight(req.Height)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidBlockHeight)
		}
	}
	if block == nil {
		return nil, status.Error(codes.InvalidArgument, ErrMsgInternalError)
	}

	return coreBlock2rpcBlock(block, false)
}

// GetBlocks returns blocks
func (s *APIService) GetBlocks(ctx context.Context, req *rpcpb.GetBlocksRequest) (*rpcpb.GetBlocksResponse, error) {
	var rpcBlocks []*rpcpb.GetBlockResponse

	if req.From > req.To || req.From < 1 {
		return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidRequest)
	}

	if req.To-req.From > MaxBlocksCount {
		return nil, status.Error(codes.InvalidArgument, ErrMsgTooManyBlocksRequest)
	}

	if s.bm.TailBlock().Height() < req.To {
		req.To = s.bm.TailBlock().Height()
	}

	for i := req.From; i <= req.To; i++ {
		block, err := s.bm.BlockByHeight(i)
		if err != nil {
			return nil, status.Error(codes.Internal, ErrMsgBlockNotFound)
		}

		rpcBlock, err := coreBlock2rpcBlock(block, true)
		if err != nil {
			return nil, status.Error(codes.Internal, ErrMsgConvertBlockFailed)
		}
		rpcBlocks = append(rpcBlocks, rpcBlock)
	}

	return &rpcpb.GetBlocksResponse{
		Blocks: rpcBlocks,
	}, nil
}

// GetCandidates returns all candidates
func (s *APIService) GetCandidates(ctx context.Context, req *rpcpb.NonParamRequest) (*rpcpb.GetCandidatesResponse, error) {
	var rpcCandidates []*rpcpb.Candidate
	block := s.bm.TailBlock()

	cs := block.State().DposState().CandidateState()
	candidates := make([]*dpos.Candidate, 0)

	iter, err := cs.Iterator(nil)
	if err != nil {
		return nil, status.Error(codes.Internal, ErrMsgInternalError)
	}
	exist, err := iter.Next()
	if err != nil {
		return nil, status.Error(codes.Internal, ErrMsgInternalError)
	}
	for exist {
		candidate := new(dpos.Candidate)
		err := candidate.FromBytes(iter.Value())
		if err != nil {
			return nil, status.Error(codes.Internal, ErrMsgInternalError)
		}
		candidates = append(candidates, candidate)

		exist, err = iter.Next()
		if err != nil {
			return nil, status.Error(codes.Internal, ErrMsgInternalError)
		}
	}

	for _, candidate := range candidates {
		rpcCandidates = append(rpcCandidates, dposCandidate2rpcCandidate(candidate))
	}
	return &rpcpb.GetCandidatesResponse{
		Candidates: rpcCandidates,
	}, nil
}

// GetDynasty returns all dynasty accounts
func (s *APIService) GetDynasty(ctx context.Context, req *rpcpb.NonParamRequest) (*rpcpb.GetDynastyResponse, error) {
	var addresses []string

	block := s.bm.TailBlock()
	dynasty, err := block.State().GetDynasty()
	if err != nil {
		return nil, status.Error(codes.Internal, ErrMsgInternalError)
	}
	for _, addr := range dynasty {
		addresses = append(addresses, addr.Hex())
	}
	return &rpcpb.GetDynastyResponse{
		Addresses: addresses,
	}, nil
}

// GetMedState return mednet state
func (s *APIService) GetMedState(ctx context.Context, req *rpcpb.NonParamRequest) (*rpcpb.GetMedStateResponse, error) {
	tailBlock := s.bm.TailBlock()
	lib := s.bm.LIB()

	return &rpcpb.GetMedStateResponse{
		ChainId: tailBlock.ChainID(),
		Tail:    byteutils.Bytes2Hex(tailBlock.Hash()),
		Height:  tailBlock.Height(),
		LIB:     byteutils.Bytes2Hex(lib.Hash()),
	}, nil
}

// GetPendingTransactions sends all transactions in the transaction pool
func (s *APIService) GetPendingTransactions(ctx context.Context, req *rpcpb.NonParamRequest) (*rpcpb.GetTransactionsResponse, error) {
	txs := s.tm.GetAll()
	rpcTxs, err := coreTxs2rpcTxs(txs, false)
	if err != nil {
		return nil, err
	}

	return &rpcpb.GetTransactionsResponse{
		Transactions: rpcTxs,
	}, nil
}

// GetTransaction returns transaction
func (s *APIService) GetTransaction(ctx context.Context, req *rpcpb.GetTransactionRequest) (*rpcpb.GetTransactionResponse, error) {
	if len(req.Hash) != 64 {
		return nil, status.Error(codes.NotFound, ErrMsgInvalidTxHash)
	}

	tailBlock := s.bm.TailBlock()
	if tailBlock == nil {
		return nil, status.Error(codes.NotFound, ErrMsgInternalError)
	}

	txHash := byteutils.Hex2Bytes(req.Hash)
	tx, err := tailBlock.State().GetTx(txHash)
	if err != nil && err != trie.ErrNotFound {
		return nil, status.Error(codes.Internal, ErrMsgGetTransactionFailed)
	} else if err == trie.ErrNotFound { // tx is not in txsState
		// Get tx from txPool (type *Transaction)
		tx = s.tm.Get(txHash)
		if tx == nil {
			return nil, status.Error(codes.NotFound, ErrMsgTransactionNotFound)
		}
		return CoreTx2rpcTx(tx, false)
	}
	// If tx is already included in a block
	return CoreTx2rpcTx(tx, true)
}

// GetTransactionReceipt returns transaction receipt
func (s *APIService) GetTransactionReceipt(ctx context.Context, req *rpcpb.GetTransactionRequest) (*rpcpb.
	GetTransactionReceiptResponse, error) {
	if len(req.Hash) != 64 {
		return nil, status.Error(codes.NotFound, ErrMsgInvalidTxHash)
	}

	tailBlock := s.bm.TailBlock()
	if tailBlock == nil {
		return nil, status.Error(codes.NotFound, ErrMsgInternalError)
	}

	txHash := byteutils.Hex2Bytes(req.Hash)
	tx, err := tailBlock.State().GetTx(txHash)
	if err != nil && err != trie.ErrNotFound {
		return nil, status.Error(codes.Internal, ErrMsgGetTransactionFailed)
	} else if err == trie.ErrNotFound { // tx is not in txsState
		return nil, status.Error(codes.NotFound, ErrMsgTransactionNotFound)
	}
	// If tx is already included in a block
	return coreReceipt2rpcReceipt(tx)
}

// SendTransaction sends transaction
func (s *APIService) SendTransaction(ctx context.Context, req *rpcpb.SendTransactionRequest) (*rpcpb.SendTransactionResponse, error) {
	value, err := util.NewUint128FromString(req.Value)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidTxValue)
	}
	valueBytes, err := value.ToFixedSizeByteSlice()
	if err != nil {
		return nil, status.Error(codes.Internal, ErrMsgInternalError)
	}

	payloadBuf, err := hex.DecodeString(req.Payload)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, ErrMsgBuildTransactionFail)
	}

	pbTx := &corepb.Transaction{
		Hash:      byteutils.Hex2Bytes(req.Hash),
		TxType:    req.TxType,
		To:        byteutils.Hex2Bytes(req.To),
		Value:     valueBytes,
		Timestamp: req.Timestamp,
		Nonce:     req.Nonce,
		ChainId:   req.ChainId,
		Payload:   payloadBuf,
		CryptoAlg: req.CryptoAlg,
		HashAlg:   req.HashAlg,
		Sign:      byteutils.Hex2Bytes(req.Sign),
		PayerSign: byteutils.Hex2Bytes(req.PayerSign),
		Receipt:   nil,
	}

	tx := &core.Transaction{}
	if err := tx.FromProto(pbTx); err != nil {
		return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidTransaction)
	}

	if err = s.tm.PushAndRelay(tx); err != nil {
		return nil, status.Error(codes.InvalidArgument, ErrMsgInvalidTransaction)
	}
	return &rpcpb.SendTransactionResponse{
		Hash: byteutils.Bytes2Hex(tx.Hash()),
	}, nil
}

// Subscribe to listen event
func (s *APIService) Subscribe(req *rpcpb.SubscribeRequest, stream rpcpb.ApiService_SubscribeServer) error {

	eventSub, err := core.NewEventSubscriber(1024, req.Topics)
	if err != nil {
		return err
	}

	s.ee.Register(eventSub)
	defer s.ee.Deregister(eventSub)

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case event := <-eventSub.EventChan():
			err := stream.Send(&rpcpb.SubscribeResponse{
				Topic: event.Topic,
				Hash:  event.Data,
			})
			// TODO : Send timeout
			if err != nil {
				return err
			}
		}
	}
}

// HealthCheck returns success.
func (s *APIService) HealthCheck(ctx context.Context, req *rpcpb.NonParamRequest) (*rpcpb.HealthCheckResponse, error) {
	return &rpcpb.HealthCheckResponse{
		Ok: true,
	}, nil
}
