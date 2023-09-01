package services

import "src/EjercicioTren/dto"

type TrenService struct {
}

func NewTrenService() *TrenService {
	return &TrenService{}
}

var trenes []*dto.Tren

func (service *TrenService) ObtenerSegunPeso(pesoVagon int) *dto.Tren {
	return ObtenerSegunPeso(pesoVagon)
}

func ObtenerSegunPeso(pesoVagon int) *dto.Tren {
	trenesArray := []*dto.Tren{}

	for _, tren := range trenes {
		sumaVagon := 0
		for _, vagon := range tren.Vagones {
			sumaVagon += vagon.Peso
		}

		if len(trenesArray) == 0 && tren.Locomotora.Peso > (sumaVagon+pesoVagon) {
			trenesArray = append(trenesArray, tren)
		} else if tren.Locomotora.Peso > (sumaVagon + pesoVagon) {
			if (tren.Locomotora.Peso - (sumaVagon + pesoVagon)) > trenesArray[len(trenesArray)-1].Locomotora.Peso-(sumaVagon+pesoVagon) {
				trenesArray = []*dto.Tren{}
				trenesArray = append(trenesArray, tren)
			} else if (tren.Locomotora.Peso - (sumaVagon + pesoVagon)) == trenesArray[len(trenesArray)-1].Locomotora.Peso-(sumaVagon+pesoVagon) {
				trenesArray = append(trenesArray, tren)
			}
		}
	}

	if len(trenesArray) == 1 {
		return trenesArray[0]
	} else {
		trenMenor := trenesArray[0]
		for _, tren := range trenesArray {
			if tren.Locomotora.PrecioPorKm < trenMenor.Locomotora.PrecioPorKm {
				trenMenor = tren
			}
		}
	}

	return nil
}

func (service *TrenService) InsertarVagon(idTren int, vagon dto.Vagon) *dto.Tren {
	for i, tren := range trenes {
		if tren.Id == idTren {
			sumaVagon := 0
			for _, vagon := range tren.Vagones {
				sumaVagon += vagon.Peso
			}

			if tren.Locomotora.Peso > (sumaVagon + vagon.Peso) {
				trenes[i].Vagones = append(trenes[i].Vagones, vagon)
				return trenes[i]
			} else {
				tren := ObtenerSegunPeso(vagon.Peso)
				if tren != nil {
					tren.Vagones = append(tren.Vagones, vagon)
					return tren
				}
			}
		}
	}

	return nil
}

func (service *TrenService) CalcularCosto(idTren int, cantKm int) int {
	for _, tren := range trenes {
		if tren.Id == idTren {
			return tren.Locomotora.PrecioPorKm * cantKm
		}
	}

	return 0
}
