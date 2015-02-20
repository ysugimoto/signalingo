///<reference path="event_interface.ts" />
///<reference path="socket.ts" />
module Signaling {
    
    class Connector extends EventInterface {
        socket: SignalSocket;

        constructor(url: string) {
            super();
            this.socket = new SignalSocket(url, this);
            this.socket.connect();
        }
    }

        
    export function connect(url: string): Connector {
        return new Connector(url);
    }

}


