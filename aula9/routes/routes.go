package routes

import (
	"4cache/cached"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Routes struct {
	cd *cached.CacheD
}

func NewRoutes() *Routes {
	return &Routes{
		cd: cached.NewCacheD(),
	}
}

func (ws *Routes) checkExistsAll(q url.Values, queryParam ...string) (map[string]string, bool) {
	m := make(map[string]string)

	for _, param := range queryParam {
		if !q.Has(param) {
			return nil, false
		}
		m[param] = q.Get(param)
	}

	return m, true
}

func (ws *Routes) index(w http.ResponseWriter, r *http.Request) {
	links := []string{
		"<a href=add>add</a>",
		"<a href=get>get</a>",
		"<a href=update>update</a>",
		"<a href=del>del</a>",
		"<a href=add-all>add all</a>",
		"<a href=del-all>del all</a>",
	}
	content := strings.Join(links, "<br>")
	fmt.Fprintln(w, content)
}

func (ws *Routes) add(w http.ResponseWriter, r *http.Request) {
	m, ok := ws.checkExistsAll(r.URL.Query(), "key", "value")
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "400 Bad Request - faltando parâmetro 'key' ou 'value'")
		return
	}

	ws.cd.Add(m["key"], m["value"])
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, "ok")
}

func (ws *Routes) get(w http.ResponseWriter, r *http.Request) {
	m, ok := ws.checkExistsAll(r.URL.Query(), "key")
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "400 Bad Request - faltando parâmetro 'key'")
		return
	}

	value, err := ws.cd.Get(m["key"])

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, "chave não encontrada")
	} else {
		fmt.Fprintln(w, value)
	}
}

func (ws *Routes) update(w http.ResponseWriter, r *http.Request) {
	m, ok := ws.checkExistsAll(r.URL.Query(), "key", "value")
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "400 Bad Request - faltando parâmetro 'key' ou 'value'")
		return
	}

	ws.cd.Update(m["key"], m["value"])

	fmt.Fprintln(w, "ok")
}

func (ws *Routes) del(w http.ResponseWriter, r *http.Request) {
	m, ok := ws.checkExistsAll(r.URL.Query(), "key")
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "400 Bad Request - faltando parâmetro 'key'")
		return
	}

	err := ws.cd.Del(m["key"])

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, "chave não encontrada")
	} else {
		fmt.Fprintln(w, "ok")
	}
}

func (ws *Routes) getAll(w http.ResponseWriter, r *http.Request) {
	pairs := ws.cd.GetAll()
	content := ""

	for _, pair := range pairs {
		content += pair[0] + " => " + pair[1] + "\n"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, content)
}

func (ws *Routes) delAll(w http.ResponseWriter, r *http.Request) {
	ws.cd.DelAll()
	fmt.Fprintln(w, "ok")
}

func GetRoutes() *http.ServeMux {
	r := NewRoutes()

	r.cd.Add("curso", "go")
	r.cd.Add("turma", "8058")
	r.cd.Add("mês", "março")

	mux := http.NewServeMux()

	mux.HandleFunc("/", r.index)
	mux.HandleFunc("/add", r.add)
	mux.HandleFunc("/get", r.get)
	mux.HandleFunc("/update", r.update)
	mux.HandleFunc("/del", r.del)
	mux.HandleFunc("/get-all", r.getAll)
	mux.HandleFunc("/del-all", r.delAll)

	return mux
}
