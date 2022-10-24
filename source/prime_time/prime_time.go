package prime_time

import (
	"bufio"
	"encoding/json"
	"math"
	"net"

	"github.com/ashmeet13/protohackers/source/utils"
)

type Request struct {
	Method *string  `json:"method"`
	Number *float64 `json:"number"`
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func NewPrimeTime() *PrimeTime {
	return &PrimeTime{}
}

type PrimeTime struct{}

func (p *PrimeTime) GetConnectionNetwork() string {
	return "tcp"
}

func (p *PrimeTime) Handle(connection net.Conn) {
	logger := utils.GetLogger()
	logger.Info("Received new connection")

	defer connection.Close()

	reader := bufio.NewReader(connection)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			logger.WithError(err).Error("Error in reading bytes from connection")
			return
		}

		request, err := p.buildRequest(data)
		if err != nil {
			logger.WithError(err).Error("Error in decoding request")
			return
		}

		if !p.isRequestConforming(request) {
			data, err = p.buildResponse("malformed", false)
			if err != nil {
				logger.WithError(err).Error("Error in building response")
				return
			}

			data = append(data, byte('\n'))

			_, err = connection.Write(data)
			if err != nil {
				logger.WithError(err).Error("Error in writing malformed response")
				return
			}
			return
		}

		data, err = p.buildResponse("isPrime", p.isPrime(int(*request.Number)))
		if err != nil {
			logger.WithError(err).Error("Error in building response")
			return
		}

		data = append(data, byte('\n'))

		_, err = connection.Write(data)
		if err != nil {
			logger.WithError(err).Error("Error in writing response")
			return
		}
	}

}

func (p *PrimeTime) buildRequest(data []byte) (*Request, error) {
	var request *Request
	err := json.Unmarshal(data, &request)
	return request, err
}

func (p *PrimeTime) isRequestConforming(request *Request) bool {
	if request.Method == nil || *request.Method != "isPrime" {
		return false
	}
	if request.Number == nil {
		return false
	}
	return true
}

func (p *PrimeTime) buildResponse(method string, isPrime bool) ([]byte, error) {
	response := &Response{
		Method: method,
		Prime:  isPrime,
	}

	return json.Marshal(response)
}

func (p *PrimeTime) isPrime(number int) bool {
	if number < 2 {
		return false
	}

	sqroot := int(math.Sqrt(float64(number)))

	for i := 2; i <= sqroot; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
