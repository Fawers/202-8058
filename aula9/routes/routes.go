package routes

import (
	"4cache/cached"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Routes struct {
	cd *cached.CacheD
}

func NewRoutes() *Routes {
	return &Routes{
		cd: cached.NewCacheD(),
	}
}

func (r *Routes) Cd() *cached.CacheD {
	return r.cd
}

func (ws *Routes) add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	response := ResponseMsg{}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		response.Msg = fmt.Sprintf("método precisa ser POST; é %s", r.Method)
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	jsonReq, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Msg = err.Error()
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	var payload RequestKeyVal
	err = json.Unmarshal(jsonReq, &payload)

	if err == nil && (payload.Key == "" || payload.Value == "") {
		err = fmt.Errorf("faltando campo key ou value no json")
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Msg = fmt.Sprintf("json inválido: %s", err)
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	err = ws.cd.Add(payload.Key, payload.Value)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response.Msg = err.Error()
	} else {
		w.WriteHeader(http.StatusCreated)
		response.Ok = true
	}

	jsonResp, _ := json.Marshal(&response)
	fmt.Fprintf(w, "%s\n", jsonResp)
}

func (ws *Routes) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		response := ResponseMsg{
			Msg: fmt.Sprintf("método precisa ser GET; é %s", r.Method)}
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	if !r.URL.Query().Has("key") {
		w.WriteHeader(http.StatusBadRequest)
		response := ResponseMsg{Msg: "parâmetro key não encontrado"}
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	value, err := ws.cd.Get(r.URL.Query().Get("key"))

	if err != nil {
		w.WriteHeader(404)
		response := ResponseMsg{Msg: err.Error()}
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	response := ResponseGet{
		Ok:    true,
		Value: value,
	}

	jsonResp, _ := json.Marshal(&response)
	fmt.Fprintf(w, "%s\n", jsonResp)
}

func (ws *Routes) update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	response := ResponseMsg{}

	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		response.Msg = fmt.Sprintf("método precisa ser PUT; é %s", r.Method)
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	jsonReq, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Msg = err.Error()
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	var payload RequestKeyVal
	err = json.Unmarshal(jsonReq, &payload)

	if err == nil && (payload.Key == "" || payload.Value == "") {
		err = fmt.Errorf("faltando campo key ou value no json")
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Msg = fmt.Sprintf("json inválido: %s", err)
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	err = ws.cd.Update(payload.Key, payload.Value)

	if err != nil {
		w.WriteHeader(404)
		response.Msg = err.Error()
	} else {
		response.Ok = true
	}

	jsonResp, _ := json.Marshal(&response)
	fmt.Fprintf(w, "%s\n", jsonResp)
}

func (ws *Routes) del(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		response := ResponseMsg{
			Msg: fmt.Sprintf("método precisa ser DELETE; é %s", r.Method),
		}
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	if !r.URL.Query().Has("key") {
		w.WriteHeader(http.StatusBadRequest)
		response := ResponseMsg{Msg: "parâmetro key não encontrado"}
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	key := r.URL.Query().Get("key")
	err := ws.cd.Del(key)

	if err != nil {
		w.WriteHeader(404)
		response := ResponseMsg{Msg: err.Error()}
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	response := ResponseDel{
		Ok:  true,
		Key: key,
	}

	jsonResp, _ := json.Marshal(&response)
	fmt.Fprintf(w, "%s\n", jsonResp)
}

func (ws *Routes) getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		response := ResponseMsg{
			Msg: fmt.Sprintf("método precisa ser GET; é %s", r.Method),
		}
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	response := ResponseGetAll{Ok: true, Pairs: []Pair{}}
	pairs := ws.cd.GetAll()

	for _, pair := range pairs {
		response.Pairs = append(response.Pairs, Pair{pair[0], pair[1]})
	}

	jsonResp, _ := json.MarshalIndent(&response, "", "    ")
	fmt.Fprintf(w, "%s\n", jsonResp)
}

func (ws *Routes) delAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		response := ResponseMsg{
			Msg: fmt.Sprintf("método precisa ser DELETE; é %s", r.Method),
		}
		jsonResp, _ := json.Marshal(&response)
		fmt.Fprintf(w, "%s\n", jsonResp)
		return
	}

	response := ResponseMsg{Ok: true}
	ws.cd.DelAll()

	jsonResp, _ := json.Marshal(&response)
	fmt.Fprintf(w, "%s\n", jsonResp)
}

func GetRoutes(r *Routes) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/add", r.add)
	mux.HandleFunc("/get", r.get)
	mux.HandleFunc("/update", r.update)
	mux.HandleFunc("/del", r.del)
	mux.HandleFunc("/get-all", r.getAll)
	mux.HandleFunc("/del-all", r.delAll)

	return mux
}
