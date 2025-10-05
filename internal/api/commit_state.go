package api

import (
	"context"
	"encoding/hex"
	"m3terscan-api/internal/m3tering"
	"m3terscan-api/internal/util"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func GetCommitState(ctx *gin.Context, client *ethclient.Client) {
	hexStr := ctx.Query("hash")
	if hexStr == "" {
		ctx.JSON(400, gin.H{"error": "Hash is required"})
		return
	}
	hsh := common.HexToHash(hexStr)
	tx, isPending, err := client.TransactionByHash(context.Background(), hsh)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if isPending {
		ctx.JSON(200, gin.H{
			"status": "pending",
			"hash":   hsh.Hex(),
		})
		return
	}
	inputData := tx.Data()
	parsedABI, err := abi.JSON(strings.NewReader(m3tering.M3teringABI))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(inputData) < 4 {
		ctx.JSON(500, gin.H{"error": "invalid transaction data"})
		return
	}
	selector := inputData[:4]

	method, err := parsedABI.MethodById(selector)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	args, err := method.Inputs.Unpack(inputData[4:])
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	accountsHex := hex.EncodeToString(args[0].([]byte))

	noncesHex := hex.EncodeToString(args[1].([]byte))

	accounts, err := util.HexToChunks(accountsHex, 6)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	nonces, err := util.HexToChunks(noncesHex, 6)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	meters := util.CombineAccountsNonces(accounts, nonces)

	ctx.JSON(200, gin.H{
		"data": meters,
	})
}
