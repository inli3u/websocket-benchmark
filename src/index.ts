import { program, InvalidArgumentError } from 'commander';
import { listen } from './server';

const DEFAULT_PORT = 8080;

function validateInt(value: string) {
  const parsed = parseInt(value, 10);
  if (isNaN(parsed)) {
    throw new InvalidArgumentError('Not a valid int');
  }
  return parsed;
}

program
  .option('--port [number]', `Defaults to ${DEFAULT_PORT}`, validateInt);

program.parse();

const opts = program.opts();
const port = opts.port ?? DEFAULT_PORT;

listen(port);
console.log('Listening on port', port);

// output:
// - running on...
// - { connections: n, event: connect | disconnect | socket-error | server-error | message }
// 
// eventually only output aggregate stats
