package entity

type RigidBody struct {
	resilient float64
}

func NewRigidBody(resilient float64) *RigidBody {
	return &RigidBody{resilient}
}

func (r RigidBody) GetResilient() float64 {
	return r.resilient
}
