package entity

import "fmt"

type RigidBody struct {
	resilient float64
}

func NewRigidBody(resilient float64) *RigidBody {
	return &RigidBody{resilient}
}

func (r *RigidBody) String() string {
	return fmt.Sprintf("RigidBody resilent %.1f", r.resilient)
}

func (r *RigidBody) GetResilient() float64 {
	return r.resilient
}
