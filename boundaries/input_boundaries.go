package boundaries

// Functions on the UseCaseInteractors
type InstructionsInputBoundary interface {
	TranslateInstructions(InstructionsInputData) InstructionsOuputData
}

// Instructions expected by the UseCaseInteractor
type InstructionsInputData struct {
	Instructions []string
}

