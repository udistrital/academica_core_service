package academica

import (
	"context"
	"fmt"

	"github.com/udistrital/academica_core_service/models"
	"github.com/udistrital/academica_core_service/responses"
)

type EstudiantesRepository struct{}

func NewEstudiantesRepository() *EstudiantesRepository {
	return &EstudiantesRepository{}
}

func (r *EstudiantesRepository) GetDatosDiploma(ctx context.Context, codigo int64) (*responses.EstudianteDatosDiploma, error) {
	const query = `
SELECT
	AE.EST_NOMBRE AS EST_NOMBRE,
	TO_CHAR(AE.EST_NRO_IDEN) AS EST_NRO_IDEN,
	AT.TIT_NOMBRE AS TIT_NOMBRE,
	G.TDO_ABREV AS TDO_ABREV,
	NVL(GM.MUN_NOMBRE, '') AS MUN_NOMBRE
FROM MNTAC.ACEST AE
JOIN MNTAC.ACTITULO AT ON AE.EST_CRA_COD = AT.TIT_CRA_COD
JOIN MNTGE.GETIPDOCU G ON AE.EST_TIPO_IDEN = G.TDO_CODVAR
JOIN MNTAC.ACESTOTR AEO ON AE.EST_COD = AEO.EOT_COD
LEFT JOIN MNTGE.GEMUNICIPIO GM ON AEO.EOT_COD_MUN_EXP = GM.MUN_COD
WHERE AE.EST_COD = :1
	AND AE.EST_SEXO = AT.TIT_SEXO`

	var datos responses.EstudianteDatosDiploma
	if err := models.DB.GetContext(ctx, &datos, query, codigo); err != nil {
		return nil, fmt.Errorf("consultando datos diploma estudiante: %w", err)
	}
	return &datos, nil
}
