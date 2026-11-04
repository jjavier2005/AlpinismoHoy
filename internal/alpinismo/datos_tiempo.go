package alpinismo

import "fmt"

type DatosTiempo struct {
	temperaturaMaxima float64
	temperaturaMinima float64
	precipitacion     float64
	viento            float64
}

type MinimaMayorQueMaximaError struct {
	Minima float64
	Maxima float64
}

func (e *MinimaMayorQueMaximaError) Error() string {
	return fmt.Sprintf("la temperatura mínima es %v y no puede ser mayor que la máxima, que es %v", e.Minima, e.Maxima)
}

type PrecipitacionNegativaError struct {
	Precipitacion float64
}

func (e *PrecipitacionNegativaError) Error() string {
	return fmt.Sprintf("la precipitación es %v y no puede ser negativa", e.Precipitacion)
}

type VientoNegativoError struct {
	Viento float64
}

func (e *VientoNegativoError) Error() string {
	return fmt.Sprintf("el viento es %v y no puede ser negativo", e.Viento)
}
