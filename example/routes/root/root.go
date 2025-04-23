package root

import "net/http"

type (
	rootImp struct{}
)

func (r2 *rootImp) GetRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/oas", http.StatusMovedPermanently)
}

func New() ServerInterface {
	return &rootImp{}
}
