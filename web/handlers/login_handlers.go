package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/actions"
	"bones/web/forms"
	"bones/web/templating"
	"net/http"
	"strconv"
)

type ProfileContext struct {
	*templating.BaseContext
	User *entities.User
}

func LoadLoginPage(res http.ResponseWriter, req *http.Request) {
	actions.RenderPage(res, newLoginContext())
}

func CreateNewSession(res http.ResponseWriter, req *http.Request) {
	form := &forms.LoginForm{Request: req}
	err := actions.ProcessForm(req, form)

	if err != nil {
		actions.RenderPageWithErrors(res, newLoginContext(), err)

		return
	}

	url := routeURL("userProfile", "id", strconv.Itoa(form.User.Id))
	http.Redirect(res, req, url, 302)
}

func LoadUserProfilePage(res http.ResponseWriter, req *http.Request) {
	entity := actions.FindEntityOr404(res, req, repositories.Users, ":id")

	if user, ok := entity.(*entities.User); ok {
		ctx := ProfileContext{templating.NewBaseContext("profile.html"), user}
		actions.RenderPage(res, &ctx)
	}
}

func newLoginContext() *templating.BaseContext {
	return templating.NewBaseContext("login.html")
}
