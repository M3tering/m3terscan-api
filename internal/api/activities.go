package api

import (
	"log"
	"m3terscan-api/internal/models"
	"m3terscan-api/internal/subgraph"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

func GetActivities(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := graphql.NewRequest(`
		query ExampleQuery($meterNumber: Int!) {
			meterDataPoints(meterNumber: $meterNumber, first: 10, sortBy: HEIGHT_DESC) {
				node {
					timestamp
					payload {
						energy
						signature
					}
				}
			}
		}
	`)
	req.Var("meterNumber", idInt)

	var resp models.ActivityReqStruct
	if err := subgraph.Client.Run(ctx, req, &resp); err != nil {
		log.Printf("GraphQL error for meter %d: %v", idInt, err)
		ctx.JSON(500, gin.H{"error": "Failed to fetch meter data"})
		return
	}

	// Map to []ActivityResponse
	var activities []models.ActivityResponse
	for _, item := range resp.MeterDataPoints {
		activities = append(activities, models.ActivityResponse{
			Timestamp: item.Node.Timestamp,
			Energy:    item.Node.Payload.Energy,
			Signature: item.Node.Payload.Signature,
		})
	}

	ctx.JSON(200, gin.H{"data": activities})
}
