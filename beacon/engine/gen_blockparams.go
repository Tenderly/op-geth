// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package engine

import (
	"encoding/json"
	"errors"

	"github.com/tenderly/op-geth/common"
	"github.com/tenderly/op-geth/common/hexutil"
	"github.com/tenderly/op-geth/core/types"
)

var _ = (*payloadAttributesMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (p PayloadAttributes) MarshalJSON() ([]byte, error) {
	type PayloadAttributes struct {
		Timestamp             hexutil.Uint64      `json:"timestamp"             gencodec:"required"`
		Random                common.Hash         `json:"prevRandao"            gencodec:"required"`
		SuggestedFeeRecipient common.Address      `json:"suggestedFeeRecipient" gencodec:"required"`
		Withdrawals           []*types.Withdrawal `json:"withdrawals,omitempty" gencodec:"optional"`
		Transactions          []hexutil.Bytes     `json:"transactions,omitempty"  gencodec:"optional"`
		NoTxPool              bool                `json:"noTxPool,omitempty" gencodec:"optional"`
		GasLimit              *hexutil.Uint64     `json:"gasLimit,omitempty" gencodec:"optional"`
	}
	var enc PayloadAttributes
	enc.Timestamp = hexutil.Uint64(p.Timestamp)
	enc.Random = p.Random
	enc.SuggestedFeeRecipient = p.SuggestedFeeRecipient
	enc.Withdrawals = p.Withdrawals
	if p.Transactions != nil {
		enc.Transactions = make([]hexutil.Bytes, len(p.Transactions))
		for k, v := range p.Transactions {
			enc.Transactions[k] = v
		}
	}
	enc.NoTxPool = p.NoTxPool
	enc.GasLimit = (*hexutil.Uint64)(p.GasLimit)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (p *PayloadAttributes) UnmarshalJSON(input []byte) error {
	type PayloadAttributes struct {
		Timestamp             *hexutil.Uint64     `json:"timestamp"             gencodec:"required"`
		Random                *common.Hash        `json:"prevRandao"            gencodec:"required"`
		SuggestedFeeRecipient *common.Address     `json:"suggestedFeeRecipient" gencodec:"required"`
		Withdrawals           []*types.Withdrawal `json:"withdrawals,omitempty" gencodec:"optional"`
		Transactions          []hexutil.Bytes     `json:"transactions,omitempty"  gencodec:"optional"`
		NoTxPool              *bool               `json:"noTxPool,omitempty" gencodec:"optional"`
		GasLimit              *hexutil.Uint64     `json:"gasLimit,omitempty" gencodec:"optional"`
	}
	var dec PayloadAttributes
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for PayloadAttributes")
	}
	p.Timestamp = uint64(*dec.Timestamp)
	if dec.Random == nil {
		return errors.New("missing required field 'prevRandao' for PayloadAttributes")
	}
	p.Random = *dec.Random
	if dec.SuggestedFeeRecipient == nil {
		return errors.New("missing required field 'suggestedFeeRecipient' for PayloadAttributes")
	}
	p.SuggestedFeeRecipient = *dec.SuggestedFeeRecipient
	if dec.Withdrawals != nil {
		p.Withdrawals = dec.Withdrawals
	}
	if dec.Transactions != nil {
		p.Transactions = make([][]byte, len(dec.Transactions))
		for k, v := range dec.Transactions {
			p.Transactions[k] = v
		}
	}
	if dec.NoTxPool != nil {
		p.NoTxPool = *dec.NoTxPool
	}
	if dec.GasLimit != nil {
		p.GasLimit = (*uint64)(dec.GasLimit)
	}
	return nil
}
