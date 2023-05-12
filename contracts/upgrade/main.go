package upgrade

import (
	"github.com/near/borsh-go"
	"github.com/olegfomenko/solana-go"
)

type Instruction borsh.Enum

const (
	InstructionInitAdmin Instruction = iota
	InstructionTransferOwnership
	InstructionUpgrade
)

const (
	InitializeAdminUpgradeAdminIndex = iota
	InitializeAdminFeePayerIndex
)

type InitializeAdminArgs struct {
	Instruction Instruction
	PublicKey   [64]byte
	Contract    [32]byte
}

const (
	TransferOwnershipUpgradeAdminIndex = iota
)

type TransferOwnershipArgs struct {
	Instruction  Instruction
	NewPublicKey [64]byte
	Signature    [64]byte
	RecoveryId   byte
}

const (
	UpgradeAdminIndex = iota
	UpgradeProgramDataIndex
	UpgradeProgramIndex
	UpgradeBufferIndex
	UpgradeSpillIndex
)

type UpgradeArgs struct {
	Instruction Instruction
	Signature   [64]byte
	RecoveryId  byte
}

func InitializeAdminInstruction(programId, upgradeAdmin, feePayer solana.PublicKey, args InitializeAdminArgs) (solana.Instruction, error) {
	args.Instruction = InstructionInitAdmin

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 5))
	accounts.Append(solana.NewAccountMeta(upgradeAdmin, true, false))
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

func TransferOwnershipInstruction(programId, upgradeAdmin solana.PublicKey, args TransferOwnershipArgs) (solana.Instruction, error) {
	args.Instruction = InstructionTransferOwnership

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 1))
	accounts.Append(solana.NewAccountMeta(upgradeAdmin, true, false))

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

func UpgradeInstruction(programId, upgradeAdmin, program, buffer, spill solana.PublicKey, args UpgradeArgs) (solana.Instruction, error) {
	args.Instruction = InstructionUpgrade

	programData, _, err := solana.FindProgramAddress([][]byte{program[:]}, solana.BPFLoaderUpgradeableProgramID)
	if err != nil {
		return nil, err
	}

	accounts := solana.AccountMetaSlice(make([]*solana.AccountMeta, 0, 7))
	accounts.Append(solana.NewAccountMeta(upgradeAdmin, true, false))
	accounts.Append(solana.NewAccountMeta(programData, true, false))
	accounts.Append(solana.NewAccountMeta(program, true, false))
	accounts.Append(solana.NewAccountMeta(buffer, true, false))
	accounts.Append(solana.NewAccountMeta(spill, true, true))
	accounts.Append(solana.NewAccountMeta(solana.SysVarRentPubkey, false, false))
	accounts.Append(solana.NewAccountMeta(solana.SysVarClockPubkey, false, false))
	accounts.Append(solana.NewAccountMeta(solana.BPFLoaderUpgradeableProgramID, false, false))

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
