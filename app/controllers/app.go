package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/revel/revel"
)

// App is the main controller struct
type App struct {
	*revel.Controller
}

// JSONResponse Holds a simple JSON response
type JSONResponse struct {
	Message string `json:"message"`
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

	return renderJSONSuccess(t.Format("2006-01-02T15:04:05Z"), &c)
}

// ConvertDate takes a date and turns it into a timestamp
func (c App) ConvertDate(userDate string) revel.Result {
	c.Validation.Required(userDate).Message("date is a required parameter")
	c.Validation.Match(userDate, regexp.MustCompile("^([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\\.[0-9]+)?(([Zz])|([\\+|\\-]([01][0-9]|2[0-3]):[0-5][0-9]))$"))

	c.Log.Infof("Date: %v", userDate)
	date, err := time.Parse("2006-01-02T15:04:05Z", userDate)
	if err != nil || c.Validation.HasErrors() {
		return renderJSONError("Date must be in 2006-01-02T15:04:05Z format", &c)
	}

	return renderJSONSuccess(fmt.Sprint(date.Unix()), &c)
}

func renderJSONError(message string, c *App) revel.Result {
	errorStruct := JSONResponse{Message: message}
	c.Response.Status = http.StatusBadRequest
	return c.RenderJSON(errorStruct)
}

func renderJSONSuccess(message string, c *App) revel.Result {
	response := JSONResponse{Message: message}
	return c.RenderJSON(response)
}
