package web

import (
	"course-phones-review/gadgets/smartphones/gateway"
	"course-phones-review/gadgets/smartphones/models"
	"course-phones-review/internal/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *CreateSmartphoneHandler) SaveSmartphoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)		
	/*res, err := h.Create(&models.CreateSmartphoneCMD{
		Name: "Samsung S9",
		Price: 800,
		CountryOrigin: "South Korea",
		OperativeSystem: "Android 10",
	})*/

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create smartphone"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

func (h *CreateSmartphoneHandler) GetSmartphoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	smartphoneID ,err := strconv.ParseInt(chi.URLParam(r,"smartphoneID"),10,64)
			
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}
	
	res := h.GetSmartphoneByID(smartphoneID)

	
	if res == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "smartphone id doesn't exist"})		
	}	
	
	json.NewEncoder(w).Encode(&res)
}

type CreateSmartphoneHandler struct {	
	gateway.SmartphoneCreateGateway
}

func NewCreateSmartphoneHandler(client *database.MySqlClient) *CreateSmartphoneHandler {	
	
	return &CreateSmartphoneHandler{gateway.NewSmartphoneCreateGateway(client)}
}

func parseRequest(r *http.Request) *models.CreateSmartphoneCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateSmartphoneCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}
