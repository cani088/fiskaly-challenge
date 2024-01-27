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
	DeviceLabel    string `json:"device_label"`
	DataToBeSigned string `json:"data_to_be_signed"`
}

type SignTransactionResponse struct {
	Signature  string `json:"signature"`
	SignedData string `json:"signed_data"`
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

	device := domain.NewDevice(requestBody.Label, requestBody.Algorithm)

	err = s.repo.AddDevice(*device)

	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			err.Error(),
		})
	} else {
		WriteAPIResponse(response, http.StatusOK, device.Label)
	}

}

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	var requestBody SignTransactionRequestBody
	var signatureResponse SignTransactionResponse
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			http.StatusText(http.StatusInternalServerError),
		})
	}

	var deviceLabel string = requestBody.DeviceLabel
	device, err := s.repo.GetDeviceByLabel(deviceLabel)

	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			err.Error(),
		})
	}

	signatureResponse.Signature, signatureResponse.SignedData = device.SignData(requestBody.DataToBeSigned)

	// TODO: make signatureCounter only private
	err1 := s.repo.IncreaseDeviceCounter(deviceLabel)
	err2 := s.repo.UpdateLastSignature(deviceLabel, signatureResponse.Signature)

	if err1 != nil && err2 != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			err1.Error(), err2.Error(),
		})
	} else {
		WriteAPIResponse(response, http.StatusOK, signatureResponse)
	}
}
