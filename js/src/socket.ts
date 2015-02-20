///<reference path="event_interface.ts" />

module Signaling {

    enum WebSocketEvents {
        open,
        close,
        error,
        message
    }

    export class SignalSocket {

        ws: WebSocket;

        constructor(private url: string, private handler: EventInterface) {
        }

        connect() {
            this.ws = new WebSocket(this.url);

            for ( var i : WebSocketEvents = WebSocketEvents.open; i < WebSocketEvents.message; ++i ) {
                this.ws.addEventListener(WebSocketEvents[i], (evt: MessageEvent) => {
                    this.handleEvent(evt);
                });
            }
        }

        handleEvent(evt: MessageEvent) {
            switch (evt.type) {
                case WebSocketEvents[WebSocketEvents.open]:
                   console.log("OPENED");
                   this.handler.trigger(WebSocketEvents[WebSocketEvents.open]);
                   break;
                case WebSocketEvents[WebSocketEvents.close]:
                   console.log("CLOSED");
                   this.handler.trigger(WebSocketEvents[WebSocketEvents.close]);
                   break;
                case WebSocketEvents[WebSocketEvents.error]:
                   console.log("ERROR");
                   this.handler.trigger(WebSocketEvents[WebSocketEvents.error]);
                   break;
                case WebSocketEvents[WebSocketEvents.message]:
                   console.log("MESSAGE");
                   this.handler.trigger(WebSocketEvents[WebSocketEvents.message], {data: evt.data});
                   break;
            }
        }

    }
}
