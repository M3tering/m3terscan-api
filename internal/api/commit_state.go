package api

import (
	"fmt"
	"m3terscan-api/internal/m3tering"
	"m3terscan-api/internal/util"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var parsedABI abi.ABI

func init() {
	var err error
	parsedABI, err = abi.JSON(strings.NewReader(m3tering.M3teringABI))
	if err != nil {
		panic(fmt.Errorf("failed to parse ABI: %w", err))
	}
}

func GetCommitState(ctx *gin.Context, client *ethclient.Client) {
	hexStr := ctx.Query("hash")
	if hexStr == "" {
		ctx.JSON(400, gin.H{"error": "Hash is required"})
		return
	}
	hsh := common.HexToHash(hexStr)

	// Use ctx.Request.Context() for cancellation support
	tx, isPending, err := client.TransactionByHash(ctx.Request.Context(), hsh)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if isPending {
		ctx.JSON(200, gin.H{"status": "pending", "hash": hsh.Hex()})
		return
	}

	inputData := tx.Data()
	if len(inputData) < 4 {
		ctx.JSON(500, gin.H{"error": "invalid transaction data"})
		return
	}

	method, err := parsedABI.MethodById(inputData[:4])
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	args, err := method.Inputs.Unpack(inputData[4:])
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// args[0] and args[1] are []byte already
	accounts, err := util.BytesToChunks(args[0].([]byte), 6)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	nonces, err := util.BytesToChunks(args[1].([]byte), 6)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	meters, err := util.CombineAccountsNonces(accounts, nonces)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": meters})
}
