package flyWeightPattern

import "fmt"

type SoldierFactory struct {
	soldiers map[string]Soldier
}

func NewSoldierFactory() *SoldierFactory {
	return &SoldierFactory{soldiers: make(map[string]Soldier)}
}

func (sf *SoldierFactory) GetSoldier(soldierType string) Soldier {
	if sf.soldiers[soldierType] == nil {
		switch soldierType {
		case "sniper":
			fmt.Println("Creating a sniper")
			sf.soldiers[soldierType] = &Sniper{}
		case "infantry":
			fmt.Println("Creating an infantry")
			sf.soldiers[soldierType] = &Infantry{}
		}
	}
	return sf.soldiers[soldierType]
}

// example
//soldierFactory := flyWeightPattern.NewSoldierFactory()
//
//soldier1 := soldierFactory.GetSoldier("sniper")
//soldier1.Report(10, 20)
//
//soldier2 := soldierFactory.GetSoldier("sniper")
//soldier2.Report(15, 25)
//
//soldier3 := soldierFactory.GetSoldier("infantry")
//soldier3.Report(30, 40)
//
//soldier4 := soldierFactory.GetSoldier("infantry")
//soldier4.Report(35, 45)
