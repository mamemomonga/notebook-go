package objfunc

import (
	"fmt"
)

// ObjFunc is object struct
type ObjFunc struct {
	PublicValue  string
	privateValue string
}

// NewObjFunc is constructor
func NewObjFunc() (t *ObjFunc) {
	t = new(ObjFunc)
	return t
}

// Set is set private value
func (t *ObjFunc) Set(s string) {
	t.privateValue = s
}

// HelloPrivate is print private value
func (t *ObjFunc) HelloPrivate() {
	fmt.Printf("Hello %s\n", t.privateValue)
}

// HelloPublic is print public value
func (t *ObjFunc) HelloPublic() {
	fmt.Printf("Hello %s\n", t.PublicValue)
}
