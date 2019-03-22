package main

import (
	"jcox15/sentry-poc/app"
	"log"
	"strconv"

	raven "github.com/getsentry/raven-go"

	"github.com/goadesign/goa"
)

type operandsInteractor interface {
	Add(int, int) (int, error)
}

type exceptionMonitor interface {
	MonitorPanic(func()) (interface{}, string)
	CaptureError(error)
	CaptureMessage(string)
}

// START OMIT
// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
	monitor    exceptionMonitor
	interactor operandsInteractor
}

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

// END OMIT

// NewOperandsController creates a operands controller.
func NewOperandsController(s *goa.Service, m exceptionMonitor, i operandsInteractor) *OperandsController {
	return &OperandsController{
		Controller: s.NewController("OperandsController"),
		monitor:    m,
		interactor: i,
	}
}
