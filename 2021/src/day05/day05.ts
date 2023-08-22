import { getLinesFromFile } from '../file/parse.js'

const day = "day05"

interface vent {
    x1: number,
    y1: number,
    x2: number,
    y2: number,
}

function notDiagonal(v: vent): boolean {
    return v.x1 === v.x2 || v.y1 === v.y2
}

function getVent(line: string): vent {
    let coordinates = line.split(" -> ")
    let start = coordinates[0].split(",")
    let end = coordinates[1].split(",")

    let v: vent = {
        x1: parseInt(start[0]),
        y1: parseInt(start[1]),
        x2: parseInt(end[0]),
        y2: parseInt(end[1])
    }

    return v
}

function markMap(map: number[][], v: vent) {
    if (v.x1 === v.x2) {
        // vertical line
        let minY = Math.min(v.y1, v.y2)
        let maxY = Math.max(v.y1, v.y2)

        let x = v.x1
        for (let i = minY; i <= maxY; i++) {
            let val = map[i][x]
            map[i][x] = ++val
        }
    }

    else if (v.y1 === v.y2) {
        // horizontal line
        let minX = Math.min(v.x1, v.x2)
        let maxX = Math.max(v.x1, v.x2)

        let y = v.y1
        for (let i = minX; i <= maxX; i++) {
            let val = map[y][i]
            map[y][i] = ++val
        }
    } else if (v.x1 < v.x2 && v.y1 < v.y2) {
        // down and right
        let y = v.y1
        for (let x = v.x1; x <= v.x2; x++) {
            let val = map[y][x]
            map[y][x] = ++val
            y++
        }
    } else if (v.x2 < v.x1 && v.y2 < v.y1) {
        // up and left
        let y = v.y2
        for (let x = v.x2; x <= v.x1; x++) {
            let val = map[y][x]
            map[y][x] = ++val
            y++
        }
    } else if (v.x1 < v.x2 && v.y2 < v.y1) {
        // up and right
        let y = v.y1
        for (let x = v.x1; x <= v.x2; x++) {
            let val = map[y][x]
            map[y][x] = ++val
            y--
        }
    } else {
        // down and left
        let y = v.y2
        for (let x = v.x2; x <= v.x1; x++) {
            let val = map[y][x]
            map[y][x] = ++val
            y--
        }
    }
}

function getGreaterThan(map: number[][], num: number): number {
    let count = 0
    for (let r = 0; r < map.length; r++) {
        for (let c = 0; c < map[0].length; c++) {
            if (map[r][c] > num) {
                count++
            }
        }
    }

    return count
}

export default async function day05() {
    console.log(`2021/${day}:`)

    //const lines = await getLinesFromFile(`src/${day}/sample.txt`)
    const lines = await getLinesFromFile(`src/${day}/input.txt`)

    const size = 1000

    let map: number[][] = [];

    for (let i = 0; i < size; i++) {
        map.push(new Array<number>(size).fill(0));
    }

    for (const line of lines) {
        let v: vent = getVent(line)
        if (notDiagonal(v)) {
            markMap(map, v)
        }
    }

    console.log(`Part 1: ${getGreaterThan(map, 1)} points where at least 2 lines overlap`)


    for (const line of lines) {
        let v: vent = getVent(line)
        if (!notDiagonal(v)) {
            markMap(map, v)
        }
    }

    console.log(`Part 2: ${getGreaterThan(map, 1)} points where at least 2 lines overlap`)
}
