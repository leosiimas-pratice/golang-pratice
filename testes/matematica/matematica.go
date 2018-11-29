package matematica2

import (
	"fmt"
	"strconv"
)

//Media é responsavel por calcular o que já sabee//
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", media), 64)
	return mediaArredondada
}
