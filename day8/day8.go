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

func doesProgramHaveInfiniteLoop(instructions []Instruction) bool {
	executed := map[int]struct{}{}
	instructionNum := 0

	for {
		// Finished execution
		if instructionNum == len(instructions) {
			return false
		}

		instruction := instructions[instructionNum]

		if _, ok := executed[instructionNum]; ok {
			return true
		}

		executed[instructionNum] = struct{}{}

		switch instruction.operation {
		case acc:
			instructionNum++

		case jmp:
			instructionNum += instruction.argument

		case nop:
			instructionNum++
		}
	}
}

func runProgram(instructions []Instruction) int {
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
			return accumulator
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

	return accumulator
}

func Solve(rawInstructions []string) int {
	instructions := getInstructions(rawInstructions)

	return runProgram(instructions)
}

func SolvePart2(rawInstructions []string) int {
	instructions := getInstructions(rawInstructions)

	for i, instruction := range instructions {
		switch instruction.operation {
		case nop:
			modifiedProgram := make([]Instruction, len(instructions))
			copy(modifiedProgram, instructions)
			modifiedProgram[i].operation = jmp

			if !doesProgramHaveInfiniteLoop(modifiedProgram) {
				return runProgram(modifiedProgram)
			}

		case jmp:
			modifiedProgram := make([]Instruction, len(instructions))
			copy(modifiedProgram, instructions)
			modifiedProgram[i].operation = nop

			if !doesProgramHaveInfiniteLoop(modifiedProgram) {
				return runProgram(modifiedProgram)
			}
		}
	}

	panic("NO FIX FOUND!")
}
