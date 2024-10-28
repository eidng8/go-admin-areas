package api

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// CopyUrl returns a deep-clone of the provided URL.
func CopyUrl(request *url.URL) url.URL {
	return url.URL{
		Scheme:      request.Scheme,
		Opaque:      request.Opaque,
		User:        CopyUserInfo(request),
		Host:        request.Host,
		Path:        request.Path,
		RawPath:     request.RawPath,
		OmitHost:    request.OmitHost,
		ForceQuery:  request.ForceQuery,
		RawQuery:    request.RawQuery,
		Fragment:    request.Fragment,
		RawFragment: request.RawFragment,
	}
}

// CopyUrlWithoutQueryFragment returns a deep-clone of the provided URL.
func CopyUrlWithoutQueryFragment(request *url.URL) url.URL {
	return url.URL{
		Scheme:      request.Scheme,
		Opaque:      request.Opaque,
		User:        CopyUserInfo(request),
		Host:        request.Host,
		Path:        request.Path,
		RawPath:     request.RawPath,
		OmitHost:    request.OmitHost,
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}
}

// CopyUserInfo returns a clone of the provided URL's user info.
func CopyUserInfo(request *url.URL) *url.Userinfo {
	user := request.User
	if user == nil {
		return nil
	}
	pass, set := user.Password()
	if set {
		return url.UserPassword(user.Username(), pass)
	}
	return url.User(user.Username())
}

// GetUrlPath returns the path of the provided URL.
func GetUrlPath(request *url.URL) string {
	return request.Path
}

// Respond200JSON responds with a JSON object.
func Respond200JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusOK, data)
}

// Respond201JSON responds with a JSON object.
func Respond201JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusCreated, data)
}

// Respond202JSON responds with a JSON object.
func Respond202JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusAccepted, data)
}

// Respond204JSON responds with a JSON object.
func Respond204JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusNoContent, data)
}

// Respond205JSON responds with a JSON object.
func Respond205JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusResetContent, data)
}

// Respond206JSON responds with a JSON object.
func Respond206JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusPartialContent, data)
}

// Respond301JSON responds with a JSON object.
func Respond301JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusMovedPermanently, data)
}

// Respond302JSON responds with a JSON object.
func Respond302JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusFound, data)
}

// Respond304JSON responds with a JSON object.
func Respond304JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusFound, data)
}

// Respond307JSON responds with a JSON object.
func Respond307JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusTemporaryRedirect, data)
}

// Respond308JSON responds with a JSON object.
func Respond308JSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusPermanentRedirect, data)
}

// RespondWithCodeJSON responds with a JSON object and a status code.
func RespondWithCodeJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

// Error400JSON responds with a 400 error message.
func Error400JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusBadRequest, err)
}

// Error401JSON responds with a 401 error message.
func Error401JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusUnauthorized, err)
}

// Error403JSON responds with a 403 error message.
func Error403JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusForbidden, err)
}

// Error404JSON responds with a 404 error message.
func Error404JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusNotFound, err)
}

// Error405JSON responds with a 404 error message.
func Error405JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusMethodNotAllowed, err)
}

// Error422JSON responds with a 422 error message.
func Error422JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusUnprocessableEntity, err)
}

// Error429JSON responds with a 422 error message.
func Error429JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusTooManyRequests, err)
}

// Error500JSON responds with a 500 error message.
func Error500JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusInternalServerError, err)
}

// Error501JSON responds with a 500 error message.
func Error501JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusNotImplemented, err)
}

// Error503JSON responds with a 500 error message.
func Error503JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusServiceUnavailable, err)
}

// ErrorWithCodeJSON responds with an error message and a status code.
func ErrorWithCodeJSON(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}
