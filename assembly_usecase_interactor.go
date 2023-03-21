package main

import (
	asm "assembler/boundaries"
	enc "assembler/encoder"
	psr "assembler/parser"
)

type AssemblyUseCaseInteractor struct {
}

func MakeAssemblyUseCaseInteractor() AssemblyUseCaseInteractor {
	return AssemblyUseCaseInteractor{}
}

func (assembler *AssemblyUseCaseInteractor) TranslateInstructions(instructions asm.InstructionsInputData) asm.InstructionsOuputData {
	parser := psr.MakeParser()
	instructions_without_symbols := parser.RemoveWhiteSpaceAndProcesSymbols(instructions)

	encoder := enc.MakeEncoder()
	processed_instructions := encoder.EncodeToBinary(instructions_without_symbols)
	return asm.InstructionsOuputData{
		Instructions: processed_instructions,
	}
}
