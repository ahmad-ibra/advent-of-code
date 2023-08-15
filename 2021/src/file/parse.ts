import { createReadStream } from 'fs'
import { createInterface } from 'readline'

export async function getLinesFromFile(filePath: string): Promise<string[]> {
    const lines: string[] = [];

    const readStream = createReadStream(filePath);
    const rl = createInterface({
        input: readStream,
    });

    for await (const line of rl) {
        lines.push(line);
    }

    return lines;
}
