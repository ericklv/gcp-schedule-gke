package main

import (
	"log"
	"net/http"
	"scheduler/gcp"
	"scheduler/utils"

	"github.com/labstack/echo/v4"
)

const port = ":5432"
const gcp_res = "Action applied"
const gcp_err = "Action cannot be applied"

func execCmd(c echo.Context) error {

	req := new(utils.Request)
	if err := c.Bind(req); err != nil {
		log.Println(("param err"))
		log.Println(err)
		return err
	}

	body := req.KBody
	params := req.Params

	cmd_l, err := gcp.Action(params, body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.S4xx(err.Error()))
	}

	_, err_ := gcp.CallGCP(cmd_l)

	if err_ != nil {
		return c.JSON(http.StatusBadRequest, utils.S4xx(gcp_err))
	}
	return c.JSON(http.StatusOK, utils.S200(gcp_res))

}

func health(c echo.Context) error {
	log.Println("Everything is fine ...")
	return c.JSON(http.StatusOK, utils.S200("up"))
}

func main() {
	e := echo.New()

	e.POST("/k8s/:action", execCmd)
	e.GET("/health", health)

	err := e.Start(port)

	if err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
