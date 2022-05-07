package ticket

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/handle/rbac"
	"github.com/ez-deploy/ezdeploy/models"
	"github.com/ez-deploy/ezdeploy/restapi/operations/pod"
	"github.com/go-openapi/runtime/middleware"
	"github.com/thanhpk/randstr"
	"github.com/wuhuizuo/sqlm"
)

// TicketOperationsImpl.
type TicketOperationsImpl struct {
	Tables *db.Tables
}

/* CheckPodTicket Check Pod Ticket, and delete it */
func (t *TicketOperationsImpl) CheckPodTicket(params pod.CheckPodTicketParams) middleware.Responder {
	// get ticket by ticket.
	ticket := &models.SSHPodTicket{}
	err := t.Tables.Ticket.Get(sqlm.SelectorFilter{"ticket": params.TicketValue}, ticket)
	if errors.Is(err, sql.ErrNoRows) {
		errBody := newError("ticket not found")
		return pod.NewCheckPodTicketForbidden().WithPayload(errBody)
	}

	if err != nil {
		errBody := newError("get ticket error, get err = ", err.Error())
		return pod.NewCheckPodTicketInternalServerError().WithPayload(errBody)
	}

	// delete ticket
	log.Println(t.Tables.Ticket.Delete(sqlm.SelectorFilter{"id": ticket.ID}))

	return pod.NewCheckPodTicketOK().WithPayload(ticket)
}

/* CreatePodTicket Create Visit Pod Once Ticket */
func (t *TicketOperationsImpl) CreatePodTicket(params pod.CreatePodTicketParams, principal *models.AuthInfo) middleware.Responder {
	allowed, err := newRBACManager(t.Tables).Check(principal.UserInfo.ID, params.Body.ProjectID, models.RolePermissionPermissionWrite)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return pod.NewCreatePodTicketInternalServerError().WithPayload(errBody)
	}
	if !allowed {
		errBody := newError("permission denied")
		return pod.NewCreatePodTicketForbidden().WithPayload(errBody)
	}

	// get project by id.
	proj := &models.ProjectInfo{}
	err = t.Tables.ProjectInfo.Get(sqlm.SelectorFilter{"id": params.Body.ProjectID}, proj)
	if errors.Is(err, sql.ErrNoRows) {
		errBody := newError("project not found")
		return pod.NewCreatePodTicketForbidden().WithPayload(errBody)
	}
	if err != nil {
		errBody := newError("get project failed", err.Error())
		return pod.NewCreatePodTicketInternalServerError().WithPayload(errBody)
	}

	createTicket := &models.SSHPodTicket{
		NamespaceName: proj.Name,
		PodName:       params.Body.PodName,
		Ticket:        randstr.Base64(128),
		UserID:        principal.UserInfo.ID,
		CreateAt:      time.Now().Unix(),
	}

	createTicket.ID, err = t.Tables.Ticket.Insert(createTicket)
	if err != nil {
		errBody := newError("create ticket failed", err.Error())
		return pod.NewCreatePodTicketInternalServerError().WithPayload(errBody)
	}

	return pod.NewCreatePodTicketCreated().WithPayload(createTicket)
}

func newRBACManager(tables *db.Tables) *rbac.RBACManager {
	return &rbac.RBACManager{
		Tables: tables,
	}
}

func newError(msg ...string) *models.Error {
	return &models.Error{Message: strings.Join(msg, " ")}
}
