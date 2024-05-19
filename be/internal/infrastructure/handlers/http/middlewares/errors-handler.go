package http_middlewares

import (
	"net/http"

	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	"github.com/gin-gonic/gin"
)

func ErrorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}
		latestError := err.Err
		switch latestError.(type) {
		case *errors.InvalidInputError:
			c.JSON(http.StatusBadRequest, latestError)
		case *errors.ForbiddenError:
			c.JSON(http.StatusForbidden, latestError)
		case *errors.NotFoundError:
			c.JSON(http.StatusNotFound, latestError)
		// case *errors.UnknownError:
		// case *errors.InternalServerError:
		// 	c.JSON(http.StatusInternalServerError, latestError)
		default:
			c.JSON(http.StatusInternalServerError, c.Errors.Errors())
		}
	}
}
