
import (
	"jcox15/sentry-poc/app"
	"log"
	"strconv"

	raven "github.com/getsentry/raven-go"
)

//START OMIT
// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	tags := map[string]string{`service-name`: `service-foo`}
	var respErr error

	raven.CapturePanic(func() {
		sum, err := c.interactor.Add(ctx.Left, ctx.Right)
		if err != nil {
			log.Println(`an error occured: `, err)
			c.monitor.CaptureError(err)
			respErr = ctx.InternalServerError()
			return
		}

		raven.CaptureMessage(`everything is ok`, tags)
		respErr = ctx.OK([]byte(strconv.Itoa(sum)))
		return

	}, tags)

	return respErr
}

//END OMIT
