
import (
	"jcox15/sentry-poc/app"
	"strconv"
)

//START OMIT
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	var respErr error

	c.monitor.CapturePanic(func() { // HLmonitor
		sum, err := c.interactor.Add(ctx.Left, ctx.Right)
		if err != nil {
			c.monitor.CaptureError(err) // HLmonitor
			respErr = ctx.InternalServerError()
			return
		}

		c.monitor.CaptureMessage(`everything is ok`) // HLmonitor
		respErr = ctx.OK([]byte(strconv.Itoa(sum)))
		return
	})

	return respErr
}

//END OMIT
