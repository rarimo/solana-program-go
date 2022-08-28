package scripts

import (
	"crypto/sha256"
	"fmt"

	"github.com/cbergoon/merkletree"
	"gitlab.com/rarify-protocol/solana-program-go/contract"
)

var _ merkletree.Content = HashContent{}

// HashContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type HashContent struct {
	// Hash of the deposit tx
	TxHash string
	// Collection address on current chain
	CurrentAddress string
	// TokenId on current chain
	CurrentId string
	// Collection address on target chain
	TargetAddress string
	// TokenId on target chain
	TargetId string
	// Receiver address on target network
	Receiver       string
	CurrentNetwork string
	TargetNetwork  string
	Amount         string
	Type           contract.TokenType
}

func (c HashContent) String() string {
	return c.TxHash + c.CurrentAddress + c.CurrentId +
		c.TargetAddress + c.TargetId + c.Receiver + c.CurrentNetwork + c.TargetNetwork + fmt.Sprint(c.Amount) + fmt.Sprint(c.Type)
}

func (c HashContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(c.String())); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (c HashContent) Equals(other merkletree.Content) (bool, error) {
	if oc, ok := other.(HashContent); ok {
		return c.TxHash == oc.TxHash &&
			c.CurrentAddress == oc.CurrentAddress &&
			c.TargetAddress == oc.TargetAddress &&
			c.CurrentId == oc.CurrentId &&
			c.TargetId == oc.TargetId &&
			c.Receiver == oc.Receiver &&
			c.CurrentNetwork == oc.CurrentNetwork &&
			c.TargetNetwork == oc.TargetNetwork &&
			c.Amount == oc.Amount &&
			c.Type == oc.Type, nil
	}
	return false, nil
}
