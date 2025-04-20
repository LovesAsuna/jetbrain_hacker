package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ObtainTicket(context *gin.Context) {
	response, err := NewObtainTicketResponse(context)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	context.Render(http.StatusOK, &SignedResponse{response})
}

type ObtainTicketResponse struct {
	*Helper            `xml:"-"`
	Action             string `xml:"action"`
	ConfirmationStamp  string `xml:"confirmationStamp"`
	LeaseSignature     string `xml:"leaseSignature"`
	Message            string `xml:"message"`
	ProlongationPeriod int    `xml:"prolongationPeriod,omitempty"`
	ResponseCode       string `xml:"responseCode"`
	Salt               string `xml:"salt"`
	ServerLease        string `xml:"serverLease"`
	ServerUid          string `xml:"serverUid"`
	TicketID           string `xml:"ticketId"`
	TicketProperties   string `xml:"ticketProperties"`
	ValidationDeadline int    `xml:"validationDeadlinePeriod"`
	ValidationPeriod   int    `xml:"validationPeriod"`
}

func NewObtainTicketResponse(context *gin.Context) (*ObtainTicketResponse, error) {
	baseRequest := new(BaseRequest)
	if err := context.Bind(baseRequest); err != nil {
		return nil, err
	}
	helper := NewHelper(context.Value(CertPoolKey).(*CertPool))
	serverUid := helper.GetServerUid()
	serverLease := "4102415999000:" + serverUid
	return &ObtainTicketResponse{
		Helper:             helper,
		Action:             "NONE",
		ConfirmationStamp:  helper.GenerateConfirmationStamp(baseRequest.MachineId),
		LeaseSignature:     helper.GenerateLeaseSignature(serverLease),
		Message:            "",
		ProlongationPeriod: 600000,
		ResponseCode:       "OK",
		Salt:               baseRequest.Salt,
		ServerLease:        serverLease,
		ServerUid:          serverUid,
		TicketID:           "666",
		TicketProperties:   fmt.Sprintf("licensee=%s", baseRequest.UserName),
		ValidationDeadline: -1,
		ValidationPeriod:   60000000,
	}, nil
}
