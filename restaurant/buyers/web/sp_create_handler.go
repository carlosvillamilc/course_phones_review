package web_buyer

import (
	"course-phones-review/internal/database"
	"course-phones-review/internal/logs"
	"course-phones-review/restaurant/buyers/gateway"
	"course-phones-review/restaurant/buyers/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CreateBuyerHandler struct {
	gateway.BuyerCreateGateway
}

func NewCreateBuyerHandler(client *database.MySqlClient) *CreateBuyerHandler {

	return &CreateBuyerHandler{gateway.NewBuyerCreateGateway(client)}
}

func parseRequest(r *http.Request) *models.CreateBuyerCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateBuyerCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}

func parseResponse(r *http.Request) *models.CreateBuyerCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateBuyerCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}

func (h *CreateBuyerHandler) SaveDataHandler(w http.ResponseWriter, r *http.Request) {

	response, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "Error connecting to buyers endpoint"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Log().Error(err.Error())
	}

	var buyer_arr []models.Buyer

	json.Unmarshal([]byte(responseData), &buyer_arr)

	res, err := h.SaveBuyers(buyer_arr)

	response, err = http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products")

	logs.Log().Debug(response)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "Error connecting to products endpoint"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Log().Error(err.Error())
	}

	//responseString := string(responseData)

	res, err = h.SaveProducts(string(responseData))

	/*cmd := parseResponse(response.Body)
	logs.Log().Debug(cmd)*/

	/*w.Header().Set("Content-Type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create user"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}*/

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

/*
func (h *CreateUserHandler) AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := parseRequest(r)

	if cmd.Password == "" || cmd.Username == "" {
		logs.Log().Debug("decode password:", cmd.Password)
		logs.Log().Debug("decode username:", cmd.Username)
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "Invalid json provided"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}
	res, err := h.Authenticate(cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		if err.Error() == "sql: no rows in result set" {
			m := map[string]interface{}{"msg": "User not found"}
			_ = json.NewEncoder(w).Encode(m)

		} else {
			m := map[string]interface{}{"msg": "Error in auth user"}
			_ = json.NewEncoder(w).Encode(m)
		}
		return
	} else if cmd.Password != res.Password && cmd.Username != cmd.Password {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "Invalid Credentials"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	ts, err := token.CreateToken(res.Id)

	if err != nil {
		logs.Log().Error(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		m := map[string]interface{}{"msg": "Token Creation Error"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokens)
}

func (h *CreateUserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res := h.GetUserByID(userID)

	if res == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "user id doesn't exist"})
	}

	json.NewEncoder(w).Encode(&res)
}
*/
