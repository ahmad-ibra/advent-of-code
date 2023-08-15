import { getLinesFromFile } from '../file/parse.js'

export default async function day01() {
    console.log("2021/01:")

    const lines = await getLinesFromFile("src/day01/sample.txt")
    //const lines = await getLinesFromFile("src/day01/input.txt")

    for (const line of lines) {
        console.log(line)
    }
}
