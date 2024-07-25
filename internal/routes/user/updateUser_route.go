package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billzayy/Booking_Web_BE/internal/db/userdb"
	"github.com/billzayy/Booking_Web_BE/internal/handlers"
	"github.com/billzayy/Booking_Web_BE/internal/pkg"
	"github.com/billzayy/Booking_Web_BE/internal/routes/middleware"
)

// @Summary Delete user
// @Description Returns Delete Successful
// @Tags Users
// @Accept multipart/form-data
// @Produce  json
// @Param image formData file true "Image File"
// @Param value formData string true "User Value"
// @Security  Bearer
// @Success 200 {object} handlers.ResponseDataType
// @Router /api/v1/update-user [patch]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPatch {
		handlers.ResponseData(w, http.StatusMethodNotAllowed, "Invalid Method")
		return
	}

	checkAuth, err := middleware.ProtectedRoute(w, r)

	if err != nil {
		handlers.ResponseData(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if checkAuth {
		// Limit the size of the incoming request body to prevent abuse
		r.ParseMultipartForm(10 << 20) // 10 MB

		// Retrieve the file from the form data
		file, handler, err := r.FormFile("image")
		if err != nil {
			handlers.ResponseData(w, http.StatusNotFound, "Error Retrieving the File")
			return
		}
		defer file.Close()

		var result handlers.UpdateUserType

		err = json.Unmarshal([]byte(r.FormValue("value")), &result)
		if err != nil {
			handlers.ResponseData(w, http.StatusBadRequest, err.Error())
			return
		}

		err = pkg.SaveImage(file, handler.Filename, "Users")

		if err != nil {
			handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
			return
		}

		result.Image = fmt.Sprintf("/images/Users/%s", handler.Filename)

		err = userdb.UpdateUserDB(result)

		if err != nil {
			handlers.ResponseData(w, http.StatusInternalServerError, err.Error())
			return
		}

		handlers.ResponseData(w, http.StatusOK, "Updated Successful !")
	} else {
		handlers.ResponseData(w, http.StatusUnauthorized, "Invalid Token")
	}
}
