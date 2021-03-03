package main

import (
	"math"

	"github.com/suvajit-sarkar/engine/models"
	"github.com/suvajit-sarkar/engine/utilities"
)

var combatOver bool

func calculateMeeleUnitBaseProp(unit *models.Units) []float64 {
	var attackMeele = 0.0
	var defendMeele = 0.0

	var attackRange = 0.0
	var defendRange = 0.0

	var attackMounted = 0.0
	var defendMounted = 0.0

	var attackSeige = 0.0
	var defendSeige = 0.0

	for _, u := range unit.MeeleUnits {
		//
		attackMeele += u.BaseAttack
		defendMeele += u.BaseDefence
		//
		attackRange += (u.BaseAttack + u.AttackToRange)
		defendRange += (u.BaseDefence + u.DefenceToRange)
		//
		attackMounted += (u.BaseAttack + u.AttackToMounted)
		defendMounted += (u.BaseDefence + u.DefenceToMounted)
		//
		attackSeige += (u.BaseAttack + u.AttackToSeige)
		defendSeige += (u.BaseDefence + u.DefenceToSeige)
	}

	prop := []float64{attackMeele, attackRange, attackMounted, attackSeige, defendMeele, defendRange, defendMounted, defendSeige}

	return prop
}
func calculateRangeUnitBaseProp(unit *models.Units) []float64 {
	var attackMeele = 0.0
	var defendMeele = 0.0

	var attackRange = 0.0
	var defendRange = 0.0

	var attackMounted = 0.0
	var defendMounted = 0.0

	var attackSeige = 0.0
	var defendSeige = 0.0

	for _, u := range unit.RangeUnits {
		//
		attackMeele += (u.BaseAttack + u.AttackToMeele)
		defendMeele += (u.BaseDefence + u.DefenceToMeele)
		//
		attackRange += (u.BaseAttack)
		defendRange += (u.BaseDefence)
		//
		attackMounted += (u.BaseAttack + u.AttackToMounted)
		defendMounted += (u.BaseDefence + u.DefenceToMounted)
		//
		attackSeige += (u.BaseAttack + u.AttackToSeige)
		defendSeige += (u.BaseDefence + u.DefenceToSeige)
	}

	prop := []float64{attackMeele, attackRange, attackMounted, attackSeige, defendMeele, defendRange, defendMounted, defendSeige}

	return prop
}
func calculateMountedUnitBaseProp(unit *models.Units) []float64 {
	var attackMeele = 0.0
	var defendMeele = 0.0

	var attackRange = 0.0
	var defendRange = 0.0

	var attackMounted = 0.0
	var defendMounted = 0.0

	var attackSeige = 0.0
	var defendSeige = 0.0

	for _, u := range unit.MountedUnits {
		//
		attackMeele += (u.BaseAttack + u.AttackToMeele)
		defendMeele += (u.BaseDefence + u.DefenceToMeele)
		//
		attackRange += (u.BaseAttack + u.AttackToRange)
		defendRange += (u.BaseDefence + u.DefenceToRange)
		//
		attackMounted += (u.BaseAttack)
		defendMounted += (u.BaseDefence)
		//
		attackSeige += (u.BaseAttack + u.AttackToSeige)
		defendSeige += (u.BaseDefence + u.DefenceToSeige)
	}

	prop := []float64{attackMeele, attackRange, attackMounted, attackSeige, defendMeele, defendRange, defendMounted, defendSeige}

	return prop
}
func calculateSeigeUnitBaseProp(unit *models.Units) []float64 {
	var attack = 0.0
	var defend = 0.0

	for _, u := range unit.SeigeUnits {
		//
		attack += u.BaseAttack
		defend += u.BaseDefence
		//
	}

	prop := []float64{attack, attack, attack, attack, defend, defend, defend, defend}

	return prop
}
func calculateUnitHealth(u *models.Units) []float64 {
	propHealth := []float64{0, 0, 0, 0, 0, 0, 0, 0}
	uSize := calculateUnitSize(u)
	for i := 0; i < len(uSize); i++ {
		if uSize[i] > 0 {
			propHealth[i] = u.TotalHealth[i]
			propHealth[i+4] = u.TotalHealth[i] / float64(uSize[i])
		}
	}
	return propHealth
}
func calculateUnitSize(u *models.Units) []int {
	return []int{len(u.MeeleUnits), len(u.RangeUnits), len(u.MountedUnits), len(u.SeigeUnits)}
}

