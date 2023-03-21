package controllers

import (
	asm "assembler/boundaries"
	"bufio"
	"log"
	"os"
)

type CommandlineController struct {
	filePath               string
	instructions           asm.InstructionsInputData
	translatedInstructions asm.InstructionsOuputData
}

func MakeCommandlineController() CommandlineController {
	return CommandlineController{}
}

func (controller *CommandlineController) getInstructions() {
	if len(os.Args) == 1 {
		panic("Specify the Instructions source file")
	}

	instructionsFilePath := os.Args[1]
	controller.filePath = instructionsFilePath
	f, err := os.Open(instructionsFilePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var instructions []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		instructions = append(instructions, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	controller.instructions = asm.InstructionsInputData{
		Instructions: instructions,
	}
}

func (controller *CommandlineController) GetInstructions() {
	controller.getInstructions()
}

func (controller *CommandlineController) GetTranslatedInstructions(assembler asm.InstructionsInputBoundary) {
	controller.translatedInstructions = assembler.TranslateInstructions(controller.instructions)
}
