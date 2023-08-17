import { getLinesFromFile } from '../file/parse.js'

const day = "day04"

interface board {
    b: number[][]
    marked: marker[]
    hasWon: boolean
}

interface marker {
    num: number,
    row: number,
    col: number
}

function sumUnselectedNums(b: board): number {
    let sum = 0

    let len = b.b[0].length
    for (let row = 0; row < len; row++) {
        for (let col = 0; col < len; col++) {
            sum += b.b[row][col]
        }
    }
    for (let m of b.marked) {
        sum -= m.num
    }

    return sum
}

function markNumber(n: number, b: board): void {
    for (let row = 0; row < b.b.length; row++) {
        for (let col = 0; col < b.b[0].length; col++) {
            if (n === b.b[row][col]) {
                let mark = { num: n, row: row, col: col }
                b.marked.push(mark)
                return
            }
        }
    }
}

function isWinner(b: board): boolean {
    const boardLen = b.b[0].length
    let r = new Array<number>(boardLen).fill(0)
    let c = new Array<number>(boardLen).fill(0)

    let marked = b.marked

    for (let m of marked) {
        r[m.row]++
        c[m.col]++

        if (r[m.row] === boardLen || c[m.col] === boardLen) {
            return true
        }
    }

    return false
}

export default async function day04() {
    console.log(`2021/${day}:`)

    //const lines = await getLinesFromFile(`src/${day}/sample.txt`)
    const lines = await getLinesFromFile(`src/${day}/input.txt`)

    let chosenNums: number[] = lines[0].split(",").map(n => parseInt(n))

    let boards: board[] = []
    let board: board = { b: [], marked: [], hasWon: false }
    for (let i = 2; i < lines.length; i++) {
        let line = lines[i]
        if (line.length == 0) {
            boards.push(board)
            board = { b: [], marked: [], hasWon: false }
            continue
        }

        let row = line.split(" ").filter(v => v.length > 0).map(n => parseInt(n))
        board.b.push(row)
    }
    boards.push(board)

    let winnerCount = 0
    for (let num of chosenNums) {
        for (let b of boards) {
            markNumber(num, b)
            if (isWinner(b) && !b.hasWon) {
                b.hasWon = true
                winnerCount++
                if (winnerCount === 1) {
                    const sum = sumUnselectedNums(b)
                    console.log(`Part1: Final score of chosen board is ${sum} * ${num} = ${sum * num}`)
                }
                if (winnerCount === boards.length) {
                    const sum = sumUnselectedNums(b)
                    console.log(`Part2: Final score of chosen board is ${sum} * ${num} = ${sum * num}`)
                }
            }
        }
    }
}
