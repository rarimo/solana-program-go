package scripts

import (
	"fmt"
	"testing"

	"github.com/mr-tron/base58"
)

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
