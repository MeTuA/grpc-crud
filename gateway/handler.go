package main

import (
	"context"
	dProto "gateway/proto"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Handler struct {
	dataServiceClient dProto.DataServiceClient
	dataLoaderClient  dProto.DataLoaderClient
}

func NewHandler(clientDataService dProto.DataServiceClient, clientDataLoader dProto.DataLoaderClient) *echo.Echo {
	h := &Handler{
		dataServiceClient: clientDataService,
		dataLoaderClient:  clientDataLoader,
	}

	a := echo.New()

	a.GET("/api/data", h.get)
	a.GET("/api/data/:id", h.getOne)
	a.POST("/api/data/:id", h.update)
	a.DELETE("/api/data/:id", h.delete)

	a.GET("api/data/load", h.load)

	return a
}

func (h *Handler) load(c echo.Context) error {
	_, err := h.dataLoaderClient.LoadData(context.Background(), &dProto.EmptyParams{})
	if err != nil{
		return c.JSON(500, err.Error())
	}

	return nil
}

func (h *Handler) get(c echo.Context) error {
	data, err := h.dataServiceClient.GetAllData(context.Background(), &dProto.EmptyParams{})
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, data.GetData())
}

func (h *Handler) getOne(c echo.Context) error {
	id := c.Param("id")
	uId, _ := strconv.Atoi(id)

	data := &dProto.Data{
		Id: int64(uId),
	}

	gotData, err := h.dataServiceClient.GetData(context.Background(), data)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, gotData)
}

func (h *Handler) update(c echo.Context) error {
	data := &dProto.Data{}
	id := c.Param("id")
	uId, _ := strconv.Atoi(id)

	err := c.Bind(data)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	data.Id = int64(uId)

	resp, err := h.dataServiceClient.UpdateData(context.Background(), data)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, resp)
}

func (h *Handler) delete(c echo.Context) error {
	id := c.Param("id")
	uId, err := strconv.Atoi(id)
	if err != nil{
		return c.JSON(500, err.Error())
	}

	data := &dProto.Data{
		Id: int64(uId),
	}

	resp, err := h.dataServiceClient.DeleteData(context.Background(), data)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, resp)
}
