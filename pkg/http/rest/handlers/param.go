package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/angelmendozacap/go-structure/message"
	"github.com/labstack/echo/v4"

	paramDom "github.com/angelmendozacap/go-structure/pkg/param/domain"
	paramInf "github.com/angelmendozacap/go-structure/pkg/param/infraestructure"
)

// ParamHandler estructura que tiene los handler de tag
type ParamHandler struct {
	Engine string
	DB     *sql.DB
}

// NewParamHandler devuelve un puntero a Handler.
func NewParamHandler(engine string, db *sql.DB) *ParamHandler {
	return &ParamHandler{engine, db}
}

// Create handler para crear un registro de user
func (h *ParamHandler) Create(c echo.Context) error {
	mr := message.ResponseMessage{}
	m := &paramDom.Param{}

	err := c.Bind(m)
	if err != nil {
		log.Printf("warning: la estructura param no es correcta. Handler user.Create: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusBadRequest),
			"¡Upps! debes enviarnos una estructura valida",
			"revisa la documentación del paquete",
		)
		return c.JSON(http.StatusBadRequest, mr)
	}

	ms := paramInf.NewStore(h.Engine, h.DB)
	err = ms.Create(m)
	if err != nil {
		log.Printf("error: no se pudo registrar el modelo. Handler user.Create: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusInternalServerError),
			"¡Upps! no pudimos crear el registro",
			"para descubrir que sucedio revisa los log del servicio",
		)
		return c.JSON(http.StatusInternalServerError, mr)
	}

	mr.AddMessage(strconv.Itoa(http.StatusCreated), "¡listo!", "")
	mr.Data = m

	return c.JSON(http.StatusCreated, mr)
}

// Update handler para actualizar un registro de user
func (h *ParamHandler) Update(c echo.Context) error {
	mr := message.ResponseMessage{}
	m := &paramDom.Param{}

	id := c.Param("id")

	err := c.Bind(m)
	if err != nil {
		log.Printf("warning: la estructura no es correcta. Handler user.Update: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusBadRequest),
			"¡Upps! debes enviarnos una estructura valida",
			"revisa la documentación del paquete",
		)
		return c.JSON(http.StatusBadRequest, mr)
	}

	ms := paramInf.NewStore(h.Engine, h.DB)
	m.ParamID = id
	err = ms.Update(m.ParamID, m)
	if err != nil {
		log.Printf("error: error al actualizar. Handler user.Update: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusInternalServerError),
			"¡Upps! no pudimos actualizar el registro",
			"para descubrir que sucedio revisa los log del servicio",
		)
		return c.JSON(http.StatusInternalServerError, mr)
	}

	mr.AddMessage(strconv.Itoa(http.StatusOK), "¡listo!", "")
	mr.Data = m

	return c.JSON(http.StatusOK, mr)
}

// GetByID handler para obtener un registro de user
func (h *ParamHandler) GetByID(c echo.Context) error {
	mr := message.ResponseMessage{}

	id := c.Param("id")

	ms := paramInf.NewStore(h.Engine, h.DB)
	res, err := ms.GetByID(id)
	if err == sql.ErrNoRows {
		mr.AddMessage(
			strconv.Itoa(http.StatusNoContent),
			"nos dimos cuenta que no tenemos datos para este id",
			"",
		)
		return c.JSON(http.StatusOK, mr)
	}
	if err != nil {
		log.Printf("error: no se pudo obtener los datos solicitados del id: %s. Handler user.GetByID: %v", id, err)
		mr.AddError(
			strconv.Itoa(http.StatusInternalServerError),
			"¡Upps! no pudimos consultar la información",
			"para descubrir que sucedio revisa los log del servicio",
		)
		return c.JSON(http.StatusInternalServerError, mr)
	}

	mr.AddMessage(strconv.Itoa(http.StatusOK), "¡listo!", "")
	mr.Data = res

	return c.JSON(http.StatusOK, mr)
}

// ToggleActive handler toggles active field
func (h *ParamHandler) ToggleActive(c echo.Context) error {
	mr := message.ResponseMessage{}

	id := c.Param("id")

	ms := paramInf.NewStore(h.Engine, h.DB)
	res, err := ms.ToggleActive(id)
	if err == sql.ErrNoRows {
		mr.AddMessage(
			strconv.Itoa(http.StatusNoContent),
			"nos dimos cuenta que no tenemos datos para este id",
			"",
		)
		return c.JSON(http.StatusOK, mr)
	}
	if err != nil {
		log.Printf("error: no se pudo obtener los datos solicitados del id: %s. Handler param.ToggleActive: %v", id, err)
		mr.AddError(
			strconv.Itoa(http.StatusInternalServerError),
			"¡Upps! no pudimos consultar la información",
			"para descubrir que sucedio revisa los log del servicio",
		)
		return c.JSON(http.StatusInternalServerError, mr)
	}

	mr.AddMessage(strconv.Itoa(http.StatusOK), "Active field updated successfully.", "")
	mr.Data = res

	return c.JSON(http.StatusOK, mr)
}

// GetAll handler para obtener todos los registro de user
func (h *ParamHandler) GetAll(c echo.Context) error {
	mr := message.ResponseMessage{}

	ms := paramInf.NewStore(h.Engine, h.DB)
	res, err := ms.GetAll()
	if err != nil {
		log.Printf("error: no se pudo obtener la información. Handler user.GetAll: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusInternalServerError),
			"¡Upps! no pudimos consultar la información",
			"para descubrir que sucedio revisa los log del servicio",
		)
		return c.JSON(http.StatusInternalServerError, mr)
	}

	mr.AddMessage(strconv.Itoa(http.StatusOK), "¡listo!", "")
	mr.Data = res

	return c.JSON(http.StatusOK, mr)
}
