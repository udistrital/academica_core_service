package services

import (
	"context"

	"github.com/udistrital/academica_core_service/repositories/academica"
	"github.com/udistrital/academica_core_service/responses"
)

type EstudiantesService struct {
	repo *academica.EstudiantesRepository
}

func NewEstudiantesService() *EstudiantesService {
	return &EstudiantesService{
		repo: academica.NewEstudiantesRepository(),
	}
}

func (s *EstudiantesService) GetDatosDiploma(ctx context.Context, codigo int64) (*responses.EstudianteDatosDiploma, error) {
	return s.repo.GetDatosDiploma(ctx, codigo)
}
