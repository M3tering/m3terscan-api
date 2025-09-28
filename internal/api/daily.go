package api

import (
	"fmt"
	"log"
	"m3terscan-api/internal/models"
	"m3terscan-api/internal/subgraph"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

func GetDaily(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := graphql.NewRequest(`
		query ExampleQuery($meterNumber: Int!) {
			meterDataPoints(meterNumber: $meterNumber, first: 96, sortBy: HEIGHT_DESC) {
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
	req.Header.Set("Cache-Control", "no-cache")

	var resp models.DailyReqStruct
	if err := subgraph.Client.Run(ctx, req, &resp); err != nil {
		log.Printf("GraphQL error for meter %d: %v", idInt, err)
		ctx.JSON(500, gin.H{"error": "Failed to fetch meter data"})
		return
	}

	// Map: hour string -> energy sum
	hourly := make(map[string]float64)

	for _, edge := range resp.MeterDataPoints {
		node := edge.Node
		ts := time.Unix(int64(node.Timestamp), 0).UTC()
		hourStr := ts.Format("15:00") // "00:00", "01:00", ...

		hourly[hourStr] += node.Payload.Energy
	}

	// Build response in order
	result := make([]models.DailyResponseStruct, 0, 24)
	for h := range 24 {
		hour := fmt.Sprintf("%02d:00", h)
		result = append(result, models.DailyResponseStruct{
			Hour:   hour,
			Energy: hourly[hour],
		})
	}

	ctx.JSON(200, result)
}
