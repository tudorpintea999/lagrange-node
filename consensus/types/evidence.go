package types

import (
	"encoding/binary"
	"math/big"

	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/lagrange"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Evidence defines an evidence.
type Evidence struct {
	Operator                    string   `json:"operator" bson:"operator"`
	BlockHash                   [32]byte `json:"block_hash" bson:"block_hash"`
	CorrectBlockHash            [32]byte `json:"correct_block_hash" bson:"correct_block_hash"`
	CurrentCommitteeRoot        [32]byte `json:"current_committee_root" bson:"current_committee_root"`
	CorrectCurrentCommitteeRoot [32]byte `json:"correct_current_committee_root" bson:"correct_current_committee_root"`
	NextCommitteeRoot           [32]byte `json:"next_committee_root" bson:"next_committee_root"`
	CorrectNextCommitteeRoot    [32]byte `json:"correct_next_committee_root" bson:"correct_next_committee_root"`
	BlockNumber                 uint64   `json:"block_number" bson:"block_number"`
	EpochNumber                 uint64   `json:"epoch_number" bson:"epoch_number"`
	BlockSignature              []byte   `json:"block_signature" bson:"block_signature"`
	CommitSignature             []byte   `json:"commit_signature" bson:"commit_signature"`
	ChainID                     uint32   `json:"chain_id" bson:"chain_id"`
	Status                      bool     `json:"status" bson:"status"`
}

// GetLagrangeServiceEvidence returns the lagrange service evidence.
func GetLagrangeServiceEvidence(e *Evidence) lagrange.LagrangeServiceEvidence {
	return lagrange.LagrangeServiceEvidence{
		Operator:                    common.HexToAddress(e.Operator),
		BlockHash:                   e.BlockHash,
		CorrectBlockHash:            e.CorrectBlockHash,
		CurrentCommitteeRoot:        e.CurrentCommitteeRoot,
		CorrectCurrentCommitteeRoot: e.CorrectCurrentCommitteeRoot,
		NextCommitteeRoot:           e.NextCommitteeRoot,
		CorrectNextCommitteeRoot:    e.CorrectNextCommitteeRoot,
		BlockNumber:                 big.NewInt(int64(e.BlockNumber)),
		EpochNumber:                 big.NewInt(int64(e.EpochNumber)),
		BlockSignature:              e.BlockSignature,
		CommitSignature:             e.CommitSignature,
		ChainID:                     e.ChainID,
	}
}

// GetCommitRequestHash returns the hash of the commit block request.
func GetCommitRequestHash(req *networktypes.CommitBlockRequest) []byte {
	var blockNumberBuf, epochNumberBuf, tvpBuf common.Hash
	blockHash := common.FromHex(req.BlsSignature.ChainHeader.BlockHash)[:]
	currentCommitteeRoot := common.FromHex(req.BlsSignature.CurrentCommittee)[:]
	nextCommitteeRoot := common.FromHex(req.BlsSignature.NextCommittee)[:]
	blockNumber := big.NewInt(int64(req.BlsSignature.ChainHeader.BlockNumber)).FillBytes(blockNumberBuf[:])
	epochNumber := big.NewInt(int64(req.EpochNumber)).FillBytes(epochNumberBuf[:])
	tvp := big.NewInt(int64(req.BlsSignature.TotalVotingPower)).FillBytes(tvpBuf[:])
	blockSignature := common.FromHex(req.BlsSignature.Signature)[:]
	chainID := make([]byte, 4)
	binary.BigEndian.PutUint32(chainID, req.BlsSignature.ChainHeader.ChainId)

	return utils.Hash(
		blockHash,
		currentCommitteeRoot,
		nextCommitteeRoot,
		blockNumber,
		epochNumber,
		tvp,
		blockSignature,
		chainID,
	)
}

// GetEvidence returns the evidence from the commit block request.
func GetEvidence(req *networktypes.CommitBlockRequest, correctBlockHash, correctCurrentCommitteeRoot, correctNextCommitteeRoot string) (*Evidence, error) {
	hash := GetCommitRequestHash(req)
	signature := common.FromHex(req.Signature)
	pubKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return nil, err
	}
	// convert the signature to the legacy format which be able to be verified in Solidity
	if signature[64] == 0 || signature[64] == 1 {
		signature[64] += 27
	}
	addr := crypto.PubkeyToAddress(*pubKey).Hex()
	return &Evidence{
		Operator:                    addr,
		BlockHash:                   common.HexToHash(req.BlsSignature.ChainHeader.BlockHash),
		CorrectBlockHash:            common.HexToHash(correctBlockHash),
		CurrentCommitteeRoot:        common.HexToHash(req.BlsSignature.CurrentCommittee),
		CorrectCurrentCommitteeRoot: common.HexToHash(correctCurrentCommitteeRoot),
		NextCommitteeRoot:           common.HexToHash(req.BlsSignature.NextCommittee),
		CorrectNextCommitteeRoot:    common.HexToHash(correctNextCommitteeRoot),
		BlockNumber:                 req.BlsSignature.ChainHeader.BlockNumber,
		EpochNumber:                 req.EpochNumber,
		BlockSignature:              common.FromHex(req.BlsSignature.Signature),
		CommitSignature:             signature,
		ChainID:                     req.BlsSignature.ChainHeader.ChainId,
	}, nil
}
