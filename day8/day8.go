package day8

import (
	"regexp"
	"strconv"
)

type Operation int

const (
	acc Operation = 0
	jmp Operation = 1
	nop Operation = 2
)

type Instruction struct {
	operation Operation
	argument  int
}

var instructionRegex = regexp.MustCompile(`(acc|jmp|nop) ([+-]\d+)`)

func getInstructionFromRawInstruction(rawInstruction string) Instruction {
	matches := instructionRegex.FindStringSubmatch(rawInstruction)
	num, _ := strconv.Atoi(matches[2])

	var operation Operation

	switch matches[1] {
	case "acc":
		operation = acc
	case "jmp":
		operation = jmp
	case "nop":
		operation = nop
	default:
		panic("nooooo")
	}

	return Instruction{
		operation: operation,
		argument:  num,
	}
}

func getInstructions(rawInstructions []string) []Instruction {
	instructions := make([]Instruction, 0)

	for _, rawInstruction := range rawInstructions {
		instructions = append(instructions, getInstructionFromRawInstruction(rawInstruction))
	}

	return instructions
}

func runProgram(instructions []Instruction) (int, bool) {
	// NOTE: will simply stop on infinite loops and return current acc
	executed := map[int]struct{}{}
	accumulator := 0
	instructionNum := 0

	for {
		// Finished execution
		if instructionNum == len(instructions) {
			break
		}

		instruction := instructions[instructionNum]

		if _, ok := executed[instructionNum]; ok {
			return accumulator, false
		}

		executed[instructionNum] = struct{}{}

		switch instruction.operation {
		case acc:
			accumulator += instruction.argument
			instructionNum++

		case jmp:
			instructionNum += instruction.argument

		case nop:
			instructionNum++
		}
	}

	return accumulator, true
}

func Solve(rawInstructions []string) int {
	instructions := getInstructions(rawInstructions)
	accumulator, _ := runProgram(instructions)

	return accumulator
}

func SolvePart2(rawInstructions []string) int {
	instructions := getInstructions(rawInstructions)

	for i, instruction := range instructions {
		switch instruction.operation {
		case nop:
			modifiedProgram := make([]Instruction, len(instructions))
			copy(modifiedProgram, instructions)
			modifiedProgram[i].operation = jmp

			if accumulator, doesExit := runProgram(modifiedProgram); doesExit {
				return accumulator
			}

		case jmp:
			modifiedProgram := make([]Instruction, len(instructions))
			copy(modifiedProgram, instructions)
			modifiedProgram[i].operation = nop

			if accumulator, doesExit := runProgram(modifiedProgram); doesExit {
				return accumulator
			}
		}
	}

	panic("NO FIX FOUND!")
}
