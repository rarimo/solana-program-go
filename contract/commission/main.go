package commission

import (
	"github.com/near/borsh-go"
	"github.com/olegfomenko/solana-go"
	"gitlab.com/rarimo/solana-program-go/contract/bridge"
)

type Instruction borsh.Enum

const (
	InstructionInitAdmin Instruction = iota
	InstructionChargeCommission
	InstructionAddFeeToken
	InstructionRemoveFeeToken
	InstructionUpdateFeeToken
	InstructionWithdraw
)

type CommissionTokenType borsh.Enum

const (
	CommissionTokenTypeNative CommissionTokenType = iota
	CommissionTokenTypeFT     CommissionTokenType = iota
	CommissionTokenTypNFT     CommissionTokenType = iota
)

type CommissionToken struct {
	Type      CommissionTokenType `borsh_enum:"true"`
	PublicKey []byte
}

type CommissionTokenArg struct {
	Token  CommissionToken
	Amount uint64
}

const (
	InitializeAdminCommissionAdminIndex = iota
	InitializeAdminBridgeAdminIndex
	InitializeAdminFeePayerIndex
)

type InitializeAdminArgs struct {
	Instruction      Instruction
	AcceptableTokens []CommissionTokenArg
}

const (
	ChargeCommissionCommissionAdminIndex = iota
	ChargeCommissionOwnerIndex
	ChargeCommissionSystemProgramIndex
	ChargeCommissionRentIndex
	ChargeCommissionSPLIndex
	ChargeCommissionOwnerAssociatedIndex
	ChargeCommissionAdminAssociatedIndex
	ChargeCommissionMintIndex
)

type ChargeCommissionArgs struct {
	Instruction   Instruction
	Token         CommissionToken
	Deposit       bridge.TokenType
	DepositAmount uint64
}

const (
	FeeTokenCommissionAdminIndex = iota
	FeeTokenBridgeAdminIndex
	FeeTokenFeePayerIndex
	FeeTokenManagementIndex
)

type FeeTokenArgs struct {
	Instruction Instruction
	Origin      [32]byte
	Signature   [64]byte
	RecoveryId  byte
	Path        [][32]byte
	Token       CommissionTokenArg
}

const (
	WithdrawCommissionAdminIndex = iota
	WithdrawBridgeAdminIndex
	WithdrawReceiverIndex
	WithdrawManagementIndex
	WithdrawSystemProgramIndex
	WithdrawRentIndex
	WithdrawSPLIndex
	WithdrawReceiverAssociatedIndex
	WithdrawAdminAssociatedIndex
	WithdrawMintIndex
)

type WithdrawArgs struct {
	Instruction Instruction
	Origin      [32]byte
	Signature   [64]byte
	RecoveryId  byte
	Path        [][32]byte
	Token       CommissionTokenArg
	Amount      uint64
}

func InitializeAdminInstruction(programId, commissionAdmin, bridgeAdmin, feePayer solana.PublicKey, args InitializeAdminArgs) (solana.Instruction, error) {
	args.Instruction = InstructionInitAdmin

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 5))
	accounts.Append(solana.NewAccountMeta(commissionAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(bridgeAdmin, false, false))
	accounts.Append(solana.NewAccountMeta(feePayer, true, true))
	accounts.Append(solana.NewAccountMeta(solana.SystemProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))

	data, err := borsh.Serialize(args)
	if err != nil {
		return nil, err
	}

	return solana.NewInstruction(
		programId,
		accounts,
		data,
	), nil
}

func ChargeCommissionFTInstruction(programId, commissionAdmin, owner, mint solana.PublicKey, args ChargeCommissionArgs) (solana.Instruction, error) {
	args.Instruction = InstructionChargeCommission

	adminAssoc, _, err := solana.FindAssociatedTokenAddress(commissionAdmin, mint)
	if err != nil {
		return nil, err
	}

	ownerAssoc, _, err := solana.FindAssociatedTokenAddress(owner, mint)
	if err != nil {
		return nil, err
	}

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 4))
	accounts.Append(solana.NewAccountMeta(commissionAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(owner, true, true))
	accounts.Append(solana.NewAccountMeta(solana.SystemProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))
	accounts.Append(solana.NewAccountMeta(solana.TokenProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(ownerAssoc, true, false))
	accounts.Append(solana.NewAccountMeta(adminAssoc, true, false))
	accounts.Append(solana.NewAccountMeta(mint, false, false))

	data, err := borsh.Serialize(args)
	if err != nil {
		return nil, err
	}

	return solana.NewInstruction(
		programId,
		accounts,
		data,
	), nil
}

func ChargeCommissionNativeInstruction(programId, commissionAdmin, owner solana.PublicKey, args ChargeCommissionArgs) (solana.Instruction, error) {
	args.Instruction = InstructionChargeCommission

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 4))
	accounts.Append(solana.NewAccountMeta(commissionAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(owner, true, true))
	accounts.Append(solana.NewAccountMeta(solana.SystemProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))

	data, err := borsh.Serialize(args)
	if err != nil {
		return nil, err
	}

	return solana.NewInstruction(
		programId,
		accounts,
		data,
	), nil
}

