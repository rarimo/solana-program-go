package scripts

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
)

func TestGenSeed(t *testing.T) {
	key, seed := GenTokenSeed("DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n")
	fmt.Println(base58.Encode(seed[:]))
	fmt.Println(base58.Encode(key[:]))
}

func TestInitBridgeAdmin(t *testing.T) {
	_, seed := GenSeed("DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n", "9bBGSSJ68LRPQ2dLGEczip7NCo2Xc6LgLrkmowqZwfLi")
	fmt.Println(base58.Encode(seed[:]))

	pub := hexutil.MustDecode("0x04928d7a512b18fcbfd51c1b050e2d498f962e2c244bb7495e253731cddcfd164ef32c213e3b4fd4185d1de33dd596061473392aa73b532906e553e543801c0f3a")

	InitBridgeAdmin(
		base58.Encode(seed[:]),
		"DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n",
		base58.Encode(pub[1:]),
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
		"9bBGSSJ68LRPQ2dLGEczip7NCo2Xc6LgLrkmowqZwfLi",
	)
}

func TestInitCommissionAdmin(t *testing.T) {
	bridgeAdmin, _ := GetBridgeAdmin(Get32ByteFromString("DiqNmFspSvCrfLgZBWB7z8zvVTQr5xTASWtkbWofyVAc"), solana.MustPublicKeyFromBase58("DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n"))

	InitCommissionAdmin(
		"9bBGSSJ68LRPQ2dLGEczip7NCo2Xc6LgLrkmowqZwfLi",
		bridgeAdmin.String(),
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestDepositNative(t *testing.T) {
	DepositNative(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"0xd30a6d9589a4ad1845f4cfb6cdafa47e2d444fcc568cef04525f1d700f4e53aa",
		"Solana",
		10000000,
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestDepositFT(t *testing.T) {
	DepositFT(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"AwgqsvhfQLorGrqKpXzTe22DtLt8be333Efz8u3dN2hm",
		"0xd30a6d9589a4ad1845f4cfb6cdafa47e2d444fcc568cef04525f1d700f4e53aa",
		"Solana",
		1,
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestDepositFTBurned(t *testing.T) {
	DepositFTBurned(
		"BhBxfhsg2CrckxRKqHJXykLVTSD5LLF8nqNyzXD9idY",
		"GexDbBi7B2UrJDi9JkrWH9fFVhmysN7u5C9zT2HkC6yZ",
		"9KqTXfenmiyi5UbxWN2GSnMhpATekuDrQgWvoZdFkJd",
		"0xd30a6d9589a4ad1845f4cfb6cdafa47e2d444fcc568cef04525f1d700f4e53aa",
		"Solana",
		100000,
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
