import { getLinesFromFile } from '../file/parse.js'

const day = "day03"

export default async function day03() {
    console.log(`2021/${day}:`)

    //const lines = await getLinesFromFile(`src/${day}/sample.txt`)
    const lines = await getLinesFromFile(`src/${day}/input.txt`)

    const lineLen = lines[0].length
    const lineCount = lines.length

    let onesCount = new Array<number>(lineLen).fill(0)

    let O2Rating = {}
    let CO2Rating = {}

    for (const line of lines) {
        O2Rating[line] = true
        CO2Rating[line] = true
        for (let i = 0; i < line.length; i++) {
            let c = parseInt(line[i])
            onesCount[i] = onesCount[i] + c
        }
    }

    let gammaRate = ""
    let epsilonRate = ""
    for (let i = 0; i < lineLen; i++) {
        if (onesCount[i] >= lineCount / 2) {
            gammaRate += "1"
            epsilonRate += "0"
        } else {
            gammaRate += "0"
            epsilonRate += "1"
        }
    }

    let grI = parseInt(gammaRate, 2)
    let erI = parseInt(epsilonRate, 2)
    console.log(`Part1: gammaRate ${grI} * episolonRate ${erI} = ${grI * erI}`)

    let ogrRemaining = lineCount
    for (let i = 0; i < lineLen; i++) {
        let countZero = 0
        let countOne = 0
        for (const k in O2Rating) {
            if (O2Rating[k] === true) {
                if (k[i] == "0") {
                    countZero++
                } else {
                    countOne++
                }
            }
        }
        let val = countOne >= countZero ? "1" : "0"
        const keys = Object.keys(O2Rating);
        for (const k of keys) {
            if (O2Rating[k] === true) {
                if (ogrRemaining === 1) {
                    break
                }

                if (k[i] != val) {
                    O2Rating[k] = false
                    ogrRemaining--
                }
            }
        }
    }
    let ogr: number
    for (const k in O2Rating) {
        if (O2Rating[k] === true) {
            ogr = parseInt(k, 2);
        }
    }

    let csrRemaining = lineCount
    for (let i = 0; i < lineLen; i++) {
        let countZero = 0
        let countOne = 0
        for (const k in CO2Rating) {
            if (CO2Rating[k] === true) {
                if (k[i] == "0") {
                    countZero++
                } else {
                    countOne++
                }
            }
        }
        let val = countZero <= countOne ? "0" : "1"
        const keys = Object.keys(CO2Rating);
        for (const k of keys) {
            if (CO2Rating[k] === true) {
                if (csrRemaining === 1) {
                    break
                }

                if (k[i] != val) {
                    CO2Rating[k] = false
                    csrRemaining--
                }
            }
        }
    }
    let csr: number
    for (const k in CO2Rating) {
        if (CO2Rating[k] === true) {
            csr = parseInt(k, 2);
        }
    }

    console.log(`Part2: O2Rating ${ogr} * CO2Rating ${csr} = ${ogr * csr}`)
}
