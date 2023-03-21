package parser

import (
	asm "assembler/boundaries"
	"fmt"
	"regexp"
	"strings"
)

type Parser struct {
	symbolTable        map[string]int
	InputInstructions  asm.InstructionsInputData // wrong use of entity
	OutputInstructions asm.InstructionsOuputData // wrong use of entity
}

func MakeParser() Parser {
	parser := Parser{}
	parser.initSymbolTable()
	return parser
}

func (parser *Parser) initSymbolTable() {
	// add predefined symbols into the symbolTable map
	parser.symbolTable = map[string]int{"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4,
		"R0": 0, "R1": 1, "R2": 2, "R3": 3, "R4": 4, "R5": 5, "R6": 6, "R7": 7,
		"R8": 8, "R9": 9, "R10": 10, "R11": 11, "R12": 12, "R13": 13, "R14": 14, "R15": 15,
		"SCREEN": 0x4000, "KBD": 0x6000}
}

// violation the parser is not suppossed to know anything about InstructionsOutputData
func (parser *Parser) RemoveWhiteSpaceAndProcesSymbols(instructions asm.InstructionsInputData) asm.InstructionsOuputData {
	parser.removeWhiteSpace(instructions)
	parser.processSymbols()

	return parser.OutputInstructions
}

func (parser *Parser) processSymbols() {
	parser.addLabelsToSymbolTable()
	parser.translateSymbolsToAddress()
	// fmt.Println(parser.symbolTable)

	// for _, value := range parser.OutputInstructions.Instructions {
	// 	fmt.Println(value)
	// }
}

func (parser *Parser) translateSymbolsToAddress() {
	var instructions []string
	n := 16
	r, _ := regexp.Compile(`^@[a-zA-Z\s]{1}\w*`)
	for _, instruction := range parser.OutputInstructions.Instructions {
		if r.MatchString(instruction) {
			// symbol
			instruction = instruction[1:]

			address, ok := parser.symbolTable[instruction]
			if ok {
				instructions = append(instructions, fmt.Sprintf("@%d", address))
			} else {
				instructions = append(instructions, fmt.Sprintf("@%d", n))
				n++
			}

		} else {
			instructions = append(instructions, instruction)
		}
	}
	parser.OutputInstructions.Instructions = instructions
}

func (parser *Parser) addLabelsToSymbolTable() { // there's a possibility that there is more than one (label) kind of instruction
	var instructions []string
	r, _ := regexp.Compile(`^\([\w\d\s._$]*\)`)
	n := 0
	for address, instruction := range parser.InputInstructions.Instructions {
		// fmt.Printf("%d\t%s\n", address, instruction)
		if r.MatchString(instruction) {
			parser.symbolTable[instruction[1:len(instruction)-1]] = address - n
			n++
		} else {
			instructions = append(instructions, instruction)
		}
	}

	parser.OutputInstructions.Instructions = instructions
}

// This should be the work of the parser
func (parser *Parser) removeWhiteSpace(instructions asm.InstructionsInputData) {
	pInstructionStrings := instructions.Instructions
	var instructionsWithoutWhiteSpace []string

	r, _ := regexp.Compile(`^\/\/\s[\w\d\s",:=+/.([\])!#$%\^&\*_-\{\}\|:;"><,.~]*`) // remove all comments
	removeSpaceRe, _ := regexp.Compile(`^\s*$`)                                     // remove all white spaces
	// use ReplaceAllString
	for _, value := range pInstructionStrings {
		if !r.MatchString(value) && !removeSpaceRe.MatchString(value) {
			doesNotStartWithComments := strings.TrimSpace(value)
			instructionWithNoWhiteSpace := strings.Split(doesNotStartWithComments, "//")[0]
			instructionsWithoutWhiteSpace = append(instructionsWithoutWhiteSpace, instructionWithNoWhiteSpace)
		}
	}

	parser.InputInstructions = asm.InstructionsInputData{
		Instructions: instructionsWithoutWhiteSpace,
	}
}