func calculateUnitLeft() {}

//UVU does the unit vs unit
func UVU(tA []*models.Units, tB []*models.Units) ([]*models.Units, []*models.Units) {

	//Flag for combat control
	combatOver = false

	tAHealthProp := [models.UnitTypeCount * 2]float64{}
	tBHealthProp := [models.UnitTypeCount * 2]float64{}

	for _, unit := range tA {
		prop := calculateUnitHealth(unit)
		for i := 0; i < len(prop); i++ {
			tAHealthProp[i] += prop[i]
		}

	}
	for _, unit := range tB {
		prop := calculateUnitHealth(unit)
		for i := 0; i < len(prop); i++ {
			tBHealthProp[i] += prop[i]
		}

	}
	//Combat rounds start from here
	var rounds = 0
	for i := 0; i < models.UnitTypeCount; i++ {
		switch unitType := i; unitType {
		case 0:
			println("tA Meele ", len(tA[0].MeeleUnits))
		case 1:
			println("tA Range ", len(tA[0].RangeUnits))
		case 2:
			println("tA Mounted", len(tA[0].MountedUnits))
		case 3:
			println("tA Seige", len(tA[0].SeigeUnits))
		}
	}
	for i := 0; i < models.UnitTypeCount; i++ {
		switch unitType := i; unitType {
		case 0:
			println("tB Meele  ", len(tB[0].MeeleUnits))
		case 1:
			println("tB Range  ", len(tB[0].RangeUnits))
		case 2:
			println("tB Mounted  ", len(tB[0].MountedUnits))
		case 3:
			println("tB Seige  ", len(tB[0].SeigeUnits))
		}
	}
	for combatOver == false {
		rounds++
		var tATotalHealth = 0.0
		var tBTotalHealth = 0.0
		tASize := [models.UnitTypeCount]int{}
		tBSize := [models.UnitTypeCount]int{}

		tABaseProp := [models.UnitTypeCount][models.UnitTypeCount * 2]float64{}
		tBBaseProp := [models.UnitTypeCount][models.UnitTypeCount * 2]float64{}

		for _, unit := range tA {
			uSize := calculateUnitSize(unit)
			propMe := calculateMeeleUnitBaseProp(unit)
			propR := calculateRangeUnitBaseProp(unit)
			propMo := calculateMountedUnitBaseProp(unit)
			propSe := calculateSeigeUnitBaseProp(unit)
			for i := 0; i < len(propMe); i++ {
				tABaseProp[0][i] += propMe[i]
				tABaseProp[1][i] += propR[i]
				tABaseProp[2][i] += propMo[i]
				tABaseProp[3][i] += propSe[i]
			}
			for i := 0; i < len(uSize); i++ {
				tASize[i] += uSize[i]
			}
		}
		for _, unit := range tB {
			uSize := calculateUnitSize(unit)
			propMe := calculateMeeleUnitBaseProp(unit)
			propR := calculateRangeUnitBaseProp(unit)
			propMo := calculateMountedUnitBaseProp(unit)
			propSe := calculateSeigeUnitBaseProp(unit)
			for i := 0; i < len(propMe); i++ {
				tBBaseProp[0][i] += propMe[i]
				tBBaseProp[1][i] += propR[i]
				tBBaseProp[2][i] += propMo[i]
				tBBaseProp[3][i] += propSe[i]
			}
			for i := 0; i < len(uSize); i++ {
				tBSize[i] += uSize[i]
			}

		}
		// fmt.Println(tABaseProp, " ")
		// fmt.Println(tBBaseProp, " ")
		// fmt.Println(tAHealthProp, " ", tATotalHealth)
		// fmt.Println(tBHealthProp, " ", tBTotalHealth)
		for i := 0; i < len(tASize); i++ {
			if (tASize[i]) > 0 {
				for j := 0; j < models.UnitTypeCount; j++ {
					if tBSize[j] > 0 {
						tAHealthProp[i] = tAHealthProp[i] - (utilities.MaxFloat((tBBaseProp[j][i] - tABaseProp[i][j+models.UnitTypeCount]), 0))
						// println(i, "->", j)
						// println(i, "->", j, "attack done ->", tABaseProp[i][j])
						// fmt.Printf("%f", tBBaseProp[j][i+models.UnitTypeCount])
						// fmt.Printf("%f", utilities.MaxFloat((tABaseProp[i][j]-tBBaseProp[j][i+models.UnitTypeCount]), 0))
						tBHealthProp[j] = tBHealthProp[j] - (utilities.MaxFloat((tABaseProp[i][j] - tBBaseProp[j][i+models.UnitTypeCount]), 0))
						if tBHealthProp[j] < 0 {
							tBHealthProp[j] = 0
						}
						if tAHealthProp[i] < 0 {
							tAHealthProp[i] = 0
						}
						//println(int(math.Ceil(utilities.MaxFloat(tBHealthProp[i], 0) / tBHealthProp[i+unitTypeCount])))
					}
				}

			}

		}
		//Outcome for tA
		//All Units
		for i := 0; i < models.UnitTypeCount; i++ {
			tATotalHealth += utilities.MaxFloat(tAHealthProp[i], 0)
			unitLeft := int(math.Ceil(utilities.MaxFloat(tAHealthProp[i], 0) / tAHealthProp[i+models.UnitTypeCount]))
			//fmt.Println(models.UnitType[i], " unit of tA killed: ", utilities.MaxInt(tASize[i]-unitLeft, 0))
			switch unitType := i; unitType {
			case 0:
				tA[0].MeeleUnits = tA[0].MeeleUnits[utilities.MaxInt(tASize[i]-unitLeft, 0):]
			case 1:
				tA[0].RangeUnits = tA[0].RangeUnits[utilities.MaxInt(tASize[i]-unitLeft, 0):]
			case 2:
				tA[0].MountedUnits = tA[0].MountedUnits[utilities.MaxInt(tASize[i]-unitLeft, 0):]
			case 3:
				tA[0].SeigeUnits = tA[0].SeigeUnits[utilities.MaxInt(tASize[i]-unitLeft, 0):]
			}
		}
		//Outcome for tB
		//All Units
		for i := 0; i < models.UnitTypeCount; i++ {
			tBTotalHealth += utilities.MaxFloat(tBHealthProp[i], 0)
			unitLeft := int(math.Ceil(utilities.MaxFloat(tBHealthProp[i], 0) / tBHealthProp[i+models.UnitTypeCount]))
			//print(unitLeft)
			//fmt.Println(models.UnitType[i], " unit of tB killed: ", utilities.MaxInt(tBSize[i]-unitLeft, 0))
			switch unitType := i; unitType {
			case 0:
				tB[0].MeeleUnits = tB[0].MeeleUnits[utilities.MaxInt(tBSize[i]-unitLeft, 0):]
			case 1:
				tB[0].RangeUnits = tB[0].RangeUnits[utilities.MaxInt(tBSize[i]-unitLeft, 0):]
			case 2:
				tB[0].MountedUnits = tB[0].MountedUnits[utilities.MaxInt(tBSize[i]-unitLeft, 0):]
			case 3:
				tB[0].SeigeUnits = tB[0].SeigeUnits[utilities.MaxInt(tBSize[i]-unitLeft, 0):]
			}
		}

		if (tBTotalHealth) <= 0 || (tATotalHealth) <= 0 {
			combatOver = true
			println(rounds, "rounds")
			println("Combat over")
			for i := 0; i < models.UnitTypeCount; i++ {
				switch unitType := i; unitType {
				case 0:
					println("tA Meele Survived ", len(tA[0].MeeleUnits))
					println("tB Meele Survived ", len(tB[0].MeeleUnits))
				case 1:
					println("tA Mounted Survived ", len(tA[0].MountedUnits))
					println("tB Range Survived ", len(tB[0].RangeUnits))
				case 2:
					println("tA Range Survived ", len(tA[0].RangeUnits))
					println("tB Mounted Survived ", len(tB[0].MountedUnits))
				case 3:
					println("tA Seige Survived ", len(tA[0].SeigeUnits))
					println("tB Seige Survived ", len(tB[0].SeigeUnits))
				}
			}
			print()
		} else {
			combatOver = false
		}
	}
	return tA, tB

	// // Recalculate units left u1
	// if meeleAvgHealthU1 > 0 {
	// 	totalMeeleU1Left := int(math.Ceil(utilities.MaxFloat(meeleHealthU1, 0) / meeleAvgHealthU1))
	// 	totalMeeleUnitDiedU1 := totalMeeleU1 - totalMeeleU1Left
	// 	u1.MeeleUnits = u1.MeeleUnits[totalMeeleUnitDiedU1:]
	// }
	// if rangeAvgHealthU1 > 0 {
	// 	totalRangeU1Left := int(math.Ceil(utilities.MaxFloat(rangeHealthU1, 0) / rangeAvgHealthU1))
	// 	totalRangeUnitDiedU1 := totalRangeU1 - totalRangeU1Left
	// 	u1.RangeUnits = u1.RangeUnits[totalRangeUnitDiedU1:]
	// }

	// if mountedAvgHealthU1 > 0 {
	// 	totalMountedU1Left := int(math.Ceil(utilities.MaxFloat(mountedHealthU1, 0) / mountedAvgHealthU1))
	// 	totalMountedUnitDiedU1 := totalMountedU1 - totalMountedU1Left
	// 	u1.MeeleUnits = u1.MeeleUnits[totalMountedUnitDiedU1:]
	// }

	// if seigeAvgHealthU1 > 0 {
	// 	totalSeigeU1Left := int(math.Ceil(utilities.MaxFloat(seigeHealthU1, 0) / seigeAvgHealthU1))
	// 	totalSeigeUnitDiedU1 := totalSeigeU1 - totalSeigeU1Left
	// 	u1.SeigeUnits = u1.SeigeUnits[totalSeigeUnitDiedU1:]
	// }

	// // Recalculate units left u2
	// if meeleAvgHealthU2 > 0 {
	// 	totalMeeleU2Left := int(math.Ceil(utilities.MaxFloat(meeleHealthU2, 0) / meeleAvgHealthU2))
	// 	totalMeeleUnitDiedU2 := totalMeeleU2 - totalMeeleU2Left
	// 	u2.MeeleUnits = u2.MeeleUnits[totalMeeleUnitDiedU2:]
	// }
	// if rangeAvgHealthU2 > 0 {
	// 	totalRangeU2Left := int(math.Ceil(utilities.MaxFloat(rangeHealthU2, 0) / rangeAvgHealthU2))
	// 	totalRangeUnitDiedU2 := totalRangeU2 - totalRangeU2Left
	// 	u2.RangeUnits = u2.RangeUnits[totalRangeUnitDiedU2:]
	// }

	// if mountedAvgHealthU2 > 0 {
	// 	totalMountedU2Left := int(math.Ceil(utilities.MaxFloat(mountedHealthU2, 0) / mountedAvgHealthU2))
	// 	fmt.Print(mountedHealthU2)
	// 	totalMountedUnitDiedU2 := totalMountedU2 - totalMountedU2Left
	// 	u2.MountedUnits = u2.MountedUnits[totalMountedUnitDiedU2:]
	// }

	// if seigeAvgHealthU2 > 0 {
	// 	totalSeigeU2Left := int(math.Ceil(utilities.MaxFloat(seigeHealthU2, 0) / seigeAvgHealthU2))
	// 	totalSeigeUnitDiedU2 := totalSeigeU2 - totalSeigeU2Left
	// 	u2.SeigeUnits = u2.SeigeUnits[totalSeigeUnitDiedU2:]
	// }
	// if (meeleHealthU2+mountedHealthU2+rangeHealthU2+seigeHealthU2) <= 0 || (meeleHealthU1+mountedHealthU1+rangeHealthU1+seigeHealthU1) <= 0 {
	// 	combatOver = 1
	// }
}

