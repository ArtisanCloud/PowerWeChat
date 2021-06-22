package support

type Observable struct {
	handlers []func()


}

func NewObservable()*Observable{
	return &Observable{}
}