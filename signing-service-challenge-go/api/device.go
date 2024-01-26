package api

import (
	"encoding/json"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"net/http"
)

type CreateSignatureDeviceResponse struct {
	Id    string `json:"id"`
	Label string `json:"label"`
}

type NewDeviceRequestBody struct {
	Label     string `json:"label"`
	Algorithm string `json:"algorithm"`
}

type SignTransactionRequestBody struct {
	DeviceId       string `json:"device_id"`
	DataToBeSigned string `json:"data_to_be_signed"`
}

// CreateSignatureDevice Creates a new device for the user
func (s *Server) CreateSignatureDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	var requestBody NewDeviceRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			http.StatusText(http.StatusInternalServerError),
		})
	}

	d := domain.Device{}
	createdDevice, err := d.Create(requestBody.Label, requestBody.Algorithm)

	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			err.Error(),
		})
	} else {
		WriteAPIResponse(response, http.StatusOK, createdDevice)
	}

}

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	var requestBody SignTransactionRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			http.StatusText(http.StatusInternalServerError),
		})
	}

	var deviceId string = requestBody.DeviceId
	d := domain.Device{}

	device, err := d.GetById(deviceId)

	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			err.Error(),
		})
	} else {
		WriteAPIResponse(response, http.StatusOK, device)
	}
}
