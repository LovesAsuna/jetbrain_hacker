package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReleaseTicket(context *gin.Context) {
	response, err := NewPingResponse(context)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	context.Render(http.StatusOK, &SignedResponse{response})
}

type ReleaseTicketResponse struct {
	*Helper            `xml:"-"`
	Action             string `xml:"action"`
	ConfirmationStamp  string `xml:"confirmationStamp"`
	LeaseSignature     string `xml:"leaseSignature"`
	Message            string `xml:"message"`
	ResponseCode       string `xml:"responseCode"`
	Salt               string `xml:"salt"`
	ServerLease        string `xml:"serverLease"`
	ServerUid          string `xml:"serverUid"`
	ValidationDeadline int    `xml:"validationDeadlinePeriod"`
	ValidationPeriod   int    `xml:"validationPeriod"`
}

func NewReleaseTicketResponse(context *gin.Context) (*ReleaseTicketResponse, error) {
	baseRequest := new(BaseRequest)
	if err := context.Bind(baseRequest); err != nil {
		return nil, err
	}
	helper := NewHelper(context.Value(CertPoolKey).(*CertPool))
	serverUid := helper.GetServerUid()
	serverLease := "4102415999000:" + serverUid

	return &ReleaseTicketResponse{
		Helper:             nil,
		Action:             "NONE",
		ConfirmationStamp:  helper.GenerateConfirmationStamp(baseRequest.MachineId),
		LeaseSignature:     helper.GenerateLeaseSignature(serverLease),
		Message:            "",
		ResponseCode:       "OK",
		Salt:               baseRequest.Salt,
		ServerLease:        serverLease,
		ServerUid:          serverUid,
		ValidationDeadline: -1,
		ValidationPeriod:   600000,
	}, nil
}
