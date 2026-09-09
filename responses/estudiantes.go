package responses

type EstudianteDatosDiploma struct {
	Nombre               string `db:"EST_NOMBRE" json:"nombre"`
	NumeroIdentificacion string `db:"EST_NRO_IDEN" json:"numero_identificacion"`
	Titulo               string `db:"TIT_NOMBRE" json:"titulo"`
	TipoDocumento        string `db:"TDO_ABREV" json:"tipo_documento"`
	MunicipioExpedicion  string `db:"MUN_NOMBRE" json:"municipio_expedicion"`
}
