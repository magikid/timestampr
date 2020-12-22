package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"git.sr.ht/~magikid/timestamper/app/responses"
	"github.com/revel/revel"
)

// App is the main controller struct
type App struct {
	*revel.Controller
}

// Index shows the JS app
func (c App) Index() revel.Result {
	return c.Render()
}

// ConvertTimeStamp takes a timestamp and turns it into a date
func (c App) ConvertTimeStamp(userTs string) revel.Result {
	c.Validation.Required(userTs).Message("ts is a required parameter")
	c.Validation.Match(userTs, regexp.MustCompile("^[[:digit:]]+$"))
	c.Log.Infof("ts: %v", userTs)

	ts, err := strconv.Atoi(userTs)
	if err != nil || c.Validation.HasErrors() {
		return renderJSONError("Timestamp must be a number", &c)
	}
	location, _ := time.LoadLocation("UTC")

	t := time.Unix(int64(ts), 0).In(location)

	return renderNewDate(t, &c)
}

// ConvertDate takes a date and turns it into a timestamp
func (c App) ConvertDate(userDate string) revel.Result {
	c.Validation.Required(userDate).Message("date is a required parameter")
	c.Validation.Match(userDate, regexp.MustCompile("^([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\\.[0-9]+)?(([Zz]?)|([\\+|\\-]([01][0-9]|2[0-3]):[0-5][0-9]))$"))

	c.Log.Infof("Date: %v", userDate)
	date, err := time.Parse("2006-01-02T15:04:05Z", userDate)
	if err != nil || c.Validation.HasErrors() {
		return renderJSONError("Date must be in 2006-01-02T15:04:05Z format", &c)
	}

	return renderNewTimestamp(date.Unix(), &c)
}

func renderJSONError(message string, c *App) revel.Result {
	errorStruct := responses.JsonError{Message: message}
	c.Response.Status = http.StatusBadRequest
	return c.RenderJSON(errorStruct)
}

func renderNewTimestamp(ts int64, c *App) revel.Result {
	response := responses.TimestampResponse{Timestamp: ts}
	return c.RenderJSON(response)
}

func renderNewDate(date time.Time, c *App) revel.Result {
	response := responses.DateResponse{Date: date.Format("2006-01-02T15:04:05Z")}
	return c.RenderJSON(response)
}
