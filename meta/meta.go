package meta

import (
	"strconv"
)

type Meta struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
	PageCount  int64 `json:"page_count"`
	TotalCount int64 `json:"total_count"`
}

// en service lo nombró como NewService, pero como aquí no hay otra función que se llame New, se puede nombrar simplemente como New
func New(page, perPage, total int64, pagLimitDef string) (*Meta, error) {

	// si perPage es menor o igual a = 0 se le asigna un valor por defecto que se obtiene de las variables de entorno, conviertiendo el strign a int con strconv.Atoi y si hay un error en la conversión se devuelve ese error
	if perPage <= 0 {
		var err error
		perPage, err = strconv.ParseInt(pagLimitDef, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	// calcular el número total de páginas (pageCount), dividiendo el total de elementos entre el número de elementos por página, y redondeando hacia arriba para obtener el número entero más próximo. Si el total es mayor o igual a 0, se asigna 0 a pageCount. Si la página solicitada es mayor que el número total de páginas, se asigna el número total de páginas a la página solicitada para evitar devolver una página vacía.
	pageCount := int64(0)
	// total se refiere al total de registros o filas (db) para mostrar
	if total >= 0 {
		pageCount = int64(total+perPage-1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}, nil

}

// a partir de qué número de fila o registro voy a traer
func (p *Meta) Offset() int64 {
	return (p.Page - 1) * p.PerPage
}

// hasta qué número de filas voy a traer por página
func (p *Meta) Limit() int64 {
	return p.PerPage
}
