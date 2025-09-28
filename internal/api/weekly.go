package api

import (
	"context"
	"log"
	"m3terscan-api/internal/blockchain"
	"m3terscan-api/internal/m3tering"
	"m3terscan-api/internal/utils"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetWeekly(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	tokenId := big.NewInt(int64(idInt))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	contractAddress := common.HexToAddress("0xf8f2d4315DB5db38f3e5c45D0bCd59959c603d9b")
	client, err := blockchain.GetClient()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	instance, err := m3tering.NewM3tering(contractAddress, client)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	latestNonceBytes, err := instance.Nonce(&bind.CallOpts{Context: context.Background()}, tokenId)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	nonceInt := utils.Bytes6ToUint64(latestNonceBytes)
	ctx.JSON(200, gin.H{"latestNonce": nonceInt})
}
