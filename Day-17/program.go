package main

import (
	"fmt"
)

type TargetArea struct {
	MinX, MaxX, MinY, MaxY int
}

type Probe struct {
	XVelocity, YVelocity int
	XPos, YPos           int
}

func main() {
	//target := TargetArea{MinX: 20, MaxX: 30, MinY: -10, MaxY: -5} //test
	target := TargetArea{MinX: 94, MaxX: 151, MinY: -156, MaxY: -103} //real

	workingVelocityCombos := 0
	for xVelocity := 1; xVelocity < 152; xVelocity++ {
		for yVelocity := -156; yVelocity < 156; yVelocity++ {
			probe := &Probe{XVelocity: xVelocity, YVelocity: yVelocity}
			if probe.Fire(target) {
				workingVelocityCombos++
			}
		}
	}
	fmt.Println(workingVelocityCombos)
}

func (probe *Probe) IsInTargetArea(target TargetArea) bool {
	return probe.XPos >= target.MinX && probe.XPos <= target.MaxX && probe.YPos >= target.MinY && probe.YPos <= target.MaxY
}

func (probe *Probe) HasMissed(target TargetArea) bool {
	return probe.XPos > target.MaxX || probe.YPos < target.MinY
}

func (probe *Probe) Step(target TargetArea) {
	probe.XPos += probe.XVelocity
	probe.YPos += probe.YVelocity

	if probe.XVelocity > 0 {
		probe.XVelocity--
	} else if probe.XVelocity < 0 {
		probe.XVelocity++
	}

	probe.YVelocity--
	//fmt.Println("Probe at position", probe.XPos, probe.YPos)
}

func (probe *Probe) Fire(target TargetArea) bool {
	for {
		if probe.HasMissed(target) {
			return false
		}
		probe.Step(target)
		if probe.IsInTargetArea(target) {
			fmt.Println("Hit with probe", *probe)
			return true
		}
	}
}
