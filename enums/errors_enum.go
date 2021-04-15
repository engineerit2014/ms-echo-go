package enums

const (
	ErrorDBDuplicatedKeyMsg = "duplicate key"
	ErrorRequestBodyCode    = "ms_echo_go_error_request_body"
	ErrorEmailNotEmptyMsg   = "email could not be empty"
	ErrorEmailNotEmptyCode  = "ms_echo_go_error_email_not_empty"
	ErrorEmailExistsMsg     = "error email %s already exists"
	ErrorEmailExistsCode    = "ms_echo_go_error_email_exists"
	ErrorEmailNotFoundMsg   = "error email %s not found"
	ErrorEmailNotFoundCode  = "ms_echo_go_error_email_not_found"
	ErrorInsertCode         = "ms_echo_go_error_inserting_user"
	ErrorGetByEmailCode     = "ms_echo_go_error_getting_user_by_email"
	ErrorUpdateCode         = "ms_echo_go_error_updating_user"
	ErrorDeleteCode         = "ms_echo_go_error_deleting_user"
)
