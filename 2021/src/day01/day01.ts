import { parse } from 'path'
import { getLinesFromFile } from '../file/parse.js'

export default async function day01() {
    console.log("2021/01:")

    //const lines = await getLinesFromFile("src/day01/sample.txt")
    const lines = await getLinesFromFile("src/day01/input.txt")

    let increaseCount: number = -1
    let prev: number = -1
    for (const line of lines) {
        let cur = parseInt(line)
        if (cur > prev) {
            increaseCount++
        }
        prev = cur
    }

    console.log(`There are ${increaseCount} measurements larger than the previous measurement`)

    increaseCount = 0
    let prevCount: number = parseInt(lines[0]) + parseInt(lines[1]) + parseInt(lines[2])

    for (let i = 3; i < lines.length; i++) {
        let curCount = prevCount + parseInt(lines[i]) - parseInt(lines[i - 3])
        if (curCount > prevCount) {
            increaseCount++
        }
        prevCount = curCount
    }

    console.log(`There are ${increaseCount} sums larger than the previous sum.`)
}
