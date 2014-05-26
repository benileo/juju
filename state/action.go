package state

import ()

type actionDoc struct {
	Id string `bson:"_id"`

	// Name identifies the action; it should match an action defined by
	// the unit's charm.
	Name string

	// Payload holds the action's parameters, if any; it should validate
	// against the schema defined by the named action in the unit's charm
	Payload map[string]interface{}
}

// Action represents an instruction to do some "action" and is expected
// to match an action definition in a charm.
type Action struct {
	st  *State
	doc actionDoc
}

func newAction(st *State, adoc actionDoc) *Action {
	return &Action{
		st:  st,
		doc: adoc,
	}
}

// Name returns the name of the Action
func (a *Action) Name() string {
	return a.doc.Name
}

// Id returns the id of the Action
func (a *Action) Id() string {
	return a.doc.Id
}

// Payload will contain a structure representing arguments or parameters to
// an action, and is expected to be validated by the Unit using the Charm
// definition of the Action
func (a *Action) Payload() map[string]interface{} {
	return a.doc.Payload
}