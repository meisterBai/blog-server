package routers

import "github.com/kataras/iris"

type IRouter interface {
	RegisterHandlers(partyName string)
}

type Router struct {
	Iris *iris.Application
}

func (r Router) GetIrisParty(groupName string) iris.Party {
	party := r.Iris.Party(groupName)
	return party
}


