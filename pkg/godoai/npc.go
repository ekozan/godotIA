package godoai

import (
	"math"

	"github.com/godot-go/godot-go/pkg/log"

	"github.com/godot-go/godot-go/pkg/gdnative"
)

type NPC struct {
	gdnative.KinematicBody2DImpl
	gdnative.UserDataIdentifiableImpl

	speed      gdnative.Variant
	screenSize gdnative.Vector2
}

func (p *NPC) ClassName() string {
	return "NPC"
}

func (p *NPC) BaseClass() string {
	return "KinematicBody2D"
}

func (p *NPC) Init() {
	p.speed = gdnative.NewVariantReal(10)
}

func (p *NPC) OnClassRegistered(e gdnative.ClassRegisteredEvent) {
	// methods
	e.RegisterMethod("_ready", "Ready")
	//e.RegisterMethod("_process", "Process")
	e.RegisterMethod("_physics_process", "PhysicProcess")
	e.RegisterMethod("start", "Start")
	//e.RegisterMethod("_on_Player_body_entered", "OnPlayerBodyEntered")

	// signals
	//e.RegisterSignal("hit")

	// properties
	e.RegisterProperty("speed", "SetSpeed", "GetSpeed", p.speed)
}

func (p *NPC) Ready() {
	rect := p.GetViewportRect()

	// directly setting the screen_size is fine since this
	// isn't exposed as a Godot property
	p.screenSize = rect.GetSize()
	//p.Hide()
}

func (p *NPC) getInputs() gdnative.Vector2 {

	input := gdnative.GetSingletonInput()

	x := input.GetActionStrength("move_right") - input.GetActionStrength("move_left")
	y := input.GetActionStrength("move_down") - input.GetActionStrength("move_up")

	velocity := gdnative.NewVector2(x, y)
	vn := velocity.Normalized()
	return vn.OperatorMultiplyScalar(float32(p.speed.AsReal()))
}

func (p *NPC) PhysicProcess(delta float64) {

	animatedSprite := gdnative.NewAnimatedSpriteWithOwner(p.GetNode(gdnative.NewNodePath("AnimatedSprite")).GetOwnerObject())
	animatedSprite.Stop()
	//var velocity = gdnative.NewVector2(0, 0)
	var velocity = p.getInputs()
	v1 := velocity.Normalized()
	velocity = v1.OperatorMultiplyScalar(float32(p.speed.AsReal()))
	if velocity.Length() <= 0 {
		log.Debug("no move")
	} else {
		velocity = velocity.OperatorMultiplyScalar(float32(delta))
		str := velocity.AsString()
		log.Debug(str.AsGoString())
		p.MoveAndCollide(velocity, true, true, true)

		//if collision != nil {
		//log.Debug("I collided with " + collision.GetCollider().ToString())
		// /}
	}
	//log.Debug(nv.GetX())
	// if velocity.Length() > 0 {
	// 	//v1 := velocity.Normalized()
	// 	//velocity = v1.OperatorMultiplyScalar(float32(p.speed.AsReal()))
	// 	animatedSprite.Play("", false)
	// } else {
	// 	animatedSprite.Stop()
	// }

	//pos := p.GetPosition()
	//newPos := pos.OperatorAdd(velocity.OperatorMultiplyScalar(float32(delta)))
	//newPos.SetX(clamp(newPos.GetX(), 0, p.screenSize.GetX()))
	//newPos.SetY(clamp(newPos.GetY(), 0, p.screenSize.GetY()))

	// velX := nv.GetX()
	// velY := nv.GetY()

	// if velX != 0 && velX < 0 {
	// 	animatedSprite.SetAnimation("left")
	// } else if velX != 0 && velX > 0 {
	// 	animatedSprite.SetAnimation("right")
	// } else if velY != 0 && velY > 0 {
	// 	animatedSprite.SetAnimation("down")
	// } else if velY != 0 && velY < 0 {
	// 	animatedSprite.SetAnimation("up")
	// }
}

func (p *NPC) Start(pos gdnative.Vector2) {
	p.SetPosition(pos)
	p.Show()
	collisionShape2D := gdnative.NewCollisionShape2DWithOwner(p.GetNode(gdnative.NewNodePath("CollisionShape2D")).GetOwnerObject())
	collisionShape2D.SetDisabled(false)
}

func (p *NPC) OnPlayerBodyEntered(_body interface{}) {
	//p.Hide()
	//p.EmitSignal("hit")
	log.Debug("colision")
	//collisionShape2D := gdnative.NewCollisionShape2DWithOwner(p.GetNode(gdnative.NewNodePath("CollisionShape2D")).GetOwnerObject())
	//collisionShape2D.SetDeferred("disabled", gdnative.NewVariantBool(true))
}

func (p *NPC) GetSpeed() gdnative.Variant {
	return p.speed
}

func (p *NPC) SetSpeed(v gdnative.Variant) {
	newSpeed := v.AsReal()

	p.speed.Destroy()
	p.speed = gdnative.NewVariantReal(newSpeed)
}

func clamp(v, min, max float32) float32 {
	return float32(math.Max(math.Min(float64(v), float64(max)), float64(min)))
}

func NewNpcWithOwner(owner *gdnative.GodotObject) NPC {
	inst := gdnative.GetCustomClassInstanceWithOwner(owner).(*NPC)
	return *inst
}

func init() {
	gdnative.RegisterInitCallback(func() {
		gdnative.RegisterClass(&NPC{})
	})
}
