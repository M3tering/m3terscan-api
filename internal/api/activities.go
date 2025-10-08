package api

import (
	"log"
	"m3terscan-api/internal/models"
	"m3terscan-api/internal/subgraph"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

// GetActivities godoc
// @Summary      Get meter activities
// @Description  Returns a paginated list of activities for a given meter
// @Tags         activities
// @Param        id          path      int     true  "Meter ID"
// @Param        limit       query     int     false "Limit number of results" default(10)
// @Param        after       query     string  false "Cursor for pagination"
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /activities/{id} [get]
func GetActivities(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	// query params: limit, after
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if limit < 1 {
		limit = 10
	}
	after := ctx.Query("after") // optional cursor

	req := graphql.NewRequest(`
		query ExampleQuery($meterNumber: Int!, $first: Int!, $after: String) {
			meterDataPoints(
				meterNumber: $meterNumber,
				first: $first,
				after: $after,
				sortBy: HEIGHT_DESC
			) {
				cursor
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
	req.Var("first", limit)
	if after != "" {
		req.Var("after", after)
	} else {
		req.Var("after", nil)
	}

	// Response struct must match GraphQL query
	var resp models.ActivityReqStruct

	if err := subgraph.Client.Run(ctx, req, &resp); err != nil {
		log.Printf("GraphQL error for meter %d: %v", idInt, err)
		ctx.JSON(500, gin.H{"error": "Failed to fetch meter data"})
		return
	}

	// Map to []ActivityResponse
	var activities []models.ActivityResponse
	var nextCursor string
	for i, item := range resp.MeterDataPoints {
		activities = append(activities, models.ActivityResponse{
			Timestamp: int64(item.Node.Timestamp),
			Energy:    item.Node.Payload.Energy,
			Signature: item.Node.Payload.Signature,
		})
		// last cursor
		if i == len(resp.MeterDataPoints)-1 {
			nextCursor = item.Cursor
		}
	}

	ctx.JSON(200, gin.H{
		"data":       activities,
		"limit":      limit,
		"nextCursor": nextCursor,
	})
}
