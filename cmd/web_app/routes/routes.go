package routes

import (
	"net/http"
	"web_app/helpers"
	"web_app/render"
)

type RoutesHandler struct {
	Routes helpers.GRouts
}

type PageData struct {
	Title    string
	Username string
	IsAuth   bool
}

func NewRoutesHandler(cfg *helpers.GRouts) *RoutesHandler {
	return &RoutesHandler{
		Routes: *cfg,
	}
}

func (rs *RoutesHandler) PreparingRoutes() {
	http.HandleFunc("/", rs.HIndex)
	http.HandleFunc("/signup", rs.HSignUp)
	http.HandleFunc("/signin", rs.HSignIn)
	http.HandleFunc("/run", rs.HRun)
}

func (rs *RoutesHandler) HIndex(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title:    "The Game | Home",
		Username: "",
		IsAuth:   false,
	}

	tokenCookie, _ := r.Cookie("token")
	if tokenCookie != nil {
		claims := helpers.ValidateToken(tokenCookie.Value)
		if claims != nil {
			data.Username = claims.Username
			data.IsAuth = true
		}
	}

	render.TmplRender(w, "home.html", data)
}

func (rs *RoutesHandler) HSignUp(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		payload := map[string]string{
			"email":    r.FormValue("email"),
			"username": r.FormValue("username"),
			"password": r.FormValue("password"),
		}

		resp, _ := http.Post(rs.Routes.GRoutes[0].URL, "application/json", helpers.JsonEncoder(payload))

		helpers.SetCookie(helpers.JsonDecoder(resp), w)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	render.TmplRender(w, "signup.html", map[string]string{"Title": "The Game | Sign Up"})
}

func (rs *RoutesHandler) HSignIn(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		payload := map[string]string{
			"email":    r.FormValue("email"),
			"password": r.FormValue("password"),
		}

		resp, _ := http.Post(rs.Routes.GRoutes[1].URL, "application/json", helpers.JsonEncoder(payload))

		helpers.SetCookie(helpers.JsonDecoder(resp), w)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	render.TmplRender(w, "signin.html", map[string]string{"Title": "The Game | Sign In"})
}

func (rs *RoutesHandler) HRun(w http.ResponseWriter, r *http.Request) {

	helpers.ResetCookie(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
