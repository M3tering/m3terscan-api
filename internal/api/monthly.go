package api

import (
	"context"
	"log"
	"m3terscan-api/internal/blockchain"
	"m3terscan-api/internal/m3tering"
	"m3terscan-api/internal/models"
	"m3terscan-api/internal/subgraph"
	"m3terscan-api/internal/util"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

func GetMonthly(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	tokenId := big.NewInt(int64(idInt))
	year := ctx.Query("year")
	if year == "" {
		ctx.JSON(400, gin.H{"error": "year is required"})
		return
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid year format"})
		return
	}
	month := ctx.Query("month")
	if month == "" {
		ctx.JSON(400, gin.H{"error": "month is required"})
		return
	}
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid month format"})
		return
	}
	interval, err := util.GetQuarterHoursSinceYear(yearInt)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid year format"})
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
	latestNonceInt := util.Bytes6ToInt64(latestNonceBytes)
	startNonce := int64(latestNonceInt) - int64(interval)
	totalDays := util.DaysUntil(yearInt, &monthInt)
	nonceMove := totalDays * 96
	endNonce := min(latestNonceInt, startNonce+int64(nonceMove))
	if startNonce > endNonce {
		log.Println("Invalid nonce range: startNonce > endNonce")
		ctx.JSON(400, gin.H{"error": "Invalid nonce range: startNonce > endNonce"})
		return
	}
	nonces := util.NonceRange(int(startNonce), int(endNonce))
	req := graphql.NewRequest(`
    query ExampleQuery($meterNumber: Int!, $nonces: [Int!]) {
        meterDataPoints(meterNumber: $meterNumber, nonces: $nonces) {
            node {
                timestamp
                payload {
                  
                    energy
                }
            }
        }
    }
`)
	req.Var("meterNumber", idInt)
	req.Var("nonces", nonces)
	var resp models.ReqStruct
	if err := subgraph.Client.Run(ctx, req, &resp); err != nil {
		log.Printf("GraphQL error: %v", err)
		ctx.JSON(500, gin.H{"error": "Failed to fetch data"})
		return
	}
	data := util.GroupByMonth(resp, monthInt, yearInt)
	ctx.JSON(200, gin.H{"data": data})
}
