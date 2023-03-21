package encoder

import (
	asm "assembler/boundaries"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type AInstruction struct {
	Instruction string
}

type CInstruction struct {
	Dest string
	Comp string
	Jump string
}

type Encoder struct {
	CompMap map[string]string
	DestMap map[string]string
	JumpMap map[string]string
}

func MakeEncoder() Encoder {
	// build maps for reference
	encoder := Encoder{}
	encoder.CompMap = map[string]string{"0": "0101010", "1": "0111111",
		"-1": "0111010", "D": "0001100",
		"A": "0110000", "!D": "0001101",
		"!A": "0110001", "-D": "0001111",
		"-A": "0110011", "D+1": "0011111",
		"A+1": "0110111", "D-1": "0001110",
		"A-1": "0110010", "D+A": "0000010",
		"D-A": "0010011", "A-D": "0000111",
		"D&A": "0000000", "D|A": "0010101",
		"M": "1110000", "!M": "1110001",
		"-M": "1110011", "M+1": "1110111",
		"M-1": "1110010", "D+M": "1000010",
		"D-M": "1010011", "M-D": "1000111",
		"D&M": "1000000", "D|M": "1010101"}

	encoder.DestMap = map[string]string{"": "000", "M": "001", "D": "010", "MD": "011",
		"A": "100", "AM": "101", "AD": "110", "AMD": "111"}

	encoder.JumpMap = map[string]string{"": "000", "JGT": "001", "JEQ": "010", "JGE": "011",
		"JLT": "100", "JNE": "101", "JLE": "110", "JMP": "111"}
	return encoder
}

// violation the encoder is not suppossed to know anything about InstructionsOutputData
func (encoder *Encoder) EncodeToBinary(instructionsWithoutSymbols asm.InstructionsOuputData) []string {
	processedInstructions := instructionsWithoutSymbols.Instructions
	instructionsWithTypes := encoder.identifyInstructions(processedInstructions)
	for _, instruction := range instructionsWithTypes {
		// fmt.Printf("%s %T \t", instruction, instruction)
		switch instruction.(type) {
		case AInstruction:
			fmt.Printf("%s\n", encoder.encodeAInstruction(instruction.(AInstruction)))
		case CInstruction:
			fmt.Printf("%s\n", encoder.encodeCInstruction(instruction.(CInstruction)))
		}

	}

	// 		use tables to map string value to binary value
	return processedInstructions
}

func (encoder *Encoder) encodeAInstruction(instruction AInstruction) string {
	aInstruction, err := strconv.Atoi(instruction.Instruction)
	if err != nil {
		log.Fatal("AInstruction is not valid")
	}
	return fmt.Sprintf("%016b", aInstruction)

}

func (encoder *Encoder) encodeCInstruction(instruction CInstruction) string {
	// aCheck := strings.Contains(instruction.Comp, "M")
	// var a string
	// if aCheck {
	// 	a = "1"
	// } else {
	// 	a = "0"
	// }
	compInstruction := encoder.CompMap[instruction.Comp]
	destInstruction := encoder.DestMap[instruction.Dest]
	jumpInstruction := encoder.JumpMap[instruction.Jump]
	// fmt.Printf("%s %s\t", instruction.Comp, compInstruction)
	return fmt.Sprintf("111%s%s%s", compInstruction, destInstruction, jumpInstruction)
}

func (encoder *Encoder) identifyInstructions(processedInstructions []string) []interface{} {
	instructions := make([]interface{}, 0)
	r, _ := regexp.Compile(`^@`)
	for _, instruction := range processedInstructions {
		if r.MatchString(instruction) {
			// A instruction
			instructions = append(instructions, AInstruction{
				Instruction: instruction[1:],
			})
		} else {
			var cInstruction CInstruction

			components := strings.Split(instruction, "=")
			if len(components) == 2 { // D=M exists
				cInstruction.Dest = strings.TrimSpace(components[0])
				remainingInstructions := strings.TrimSpace(components[1])
				remainingInstructionsComponents := strings.Split(remainingInstructions, ";")
				if len(remainingInstructionsComponents) == 2 { // comp and jump exist
					cInstruction.Comp = strings.TrimSpace(remainingInstructionsComponents[0])
					cInstruction.Jump = strings.TrimSpace(remainingInstructionsComponents[1])
				} else if len(remainingInstructionsComponents) == 1 { // only comp exist
					cInstruction.Comp = strings.TrimSpace(remainingInstructionsComponents[0])
				}
			} else if len(components) == 1 { // = does not exist
				components = strings.Split(instruction, ";")
				if len(components) == 2 { // comp and jump exist
					cInstruction.Comp = strings.TrimSpace(components[0])
					cInstruction.Jump = strings.TrimSpace(components[1])
				} else { // what happens in the case where there is no = nor ;
					cInstruction.Comp = strings.TrimSpace(components[0])
				}
			}

			instructions = append(instructions, cInstruction)
			// split by ;
		}
	}
	return instructions
}
