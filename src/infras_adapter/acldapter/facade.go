/*
*

	@author: junwang
	@since: 2023/7/28
	@desc: //TODO

*
*/
package acldapter

type Facade struct {
	is *InternalSystem
}

func NewFacade() ExternalApi {
	return &Facade{is: &InternalSystem{}}
}
