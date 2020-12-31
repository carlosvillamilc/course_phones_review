package web_user

import (
	"course-phones-review/gadgets/users/gateway"
	"course-phones-review/gadgets/users/models"
	"course-phones-review/internal/database"
	"course-phones-review/internal/logs"
	token "course-phones-review/internal/token"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CreateUserHandler struct {
	gateway.UserCreateGateway
}

func NewCreateUserHandler(client *database.MySqlClient) *CreateUserHandler {

	return &CreateUserHandler{gateway.NewUserCreateGateway(client)}
}

func parseRequest(r *http.Request) *models.CreateUserCMD {
	body := r.Body
	defer body.Close()
	var cmd models.CreateUserCMD

	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}

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
	/*saveErr := token.CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}*/
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	//c.JSON(http.StatusOK, tokens)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokens)
}

func (h *CreateUserHandler) SaveUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create user"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
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
