import { getLinesFromFile } from '../file/parse.js'

const day = "dayXX"

export default async function dayXX() {
    console.log(`2021/${day}:`)

    const lines = await getLinesFromFile(`src/${day}/sample.txt`)
    //const lines = await getLinesFromFile(`src/${day}/input.txt`)

    for (const line of lines) {
    }
}
