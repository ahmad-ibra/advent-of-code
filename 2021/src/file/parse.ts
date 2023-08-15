import * as fs from 'fs'
import * as readline from 'readline'

export async function getLinesFromFile(filePath: string): Promise<string[]> {
    const lines: string[] = [];

    const readStream = fs.createReadStream(filePath);
    const rl = readline.createInterface({
        input: readStream,
    });

    for await (const line of rl) {
        lines.push(line);
    }

    return lines;
}
