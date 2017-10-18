package entity

type RigidBody struct{}

func NewRigidBody() *RigidBody {
	return &RigidBody{}
}

func (*RigidBody) String() string {
	return "RigidBody"
}

func (*RigidBody) RigidBodyTag() {}
