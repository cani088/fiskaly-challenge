package api

import (
	"encoding/json"
	"errors"
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
		handleError(response, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	var requestBody NewDeviceRequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		handleError(response, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
		return
	}

	device, err := domain.NewDevice(requestBody.Label, requestBody.Algorithm)

	if err != nil {
		handleError(response, http.StatusBadRequest, err)
		return
	}

	err = s.repo.AddDevice(*device)

	if err != nil {
		handleError(response, http.StatusBadRequest, err)
	} else {
		WriteAPIResponse(response, http.StatusOK, device.Label)
	}

}

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	var requestBody SignTransactionRequestBody
	var signatureResponse SignTransactionResponse
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		handleError(response, http.StatusInternalServerError, errors.New(http.StatusText(http.StatusInternalServerError)))
		return
	}

	var deviceLabel string = requestBody.DeviceLabel
	device, err := s.repo.GetDeviceByLabel(deviceLabel)

	if err != nil {
		handleError(response, http.StatusBadRequest, err)
		return
	}

	signatureResponse.Signature, signatureResponse.SignedData = device.SignData(requestBody.DataToBeSigned)

	// TODO: make signatureCounter only private
	err1 := s.repo.IncreaseDeviceCounter(deviceLabel)
	if err1 != nil {
		handleError(response, http.StatusBadRequest, err1)
		return
	}
	err2 := s.repo.UpdateLastSignature(deviceLabel, signatureResponse.Signature)
	if err2 != nil {
		handleError(response, http.StatusBadRequest, err2)
		return
	}
	WriteAPIResponse(response, http.StatusOK, signatureResponse)
}

func handleError(response http.ResponseWriter, statusCode int, err error) {
	WriteErrorResponse(response, statusCode, []string{err.Error()})
}
