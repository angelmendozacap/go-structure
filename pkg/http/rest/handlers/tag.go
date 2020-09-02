package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/angelmendozacap/go-structure/pkg/tag/infraestructure"

	"github.com/labstack/echo/v4"

	"github.com/angelmendozacap/go-structure/message"
	"github.com/angelmendozacap/go-structure/pkg/tag/domain"
)

// TagHandler estructura que tiene los handler de tag
type TagHandler struct {
	Engine string
	DB     *sql.DB
}

// NewTagHandler devuelve un puntero a Handler.
func NewTagHandler(engine string, db *sql.DB) *TagHandler {
	return &TagHandler{engine, db}
}

// Create handler para crear un registro de user
func (h *TagHandler) Create(c echo.Context) error {
	mr := message.ResponseMessage{}
	m := &domain.Tag{}

	err := c.Bind(m)
	if err != nil {
		log.Printf("warning: la estructura user no es correcta. Handler user.Create: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusBadRequest),
			"¡Upps! debes enviarnos una estructura valida",
			"revisa la documentación del paquete",
		)
		return c.JSON(http.StatusBadRequest, mr)
	}

	ms := infraestructure.NewStore(h.Engine, h.DB)
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
func (h *TagHandler) Update(c echo.Context) error {
	mr := message.ResponseMessage{}
	m := &domain.Tag{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("warning: el id debe ser numérico. Handler user.Update: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusBadRequest),
			"¡Upps! el id que nos enviaste no es un número entero",
			"",
		)
		return c.JSON(http.StatusBadRequest, mr)
	}

	err = c.Bind(m)
	if err != nil {
		log.Printf("warning: la estructura no es correcta. Handler user.Update: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusBadRequest),
			"¡Upps! debes enviarnos una estructura valida",
			"revisa la documentación del paquete",
		)
		return c.JSON(http.StatusBadRequest, mr)
	}

	ms := infraestructure.NewStore(h.Engine, h.DB)
	m.ID = uint(id)
	err = ms.Update(m.ID, m)
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

// Delete handler para eliminar un registro de user
func (h *TagHandler) Delete(c echo.Context) error {
	mr := message.ResponseMessage{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("warning: el id debe ser numérico. Handler user.Delete: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusBadRequest),
			"¡Upps! el id que nos enviaste no es un número entero",
			"",
		)
		return c.JSON(http.StatusBadRequest, mr)
	}

	ms := infraestructure.NewStore(h.Engine, h.DB)
	err = ms.Delete(uint(id))
	if err != nil {
		log.Printf("error: error al borrar el id: %d. Handler user.Delete: %v", id, err)
		mr.AddError(
			strconv.Itoa(http.StatusInternalServerError),
			"¡Upps! no pudimos eliminar el registro",
			"para descubrir que sucedio revisa los log del servicio",
		)
		return c.JSON(http.StatusInternalServerError, mr)
	}

	mr.AddMessage(strconv.Itoa(http.StatusOK), "¡listo!", "")

	return c.JSON(http.StatusOK, mr)
}

// GetByID handler para obtener un registro de user
func (h *TagHandler) GetByID(c echo.Context) error {
	mr := message.ResponseMessage{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("warning: el id debe ser numérico. Handler user.GetByID: %v", err)
		mr.AddError(
			strconv.Itoa(http.StatusBadRequest),
			"¡Upps! el id que nos enviaste no es un número entero",
			"",
		)
		return c.JSON(http.StatusBadRequest, mr)
	}

	ms := infraestructure.NewStore(h.Engine, h.DB)
	res, err := ms.GetByID(uint(id))
	if err == sql.ErrNoRows {
		mr.AddMessage(
			strconv.Itoa(http.StatusNoContent),
			"nos dimos cuenta que no tenemos datos para este id",
			"",
		)
		return c.JSON(http.StatusOK, mr)
	}
	if err != nil {
		log.Printf("error: no se pudo obtener los datos solicitados del id: %d. Handler user.GetByID: %v", id, err)
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

// GetAll handler para obtener todos los registro de user
func (h *TagHandler) GetAll(c echo.Context) error {
	mr := message.ResponseMessage{}

	ms := infraestructure.NewStore(h.Engine, h.DB)
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
