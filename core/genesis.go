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

package core

import (
	"io/ioutil"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/medibloc/go-medibloc/common"
	"github.com/medibloc/go-medibloc/core/pb"
	"github.com/medibloc/go-medibloc/crypto/signature/algorithm"
	"github.com/medibloc/go-medibloc/storage"
	"github.com/medibloc/go-medibloc/util"
	"github.com/medibloc/go-medibloc/util/logging"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

var (
	// GenesisHash is hash of genesis block
	GenesisHash = genesisHash("genesisHash")
	// GenesisTimestamp is timestamp of genesis block
	GenesisTimestamp = int64(0)
	// GenesisCoinbase coinbase address of genesis block
	//GenesisCoinbase = common.HexToAddress("02fc22ea22d02fc2469f5ec8fab44bc3de42dda2bf9ebc0c0055a9eb7df579056c")
	GenesisCoinbase = common.HexToAddress("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	// GenesisHeight is height of genesis block
	GenesisHeight = uint64(1)
)

func genesisHash(quote string) []byte {
	hasher := sha3.New256()
	hasher.Write([]byte(quote))
	return hasher.Sum(nil)
}

// LoadGenesisConf loads genesis conf file
func LoadGenesisConf(filePath string) (*corepb.Genesis, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := string(buf)

	genesis := new(corepb.Genesis)
	if err := proto.UnmarshalText(content, genesis); err != nil {
		return nil, err
	}
	return genesis, nil
}

//SaveGenesisConf save genesis conf to file
func SaveGenesisConf(conf *corepb.Genesis, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := proto.MarshalText(f, conf); err != nil {
		return err
	}

	return nil
}

// NewGenesisBlock generates genesis block
func NewGenesisBlock(conf *corepb.Genesis, consensus Consensus, txMap TxFactory, sto storage.Storage) (*Block, error) {
	if conf == nil {
		return nil, ErrNilArgument
	}
	blockState, err := newStates(consensus, sto)
	if err != nil {
		return nil, err
	}
	genesisBlock := &Block{
		BlockData: &BlockData{
			BlockHeader: &BlockHeader{
				hash:       GenesisHash,
				parentHash: GenesisHash,
				chainID:    conf.Meta.ChainId,
				coinbase:   GenesisCoinbase,
				reward:     util.NewUint128FromUint(0),
				timestamp:  GenesisTimestamp,
				hashAlg:    algorithm.SHA3256,
				cryptoAlg:  algorithm.SECP256K1,
				cpuRef:     util.NewUint128FromUint(0),
				cpuUsage:   util.NewUint128FromUint(0),
				netRef:     util.NewUint128FromUint(0),
				netUsage:   util.NewUint128FromUint(0),
			},
			transactions: make([]*Transaction, 0),
			height:       GenesisHeight,
		},
		storage: sto,
		state:   blockState,
		sealed:  false,
	}
	if err := genesisBlock.Prepare(); err != nil {
		return nil, err
	}
	if err := genesisBlock.BeginBatch(); err != nil {
		return nil, err
	}

	initialMessage := "Genesis block of MediBloc"
	payload := &DefaultPayload{
		Message: initialMessage,
	}
	payloadBuf, err := payload.ToBytes()
	if err != nil {
		return nil, err
	}

	initialTx := &Transaction{
		txType:    TxTyGenesis,
		to:        GenesisCoinbase,
		value:     util.Uint128Zero(),
		timestamp: GenesisTimestamp,
		chainID:   conf.Meta.ChainId,
		payload:   payloadBuf,
		hashAlg:   algorithm.SHA3256,
		cryptoAlg: algorithm.SECP256K1,
	}

	hash, err := initialTx.CalcHash()
	if err != nil {
		return nil, err
	}
	initialTx.hash = hash

	// Genesis transactions do not consume bandwidth(only put in txState)
	if err := genesisBlock.state.txState.Put(initialTx); err != nil {
		return nil, err
	}
	genesisBlock.AppendTransaction(initialTx) // append on block header

	// Token distribution
	supply := util.NewUint128()
	for _, dist := range conf.TokenDistribution {
		addr := common.HexToAddress(dist.Address)
		acc, err := genesisBlock.state.GetAccount(addr)
		if err != nil {
			return nil, err
		}
		balance, err := util.NewUint128FromString(dist.Balance)
		if err != nil {
			if err := genesisBlock.RollBack(); err != nil {
				return nil, err
			}
			return nil, err
		}

		acc.Balance = balance
		supply, err = supply.Add(balance)
		if err != nil {
			return nil, err
		}

		if err := genesisBlock.state.PutAccount(acc); err != nil {
			return nil, err
		}

		tx := &Transaction{
			txType:    TxTyGenesis,
			to:        addr,
			value:     acc.Balance,
			timestamp: GenesisTimestamp,
			chainID:   conf.Meta.ChainId,
			hashAlg:   algorithm.SHA3256,
			cryptoAlg: algorithm.SECP256K1,
		}
		hash, err = tx.CalcHash()
		if err != nil {
			return nil, err
		}
		tx.hash = hash

		// Genesis transactions do not consume bandwidth(only put in txState)
		if err := genesisBlock.state.txState.Put(tx); err != nil {
			return nil, err
		}
		genesisBlock.AppendTransaction(tx) // append on block header

	}
	genesisBlock.state.supply = supply

	for _, pbTx := range conf.Transactions {
		tx := new(Transaction)
		if err := tx.FromProto(pbTx); err != nil {
			return nil, err
		}
		if err := genesisBlock.Execute(tx, txMap); err != nil {
			return nil, err
		}
		genesisBlock.AppendTransaction(tx) // append on block header
	}

	if err := genesisBlock.Commit(); err != nil {
		return nil, err
	}
	if err := genesisBlock.Flush(); err != nil {
		return nil, err
	}

	genesisBlock.Seal()

	return genesisBlock, nil
}

// CheckGenesisBlock checks if a block is genesis block
func CheckGenesisBlock(block *Block) bool {
	if block == nil {
		return false
	}
	return true
}

// CheckGenesisConf checks if block and genesis configuration match
func CheckGenesisConf(block *Block, genesis *corepb.Genesis) bool {
	if block.ChainID() != genesis.GetMeta().ChainId {
		logging.Console().WithFields(logrus.Fields{
			"block":   block,
			"genesis": genesis,
		}).Error("Genesis ChainID does not match.")
		return false
	}

	accounts, err := block.State().accState.accounts()
	if err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to get accounts from genesis block.")
		return false
	}

	tokenDist := genesis.GetTokenDistribution()
	if len(accounts) != len(tokenDist) {
		logging.Console().WithFields(logrus.Fields{
			"accountCount": len(accounts),
			"tokenCount":   len(tokenDist),
		}).Error("Size of token distribution accounts does not match.")
		return false
	}

	return true
}
