package scripts

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	"github.com/olegfomenko/solana-go"
)

func TestGetFileHash(t *testing.T) {
	path := "/Users/olegfomenko/Documents/Projects/CLion/solana-bridge-program/dist/program/bridge.so"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(hexutil.Encode(crypto.Keccak256(bytes)))
	fmt.Println(len(bytes))
}

func TestGenSeed(t *testing.T) {
	key, seed := GenTokenSeed("DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n")
	fmt.Println(base58.Encode(seed[:]))
	fmt.Println(base58.Encode(key[:]))
}

func TestHash(t *testing.T) {
	program, _ := base58.Decode("DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n")
	fmt.Println(hexutil.Encode(program))

	buffer, _ := base58.Decode("8xogATUs423Lkq5jdSgJtuSUJABmjFWhqMJtK5PTfYa9")
	fmt.Println(hexutil.Encode(buffer))

	fmt.Println(hexutil.Encode(crypto.Keccak256([]byte("Solana"), To32Bytes(AmountBytes("0")), program, buffer)))
	fmt.Println(hexutil.Encode(crypto.Keccak256([]byte("Solana"), To32Bytes(AmountBytes("0")), hexutil.MustDecode("0xb5b91203ed3bb092a7cf989079e4ba49580563e26d7b9f3f6bfeb18f62fbd10b"), hexutil.MustDecode("0x764d8761504caf391c27f66b6f11db89217c4cbfc7fe337bb5060e801418232e"))))

	fmt.Println(base58.Encode([]byte("Solana"))) // iYoY7MCt
	fmt.Println(base58.Encode(To32Bytes(AmountBytes("0"))))
}

// 0xf6f436e0e7fac1226c44d5458574722d3850a8b58bb460f7c4c733bb16f22379
// 0x0425c7479247d1ee24d0fbf0232de04b844c06a5635e40bdc9571e3a6d4052d61ea234ddf62a75e936591ea4316b000ff1660c75d31092184a7abdc79b98147e9d
func TestGenECDSA(t *testing.T) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(hexutil.Encode(key.D.Bytes()))
	fmt.Println(hexutil.Encode(elliptic.Marshal(secp256k1.S256(), key.X, key.Y)))
}

func TestUpgradeContract(t *testing.T) {
	UpgradeContract(
		"RBvYeuGmV3FDXdFmT9JSKDKnEnJJ5zKeo8cjTakYZTC",
		"DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n",
		"8xogATUs423Lkq5jdSgJtuSUJABmjFWhqMJtK5PTfYa9",
		"0x3812a141dfa69d46ff480863efb62c16cc60df599b558d11607f68788c76f94b373db0699606e7d5dcb256483e040aab8588f64778893696e995cc8e3d7f60d400",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

func TestTransferOwnership(t *testing.T) {
	TransferOwnership(
		"RBvYeuGmV3FDXdFmT9JSKDKnEnJJ5zKeo8cjTakYZTC",
		"DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n",
		"0x04928d7a512b18fcbfd51c1b050e2d498f962e2c244bb7495e253731cddcfd164ef32c213e3b4fd4185d1de33dd596061473392aa73b532906e553e543801c0f3a",
		"0xf6f436e0e7fac1226c44d5458574722d3850a8b58bb460f7c4c733bb16f22379",
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
}

// solana program set-upgrade-authority <PROGRAM_ADDRESS> --new-upgrade-authority <NEW_UPGRADE_AUTHORITY>
// solana program set-upgrade-authority DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n --new-upgrade-authority HoPpdFDuk6uTFm3MsxfR1vWfZ2Uoj59yXSAQHY6G7RFj
func TestInitUpgradeAdmin(t *testing.T) {
	fmt.Println(GetUpgradeAdmin(solana.MustPublicKeyFromBase58("DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n"), solana.MustPublicKeyFromBase58("RBvYeuGmV3FDXdFmT9JSKDKnEnJJ5zKeo8cjTakYZTC")))
	pub := hexutil.MustDecode("0x04928d7a512b18fcbfd51c1b050e2d498f962e2c244bb7495e253731cddcfd164ef32c213e3b4fd4185d1de33dd596061473392aa73b532906e553e543801c0f3a")

	InitUpgradeAdmin(
		"RBvYeuGmV3FDXdFmT9JSKDKnEnJJ5zKeo8cjTakYZTC",
		"DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n",
		base58.Encode(pub[1:]),
		"4kaCgatohjE7RtkqiPW41Q9Y6CSLZft32Z5ubG5rjWgD2qp9gAmXXQTdMLRM6FT2M7Hc6SeCifd3ShkMw1uwyLnm",
	)
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
		"DiqNmFspSvCrfLgZBWB7z8zvVTQr5xTASWtkbWofyVAc",
		"DENUHfQksvSuy6eB9J8oCrLpBiU3uzar1ruTJYUehZ4n",
		"9bBGSSJ68LRPQ2dLGEczip7NCo2Xc6LgLrkmowqZwfLi",
		"0xd30a6d9589a4ad1845f4cfb6cdafa47e2d444fcc568cef04525f1d700f4e53aa",
		"Solana",
		1000,
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
