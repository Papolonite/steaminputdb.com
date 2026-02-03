/* eslint-disable no-console */

const ANSI = {
    reset: '\u001b[0m',
    gray: '\u001b[90m',
    blue: '\u001b[34m',
    green: '\u001b[32m',
    yellow: '\u001b[33m',
    red: '\u001b[31m',
    cyan: '\u001b[36m',
    white: '\u001b[37m'
};

const write = (levelColor: string, level: string, message: string, ...args: unknown[]) => {
    const now = new Date();
    const tzOffset = -now.getTimezoneOffset();
    const sign = tzOffset >= 0 ? '+' : '-';
    const tzHours = String(Math.floor(Math.abs(tzOffset) / 60)).padStart(2, '0');
    const tzMinutes = String(Math.abs(tzOffset) % 60).padStart(2, '0');
    const timestamp = `${now.toISOString().slice(0, -1)}${sign}${tzHours}:${tzMinutes}`;

    const parts = args.reduce((acc: string[], n, i) => {
        if (i % 2 === 0) {
            acc.push(`${ANSI.gray}${String(n)}${ANSI.reset}=`);
        } else {
            acc.push(`${String(n)} `);
        }
        return acc;
    },
    [
        `${ANSI.gray}${timestamp}${ANSI.reset} `,
        `${levelColor}${level.padStart(5)}${ANSI.reset} `,
        message,
        ' '
    ]
    );

    const line = parts.join('');

    switch (level) {
        case 'ERROR':
            console.error(line);
            break;
        case 'WARN':
            console.warn(line);
            break;
        case 'INFO':
            console.info(line);
            break;
        default:
            console.debug(line);
    }
};

export const log = {
    debug: (message: string, ...args: unknown[]) => write(ANSI.blue, 'DEBUG', message, ...args),
    info: (message: string, ...args: unknown[]) => write(ANSI.green, 'INFO', message, ...args),
    warn: (message: string, ...args: unknown[]) => write(ANSI.yellow, 'WARN', message, ...args),
    error: (message: string, ...args: unknown[]) => write(ANSI.red, 'ERROR', message, ...args)
};

export { ANSI };

