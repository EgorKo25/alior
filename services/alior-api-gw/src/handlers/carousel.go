package handlers

import (
	"strconv"

	"github.com/EgorKo25/alior/tree/main/services/alior-api-gw/broker"
	"github.com/gin-gonic/gin"
)

func NewCarousel(ctx *gin.Context) *Carousel {

	return &Carousel{
		ctx: ctx,
	}
}

type Carousel struct {
	ctx    *gin.Context
	broker broker.Broker

	limit int32
}

func (c *Carousel) Parse() error {
	limit, err := strconv.ParseInt(c.ctx.Query("limit"), 10, 32)
	if err != nil {
		return err
	}
	c.limit = int32(limit)
	return nil
}
func (c *Carousel) Apply() {
	/*err := c.broker.Publish()
	if err != nil {
		c.ctx.AbortWithStatusJSON(http.StatusInternalServerError, c.ctx.Error(err))
	}

	*/
}
