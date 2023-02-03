package rlm

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"test-task-rlm/pkg/rlm/models"
	"time"
)

func InitRLM(MaxLimitRate, BlockTimeIncrement int, DefaultBlockTime int64) error {
	if MaxLimitRate <= 0 || DefaultBlockTime <= 0 || BlockTimeIncrement <= 0 {
		return errors.New("invalid params")
	}
	models.DefaultParams = models.Params{
		MaxRate:     MaxLimitRate,
		BlockTime:   DefaultBlockTime,
		BTIncrement: BlockTimeIncrement}
	return nil
}

func RequestLimitMiddleware(ctx *gin.Context) {
	nilParams := models.Params{}
	if models.DefaultParams == nilParams {
		models.DefaultParams = models.Params{
			MaxRate:     5,
			BlockTime:   2,
			BTIncrement: 0,
		}
	}

	var mainLimiter = models.InitRateLimiter(1, models.DefaultParams.MaxRate)

	ipAddress := ctx.Request.RemoteAddr
	fwdAddress := ctx.ClientIP()
	// It parses IP from "X-Forwarded-For"
	if fwdAddress != "" {
		ipAddress = fwdAddress
	}
	fmt.Println(ipAddress)

	l, dur := mainLimiter.GetLimiter(ipAddress)
	if dur != 0 {
		ctx.JSON(
			429,
			models.Error{
				Status:  429,
				Message: fmt.Sprintf("Too many requests, you need to wait %s", time.Unix(dur-60*60*6, 0).Format("15:04:05")),
			})
		ctx.Abort()
		return
	}

	if !l.Allow() {
		mainLimiter.Lock(ipAddress, models.DefaultParams.BlockTime)
		ctx.JSON(429, models.Error{Status: 429, Message: "Too many requests"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
