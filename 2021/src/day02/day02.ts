import { getLinesFromFile } from '../file/parse.js'

const day = "day02"

interface position {
    horizontal: number
    depth: number
}

interface positionPart2 {
    horizontal: number
    depth: number
    aim: number
}

function move(instruction: string, pos: position): position {
    let parts = instruction.split(" ")
    let dir = parts[0]
    let dist: number = parseInt(parts[1])

    switch (dir) {
        case "forward":
            pos.horizontal += dist
            break
        case "down":
            pos.depth += dist
            break
        case "up":
            pos.depth -= dist
            break
        default:
            console.log(`ERROR, unexpected direction ${dir}`)
    }
    return pos
}

function movePart2(instruction: string, pos: positionPart2): positionPart2 {
    let parts = instruction.split(" ")
    let dir = parts[0]
    let dist: number = parseInt(parts[1])

    switch (dir) {
        case "forward":
            pos.horizontal += dist
            pos.depth += pos.aim * dist
            break
        case "down":
            pos.aim += dist
            break
        case "up":
            pos.aim -= dist
            break
        default:
            console.log(`ERROR, unexpected direction ${dir}`)
    }
    return pos
}

export default async function day02() {
    console.log(`2021/${day}:`)

    //const lines = await getLinesFromFile(`src/${day}/sample.txt`)
    const lines = await getLinesFromFile(`src/${day}/input.txt`)

    let pos: position = {
        horizontal: 0,
        depth: 0
    }

    for (const line of lines) {
        pos = move(line, pos)
    }

    console.log(`Part1: horizontal (${pos.horizontal}) * depth (${pos.depth}) = ${pos.horizontal * pos.depth}`)

    let pos2: positionPart2 = {
        horizontal: 0,
        depth: 0,
        aim: 0
    }

    for (const line of lines) {
        pos = movePart2(line, pos2)
    }

    console.log(`Part2: horizontal (${pos.horizontal}) * depth (${pos.depth}) = ${pos.horizontal * pos.depth}`)
}
