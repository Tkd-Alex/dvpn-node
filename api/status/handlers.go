package status

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/version"
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/dvpn-node/context"
	"github.com/sentinel-official/dvpn-node/types"
)

func HandlerGetStatus(ctx *context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		item := &ResponseGetStatus{
			Address: ctx.Address().String(),
			Bandwidth: &Bandwidth{
				Upload:   ctx.Bandwidth().Upload.Int64(),
				Download: ctx.Bandwidth().Download.Int64(),
			},
			Handshake: &Handshake{
				Enable: ctx.Config().Handshake.Enable,
				Peers:  ctx.Config().Handshake.Peers,
			},
			IntervalSetSessions:    ctx.IntervalSetSessions(),
			IntervalUpdateSessions: ctx.IntervalUpdateSessions(),
			IntervalUpdateStatus:   ctx.IntervalUpdateStatus(),
			Location: &Location{
				City:      ctx.Location().City,
				Country:   ctx.Location().Country,
				Latitude:  ctx.Location().Latitude,
				Longitude: ctx.Location().Longitude,
				Isp:       ctx.Location().Isp,
				Org:       ctx.Location().Org,
				As:        ctx.Location().As,
			},
			Systeminfo:     ctx.Systeminfo(),
			Moniker:        ctx.Moniker(),
			Operator:       ctx.Operator().String(),
			Peers:          ctx.Service().PeerCount(),
			GigabytePrices: ctx.GigabytePrices().String(),
			HourlyPrices:   ctx.HourlyPrices().String(),
			QOS: &QOS{
				MaxPeers: ctx.Config().QOS.MaxPeers,
			},
			Type:    ctx.Service().Type(),
			Version: version.Version,
		}

		c.JSON(http.StatusOK, types.NewResponseResult(item))
	}
}
