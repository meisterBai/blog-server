package routers

import (
	"github.com/kataras/iris"
	"server/services/article"
	"server/services/article/usecase"
	"server/services/article/repository"
	"server/storage/orm"
	"time"
	"server/config"
)

type ArticleRouter struct {
	Router
}

type ArticleHandler struct {
	Usecase article.Usecase
}

func NewArticleHandler(usecase article.Usecase) *ArticleHandler {
	return &ArticleHandler{Usecase:usecase}
}

func NewArticleRouter(app *iris.Application) IRouter {
	r := &ArticleRouter{}
	r.Iris = app
	return r
}


func (a *ArticleRouter) RegisterHandlers(groupName string) {
	party := a.GetIrisParty(groupName)

	repo := repository.NewMysqlArticleRepository(orm.GetDBService().DB)

	timeoutContext := time.Duration(config.Conf.Get("time_out").(int)) * time.Second

	articleUsecase := usecase.NewArticleUsecase(repo, timeoutContext)


	handler := &ArticleHandler{articleUsecase}

	party.Get("", handler.FetchArticle)
	party.Get("/{id:uint}", handler.GetByID)
	party.Post("", handler.Store)
	party.Delete("/{id:uint}", handler.Delete)
}

func (a *ArticleHandler) FetchArticle(c iris.Context) {

}

func (a *ArticleHandler) GetByID(c iris.Context) {
	id, _ := c.Params().GetUint("id")

}

func (a *ArticleHandler) Store(c iris.Context) {

}

func (a *ArticleHandler) Delete(c iris.Context) {
	id, _ := c.Params().GetUint("id")

}
