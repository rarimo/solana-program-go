package scripts

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mr-tron/base58"
)

func TestInputData(t *testing.T) {
	amount := make([]byte, 8)
	binary.BigEndian.PutUint64(amount, uint64(500000000000))
	address := common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
	arr := append(append(make([]byte, 24), amount...), address.Bytes()...)

	// [["0","0x000000000000000000000000000000000000000000000000000000746a5288005b38da6a701c568545dcfcb03fcb875f56beddc4"],["0","0x000000000000000000000000000000000000000000000000000000746a5288005b38da6a701c568545dcfcb03fcb875f56beddc4"]]

	/*fmt.Print("[")
	for i := 3; i < 52; i += 4 {
		fmt.Printf("\"%s\", ", hexutil.Encode(append([]byte{}, arr[i-3], arr[i-2], arr[i-1], arr[i])))
	}

	// ["0x00000000", "0x00000000", "0x00000000", "0x00000000", "0x00000000", "0x00000000", "0x000000e8", "0xd4a51000", "0x5b38da6a", "0x701c5685", "0x45dcfcb0", "0x3fcb875f", "0x56beddc4"]
	fmt.Print("]")*/

	fmt.Println(len(arr))
	fmt.Println(hexutil.Encode(arr))
}

func TestInitAdmin(t *testing.T) {
	_, seed := GenSeed("GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ")
	fmt.Println(base58.Encode(seed[:]))
	pub, prv := GenKey()
	fmt.Println("Pub key: " + base58.Encode(pub))
	fmt.Println("Prv key: " + base58.Encode(prv))

	InitBridgeAdmin(
		base58.Encode(seed[:]),
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		base58.Encode(pub),
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestDepositNative(t *testing.T) {
	DepositNative(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"FCpFKSEboCUGg1Qs8NFwH2suMAHYWvFUUiVWk8cKwNqf",
		"Solana",
		100000000,
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestDepositFT(t *testing.T) {
	DepositFT(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"ECVBA4RrPrmA3B4cheE6xoqzk2m2CJsvhpdkuGmY6T88",
		"FCpFKSEboCUGg1Qs8NFwH2suMAHYWvFUUiVWk8cKwNqf",
		"Solana",
		1000000000,
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestDepositNFT(t *testing.T) {
	DepositNFT(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"5bW7KPvZwtxqv3d2y51sdxZ4AoZsFiR5iHUWUzRCr4FQ",
		"FCpFKSEboCUGg1Qs8NFwH2suMAHYWvFUUiVWk8cKwNqf",
		"Solana",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestWithdrawNative(t *testing.T) {
	WithdrawNative(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"2iJ3erx52f16pScH8ZhpB73ENHwr9gPAzzVRfb7WYGadX8hdC6v1M2DR9zBNrBzo4WbWZS59njUJ2wMEzD5ZGW3m",
		"12134567890",
		"Solana",
		100000000,
		"4nE1f6GjLjVesk4GUEDuDoGwTc72xKdYfQpFrGRZwuNS",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestWithdrawFT(t *testing.T) {
	WithdrawFT(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"tLTG4TGSbsaL3brhkYQ2JMb72DmXZsat1RYs1yk5VGpAes2rGUTzjLGdAC8Ay67XFyUjgXHkzSmwLe9ac8JzpkK",
		"ECVBA4RrPrmA3B4cheE6xoqzk2m2CJsvhpdkuGmY6T88",
		"",
		"Solana",
		100000000,
		"4nE1f6GjLjVesk4GUEDuDoGwTc72xKdYfQpFrGRZwuNS",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestWithdrawNFT(t *testing.T) {
	WithdrawNFT(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"1iUyPZc97Kc85ePHqFmCx35HDBsXbjk2hJngPHEiF4JLo8NXA5H1PgdvwmehvDJLG6SQbfBAJvzX6QTe6ejU4bx",
		"1234567890",
		"5bW7KPvZwtxqv3d2y51sdxZ4AoZsFiR5iHUWUzRCr4FQ",
		"A32b3PiVnbwNB4pXyS8Wsfca8HodGtyiPhGbAcakZXjX",
		"Solana",
		"4nE1f6GjLjVesk4GUEDuDoGwTc72xKdYfQpFrGRZwuNS",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestMintFT(t *testing.T) {
	MintFT(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
		1000_000_000_000,
	)
}

func TestMintNFTCollection(t *testing.T) {
	MintNFTCollection(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestMintNFT(t *testing.T) {
	MintNFT(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"3rodirBaxTbx6LTjs5P6wdMQAkvXWM21QTJHFNAgfSDw",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}
