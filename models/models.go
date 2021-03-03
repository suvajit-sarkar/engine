package models

// UnitTypeCount holds the constant for unit type count;
const UnitTypeCount = 4

//UnitType arry for unit types
var UnitType = []string{"Meele", "Range", "Mounted", "Seige"}

// Units describes all possible units as a group
type Units struct {
	MeeleUnits        []*MeeleUnit
	RangeUnits        []*RangeUnit
	MountedUnits      []*MountedUnit
	SeigeUnits        []*SeigeUnit
	TotalHealth       [UnitTypeCount]float64 //Dervied property
	formationType     string                 //TBD
	StrategyPofloat64 float64                //Could be an inherit property from the origin of city
}

//CreateUnits creates a group of unit placeholder
func CreateUnits(formation string) *Units {
	return &Units{
		formationType: formation,
	}
}

type customInterface interface {
}

// AppendToMeeleUnit is Struct member function to update unit members
func (u Units) AppendToMeeleUnit(meeleUnit *MeeleUnit) *Units {
	u.MeeleUnits = append(u.MeeleUnits, meeleUnit)
	u.TotalHealth[0] = u.TotalHealth[0] + meeleUnit.HealthBar
	return &u
}

// AppendToRangeUnit is Struct member function to update unit members
func (u Units) AppendToRangeUnit(rangeUnit *RangeUnit) *Units {
	u.RangeUnits = append(u.RangeUnits, rangeUnit)
	u.TotalHealth[1] = u.TotalHealth[1] + rangeUnit.HealthBar
	return &u
}

// AppendToMountedUnit is Struct member function to update unit members
func (u Units) AppendToMountedUnit(mountedUnit *MountedUnit) *Units {
	u.MountedUnits = append(u.MountedUnits, mountedUnit)
	u.TotalHealth[2] = u.TotalHealth[2] + mountedUnit.HealthBar
	return &u
}

// AppendToSeigeUnit is Struct member function to update unit members
func (u Units) AppendToSeigeUnit(seigeUnit *SeigeUnit) *Units {
	u.SeigeUnits = append(u.SeigeUnits, seigeUnit)
	u.TotalHealth[3] = u.TotalHealth[3] + seigeUnit.HealthBar
	return &u
}

//MeeleUnit properties
type MeeleUnit struct {
	TypeID           int
	BaseAttack       float64
	BaseDefence      float64
	HealthBar        float64
	AttackToRange    float64
	AttackToMounted  float64
	AttackToSeige    float64
	DefenceToRange   float64
	DefenceToMounted float64
	DefenceToSeige   float64
	MovementSpeed    float64
}

//CreateMeeleUnit creates and returns a basic meele unit
func CreateMeeleUnit() *MeeleUnit {
	return &MeeleUnit{
		TypeID:           1,
		BaseAttack:       60,
		BaseDefence:      4,
		AttackToRange:    0,
		AttackToMounted:  5,
		AttackToSeige:    15,
		DefenceToRange:   0,
		DefenceToMounted: 0,
		DefenceToSeige:   0,
		MovementSpeed:    0.1,
		HealthBar:        350,
	}

}

//RangeUnit properties
type RangeUnit struct {
	TypeID           int
	BaseAttack       float64
	BaseDefence      float64
	AttackToMounted  float64
	AttackToMeele    float64
	AttackToSeige    float64
	DefenceToMounted float64
	DefenceToMeele   float64
	DefenceToSeige   float64
	HealthBar        float64
	MovementSpeed    float64
}

//CreateRangeUnit creates and returns a basic meele unit
func CreateRangeUnit() *RangeUnit {
	return &RangeUnit{
		TypeID:           1,
		BaseAttack:       70,
		BaseDefence:      3,
		AttackToMeele:    10,
		AttackToMounted:  15,
		AttackToSeige:    5,
		DefenceToMeele:   5,
		DefenceToMounted: 5,
		DefenceToSeige:   5,
		HealthBar:        300,
		MovementSpeed:    0.11,
	}

}

//MountedUnit properties
type MountedUnit struct {
	TypeID         int
	BaseAttack     float64
	BaseDefence    float64
	AttackToMeele  float64
	AttackToSeige  float64
	AttackToRange  float64
	DefenceToMeele float64
	DefenceToSeige float64
	DefenceToRange float64
	MovementSpeed  float64
	HealthBar      float64
}

//CreateMountedUnit creates and returns a basic meele unit
func CreateMountedUnit() *MountedUnit {
	return &MountedUnit{
		TypeID:         1,
		BaseAttack:     120,
		BaseDefence:    5,
		AttackToMeele:  10,
		AttackToRange:  15,
		AttackToSeige:  10,
		DefenceToMeele: 5,
		DefenceToRange: 5,
		DefenceToSeige: 5,
		MovementSpeed:  0.15,
		HealthBar:      500,
	}

}

//SeigeUnit properties
type SeigeUnit struct {
	TypeID          int
	BaseBreakAttack float64
	BaseAttack      float64
	BaseDefence     float64
	MovementSpeed   float64
	ManRequired     int
	HealthBar       float64
}

//CreateSeigeUnit creates and returns a basic meele unit
func CreateSeigeUnit() *SeigeUnit {
	return &SeigeUnit{
		TypeID:          1,
		BaseAttack:      20,
		BaseBreakAttack: 100,
		BaseDefence:     4,
		MovementSpeed:   0.1,
		HealthBar:       600,
	}
}
