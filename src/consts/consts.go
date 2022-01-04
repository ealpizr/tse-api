package consts

const (
	TSE_ID_QUERY_URL     = "https://servicioselectorales.tse.go.cr/chc/consulta_cedula.aspx"
	TSE_NAME_QUERY_URL   = "https://servicioselectorales.tse.go.cr/chc/consulta_nombres.aspx"
	HTML_LABEL_NOT_FOUND = "#lblmensaje1"

	// used for id querying
	HTML_INPUT_ID        = "#txtcedula"
	HTML_BUTTON_ID_QUERY = "#btnConsultaCedula"
	HTML_LABEL_FULLNAME  = "#lblnombrecompleto"
	HTML_LABEL_BIRTHDATE = "#lblfechaNacimiento"
	HTML_LABEL_AGE       = "#lbledad"

	// used for full name querying
	HTML_INPUT_NAME          = "#txtnombre"
	HTML_INPUT_FLASTNAME     = "#txtapellido1"
	HTML_INPUT_SLASTNAME     = "#txtapellido2"
	HTML_BUTTON_NAME_QUERY   = "#btnConsultarNombre"
	HTML_LABEL_RESULTS_FOUND = "#lblmensajes"
	HTML_LABEL_RESULT_ENTRY  = "label[for^=\"chk\"]"
)
