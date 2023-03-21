package main

import (
	con "assembler/controllers"
)

func main() {
	// fmt.Println("Executing Assembler")
	controller := con.MakeCommandlineController()
	controller.GetInstructions() // instructions without white space

	assembler := MakeAssemblyUseCaseInteractor()
	controller.GetTranslatedInstructions(&assembler)
}
