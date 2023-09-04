import { WebSocketServer } from 'ws';

export function listen(port: number) {
  const wss = new WebSocketServer({ port });

  wss.on('connection', (ws) => {
    ws.on('error', (e) => {
      console.error('Socket error', e);
    });

    ws.on('message', (data) => {
      console.log('Socket message', data);
    });
  });

  wss.on('error', (e) => {
    console.error('Server error', e);
  });
}
