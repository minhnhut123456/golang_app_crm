package netkit

import (
	"encoding/json"
	"net/http"
)

// HTTP headers
const (
	HeaderVary             = "Vary"
	HeaderOrigin           = "Origin"
	HeaderAccept           = "Accept"
	HeaderContentType      = "Content-Type"
	HeaderAuthorization    = "Authorization"
	HeaderAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderXFrameOptions    = "X-Frame-Options"
	HeaderXRequestID       = "X-REQUEST-ID"
)

// Content types
const (
	ContentTypeJSON           = "application/json"
	ContentTypeHTML           = "text/html"
	ContentTypeText           = "text/plain"
	ContentTypeTextXML        = "text/xml"
	ContentTypeApplicationXML = "application/xml"
)

const (
	VerdictSuccess                 = "success"
	VerdictFailure                 = "failure"
)


func SendJSON (	w http.ResponseWriter, statusCode int, verdict string, message string, data interface{},) error{
	w.WriteHeader(statusCode)
	obj := map[string]interface{}{
		"verdict": verdict,
		"message": message,
		"data": data,
	}

	body, err := json.Marshal(obj)
	if(err != nil) {
		return err
	}

	_,err = w.Write(body)
	if(err != nil){
		return err
	}

	return nil
}

func SendError(w http.ResponseWriter, what error) error {
	return SendJSON(w, http.StatusInternalServerError, VerdictFailure, what.Error(), map[string]interface{}{})
}