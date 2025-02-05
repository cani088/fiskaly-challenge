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
	DeviceId       string `json:"device_id"`
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
	var res = map[string]string{"id": device.ID, "label": device.Label}

	if err != nil {
		handleError(response, http.StatusBadRequest, err)
	} else {
		WriteAPIResponse(response, http.StatusOK, res)
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

	var deviceId string = requestBody.DeviceId
	device, err := s.repo.GetDeviceById(deviceId)

	if err != nil {
		handleError(response, http.StatusBadRequest, err)
		return
	}

	signatureResponse.Signature, signatureResponse.SignedData = device.SignData(requestBody.DataToBeSigned)

	signatureDevice, _ := s.repo.GetDeviceById(deviceId)
	transaction := domain.NewTransaction(signatureResponse.Signature, signatureResponse.SignedData, signatureDevice)

	err = s.repo.AddTransaction(*transaction)
	//if !transaction.Verify() {
	//	handleError(response, http.StatusBadRequest, errors.New("signature could not be verified"))
	//	return
	//}

	if err != nil {
		handleError(response, http.StatusBadRequest, err)
		return
	}

	err = s.repo.IncreaseDeviceCounter(deviceId)
	if err != nil {
		handleError(response, http.StatusBadRequest, err)
		return
	}
	err = s.repo.UpdateLastSignature(deviceId, signatureResponse.Signature)
	if err != nil {
		handleError(response, http.StatusBadRequest, err)
		return
	}
	WriteAPIResponse(response, http.StatusOK, signatureResponse)
}

func (s *Server) GetAllDevices(response http.ResponseWriter, request *http.Request) {
	devices := s.repo.GetAllDevices()
	WriteAPIResponse(response, http.StatusOK, devices)
}

func (s *Server) getAllTransactions(response http.ResponseWriter, request *http.Request) {
	transactions := s.repo.GetAllTransactions()
	WriteAPIResponse(response, http.StatusOK, transactions)
}

func handleError(response http.ResponseWriter, statusCode int, err error) {
	WriteErrorResponse(response, statusCode, []string{err.Error()})
}