func calculateUnitProperties() {

}

//PVP simulates player vs player 1v1 match up
func PVP(player1 *models.MeeleUnit, player2 *models.MeeleUnit) {
	for player1.HealthBar >= 0 && player2.HealthBar >= 0 {
		//Player 1 attacks player  2
		println("Damage dealt by player 1 : ", (player1.BaseAttack - player2.BaseDefence))
		player2.HealthBar = player2.HealthBar - (player1.BaseAttack - player2.BaseDefence)
		println("Health left player 2 ", player2.HealthBar)
		//Player 2 attacks player 1

		if player2.HealthBar <= 0 {
			break
		}
		println("Damage dealt by player 2 : ", (player2.BaseAttack - player1.BaseDefence))
		player1.HealthBar = player1.HealthBar - (player1.BaseAttack - player1.BaseDefence)
		println("Health left player 1 ", player1.HealthBar)

		if player1.HealthBar <= 0 {
			break
		}
	}
	if player1.HealthBar <= 0 {
		println("player 1 dies")
	} else {
		println("player 2 dies")
	}
}

//UVUSimulation tesing
func UVUSimulation() {
	const meeleUnitNumberU1 = 900
	const rangeUnitNumberU1 = 700
	const mountedUnitNumberU1 = 300
	const seigeUnitNumberU1 = 0
	u1 := models.CreateUnits("Normal")

	//attackingUnit := models.CreateMeeleUnit()

	for range [meeleUnitNumberU1]int{} {
		u1 = u1.AppendToMeeleUnit(models.CreateMeeleUnit())
	}
	for range [rangeUnitNumberU1]int{} {
		u1 = u1.AppendToRangeUnit(models.CreateRangeUnit())
	}
	for range [mountedUnitNumberU1]int{} {
		u1 = u1.AppendToMountedUnit(models.CreateMountedUnit())
	}
	for range [seigeUnitNumberU1]int{} {
		u1 = u1.AppendToSeigeUnit(models.CreateSeigeUnit())
	}

	const meeleUnitNumberU2 = 900
	const rangeUnitNumberU2 = 700
	const mountedUnitNumberU2 = 310
	const seigeUnitNumberU2 = 0
	u2 := models.CreateUnits("Normal")

	//attackingUnit := models.CreateMeeleUnit()

	for range [meeleUnitNumberU2]int{} {
		u2 = u2.AppendToMeeleUnit(models.CreateMeeleUnit())
	}
	for range [rangeUnitNumberU2]int{} {
		u2 = u2.AppendToRangeUnit(models.CreateRangeUnit())
	}
	for range [mountedUnitNumberU2]int{} {
		u2 = u2.AppendToMountedUnit(models.CreateMountedUnit())
	}
	for range [seigeUnitNumberU2]int{} {
		u2 = u2.AppendToSeigeUnit(models.CreateSeigeUnit())
	}
	const meeleUnitNumberU3 = 0
	const rangeUnitNumberU3 = 0
	const mountedUnitNumberU3 = 0
	const seigeUnitNumberU3 = 0
	u3 := models.CreateUnits("Normal")

	//attackingUnit := models.CreateMeeleUnit()

	for range [meeleUnitNumberU3]int{} {
		u3 = u3.AppendToMeeleUnit(models.CreateMeeleUnit())
	}
	for range [rangeUnitNumberU3]int{} {
		u3 = u3.AppendToRangeUnit(models.CreateRangeUnit())
	}
	for range [mountedUnitNumberU3]int{} {
		u3 = u3.AppendToMountedUnit(models.CreateMountedUnit())
	}
	for range [seigeUnitNumberU3]int{} {
		u3 = u3.AppendToSeigeUnit(models.CreateSeigeUnit())
	}
	_, _ = UVU([]*models.Units{u1}, []*models.Units{u2})

	// println("Meele units Left u1 > ", len(u1.MeeleUnits))
	// println("Range units Left u1 > ", len(u1.RangeUnits))
	// println("Mount units Left u1 > ", len(u1.MountedUnits))
	// println("Seige units Left u1 > ", len(u1.SeigeUnits))

	// println("Meele units Left u2 > ", len(u2.MeeleUnits))
	// println("Range units Left u2 > ", len(u2.RangeUnits))
	// fmt.Println("Range health Left u2 > ", u2.TotalHealthRanges)
	// println("Mount units Left u2 > ", len(u2.MountedUnits))
	// println("Seige units Left u1 > ", len(u2.SeigeUnits))
}

//PVPSimulation simulates player vs player combat
func PVPSimulation() {
	player1 := models.CreateMeeleUnit()
	player2 := models.CreateMeeleUnit()
	player2.BaseDefence = 15

	PVP(player1, player2)
}
