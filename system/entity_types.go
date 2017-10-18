package system

import (
	. "github.com/rickn42/arcade2d"
)

type position interface {
	GetPosition() Vec2
	SetPosition(Vec2)
	AddPosition(Vec2)
}

type linearVelociter interface {
	GetLinearVelocity() Vec2
	SetLinearVelocity(Vec2)
	AddLinearVelocity(Vec2)
}

type angler interface {
	GetAngle() float64
}

type angularVelociter interface {
	GetAngularVelocity() float64
	SetAngleVelocity(float64)
	AddAngleVelocity(float64)
}

type collider interface {
	ColliderShape() interface{}
	CheckEnter(this, other Entity)
	CheckExit(this, other Entity)
	CallEventCallbacks(this Entity)
}

type masser interface {
	GetMass() float64
	GetResilient() float64
}

type moved interface {
	IsMoved() bool
	ResetMoved()
}

type rigidbody interface {
	RigidBodyTag()
}

type gravity interface {
	GravityDirection(position Vec2) (direction Vec2)
}
