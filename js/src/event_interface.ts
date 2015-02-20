module Signaling {

    interface EventData {
        data: any;
    }

    interface Callback {
        fn: Function;
        once: number;
    }
    export class EventInterface {
        callbacks: {name: Callback[]};

        constructor() {
        }

        on(name: string, callback: (evt: EventData) => void) {
            this.attach(name, {
                fn: callback,
                once: 0
            });
        }

        once(name: string, callback: (evt: EventData) => void) {
            this.attach(name, {
                fn: callback,
                once: 1
            });
        }

        off(name: string, callback?: Function) {
            if ( ! (name in this.callbacks) ) {
                return;
            } else if ( ! callback ) {
                delete this.callbacks[name];
                return;
            }

            var callbacks: Callback[] = this.callbacks[name];
            var size: number = callbacks.length;
            var i: number = 0;

            for ( ; i < size; ++i ) {
                if ( callbacks[i].fn === callback ) {
                    this.callbacks[name].splice(i--, 1);
                }
            }
        }

        trigger(name: string, data?: any) {
            if ( ! (name in this.callbacks) ) {
                return;
            }

            var callbacks: Callback[] = this.callbacks[name];
            var size: number = callbacks.length;
            var i: number = 0;

            for ( ; i < size; ++i ) {
                callbacks[i].fn({data: data});
                if ( callbacks[i].once > 0 ) {
                    this.off(name, callbacks[i].fn);
                }
            }
        }

        private attach(name: string, cb: Callback) {
            if ( ! (name in this.callbacks) ) {
                this.callbacks[name] = [];
            }
            this.callbacks[name].push(cb);
        }


    }
}