func AddFeeTokenInstruction(programId, commissionAdmin, bridgeAdmin, payer, management solana.PublicKey, args FeeTokenArgs) (solana.Instruction, error) {
	args.Instruction = InstructionAddFeeToken

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 4))
	accounts.Append(solana.NewAccountMeta(commissionAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(bridgeAdmin, false, false))
	accounts.Append(solana.NewAccountMeta(management, true, false))
	accounts.Append(solana.NewAccountMeta(solana.SystemProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))

	data, err := borsh.Serialize(args)
	if err != nil {
		return nil, err
	}

	return solana.NewInstruction(
		programId,
		accounts,
		data,
	), nil
}

func RemoveFeeTokenInstruction(programId, commissionAdmin, bridgeAdmin, payer, management solana.PublicKey, args FeeTokenArgs) (solana.Instruction, error) {
	args.Instruction = InstructionRemoveFeeToken

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 4))
	accounts.Append(solana.NewAccountMeta(commissionAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(bridgeAdmin, false, false))
	accounts.Append(solana.NewAccountMeta(management, true, false))
	accounts.Append(solana.NewAccountMeta(solana.SystemProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))

	data, err := borsh.Serialize(args)
	if err != nil {
		return nil, err
	}

	return solana.NewInstruction(
		programId,
		accounts,
		data,
	), nil
}

func UpdateFeeTokenInstruction(programId, commissionAdmin, bridgeAdmin, payer, management solana.PublicKey, args FeeTokenArgs) (solana.Instruction, error) {
	args.Instruction = InstructionUpdateFeeToken

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 4))
	accounts.Append(solana.NewAccountMeta(commissionAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(bridgeAdmin, false, false))
	accounts.Append(solana.NewAccountMeta(management, true, false))
	accounts.Append(solana.NewAccountMeta(solana.SystemProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))

	data, err := borsh.Serialize(args)
	if err != nil {
		return nil, err
	}

	return solana.NewInstruction(
		programId,
		accounts,
		data,
	), nil
}

func WithdrawNativeInstruction(programId, commissionAdmin, bridgeAdmin, receiver, management solana.PublicKey, args WithdrawArgs) (solana.Instruction, error) {
	args.Instruction = InstructionWithdraw

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 4))
	accounts.Append(solana.NewAccountMeta(commissionAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(bridgeAdmin, false, false))
	accounts.Append(solana.NewAccountMeta(receiver, true, true))
	accounts.Append(solana.NewAccountMeta(management, true, false))
	accounts.Append(solana.NewAccountMeta(solana.SystemProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))

	data, err := borsh.Serialize(args)
	if err != nil {
		return nil, err
	}

	return solana.NewInstruction(
		programId,
		accounts,
		data,
	), nil
}

func WithdrawFTInstruction(programId, commissionAdmin, bridgeAdmin, receiver, management, mint solana.PublicKey, args WithdrawArgs) (solana.Instruction, error) {
	args.Instruction = InstructionWithdraw

	adminAssoc, _, err := solana.FindAssociatedTokenAddress(commissionAdmin, mint)
	if err != nil {
		return nil, err
	}

	receiverAssoc, _, err := solana.FindAssociatedTokenAddress(receiver, mint)
	if err != nil {
		return nil, err
	}

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 4))
	accounts.Append(solana.NewAccountMeta(commissionAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(bridgeAdmin, false, false))
	accounts.Append(solana.NewAccountMeta(receiver, true, true))
	accounts.Append(solana.NewAccountMeta(management, true, false))
	accounts.Append(solana.NewAccountMeta(solana.SystemProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))
	accounts.Append(solana.NewAccountMeta(solana.TokenProgramID, false, false))
	accounts.Append(solana.NewAccountMeta(receiverAssoc, true, false))
	accounts.Append(solana.NewAccountMeta(adminAssoc, true, false))
	accounts.Append(solana.NewAccountMeta(mint, false, false))

	data, err := borsh.Serialize(args)
	if err != nil {
		return nil, err
	}

	return solana.NewInstruction(
		programId,
		accounts,
		data,
	), nil
}
