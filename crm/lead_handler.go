package crm

import (
	"net/http"
	"strconv"

	"github.com/minhnhut123456/golang_app_crm/model"
	"github.com/minhnhut123456/golang_app_crm/netkit"
	"github.com/minhnhut123456/golang_app_crm/store"
)

type LeadHandler struct {
	stores *store.Stores
}

func NewLeadHandler (stores *store.Stores) *LeadHandler{
	return &LeadHandler{
		stores: stores,
	}
}

func (l *LeadHandler) handleCreateLead(w http.ResponseWriter, r *http.Request){
	var lead model.Lead
	err := l.stores.LeadStore.Create(&lead)

	if err != nil {
		netkit.SendError(w, err)
		return
	}

	netkit.SendJSON(w, http.StatusOK, netkit.VerdictSuccess, "Create Lead Successful", lead)
}

func (l *LeadHandler) handleGetLead(w http.ResponseWriter, r *http.Request){
	vars := netkit.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil{
		netkit.SendError(w, err)
		return
	}

	lead, err := l.stores.LeadStore.FindByID(ID)

	if err != nil{
		netkit.SendError(w, err)
		return
	}

	netkit.SendJSON(w, http.StatusOK, netkit.VerdictSuccess, "Get Lead Successful", lead)
}