package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {

   return !archerIsAwake && prisonerIsAwake
}


func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
  return (prisonerIsAwake && !knightIsAwake && !archerIsAwake) || (petDogIsPresent && !archerIsAwake)
}
