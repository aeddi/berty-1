package bertybridge

const (
	AppStateActive = iota
	AppStateInactive
	AppStateBackground
)

type AppState struct{ Active, Inactive, Background int }

var AppStateEnum = &AppState{AppStateActive, AppStateInactive, AppStateBackground}

type LifeCycleDriver interface {
	GetCurrentState() int
	RegisterHandler(handler AppStateHandler) bool
}

type AppStateHandler interface {
	Handler(appstate int)
}

// noop driver

type noopLifeCycleDriver struct{}

func (*noopLifeCycleDriver) GetCurrentState() int                   { return AppStateEnum.Active }
func (*noopLifeCycleDriver) RegisterHandler(_ AppStateHandler) bool { return true }
