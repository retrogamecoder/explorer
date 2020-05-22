package model

import "time"

type SpriteID int

type World struct {
	Regions []*Region
}

type Region struct {
	World   *World
	Section [][]*Section
	Tileset *Tileset
}

type Section struct {
	TileMap *TileMap
	Region  *Region
}

type Tileset struct {
}

type TileMap struct {
}

type Orientation uint8

const (
	OrientationUp Orientation = iota
	OrientationDown
	OrientationLeft
	OrientationRight
)

type Entity struct {
	X           int
	Y           int
	Orientation Orientation
	Health      int
	MaxHealth   int
	Sprite      SpriteID
}

type Monster struct {
	Entity
	MonsterClass *MonsterClass
	Behaviour    Behaviour
	Drops        DropSet
}

type MonsterClass struct {
	MaxHealth int
	Sprite    SpriteID
	Behaviour Behaviour
	Drops     DropSet
}

type Behaviour interface {
	Update(time.Duration, *Section)
}

type DropType uint8

type Drop struct {
	Entity
	Type     DropType
	Quantity int
}

type DropProbability struct {
	Drop        Drop
	Probability int // 0-100
}

type DropSet []DropProbability

type Player struct {
	Entity
	CurrentSection *Section
	Inventory      Inventory
}

type WeaponType uint8
type ArmourType uint8
type ItemType uint8

type Inventory struct {
	Weapon     WeaponType
	Armour     ArmourType
	Items      map[ItemType]bool
	ActiveItem ItemType
	Medals     []bool

	Bombs  int
	Arrows int
	Money  int
	Keys   int
}

type Projectile struct {
	Entity
	Behaviour ProjectileBehaviour
}

type ProjectileBehaviour interface {
	Update(time.Duration)
}
