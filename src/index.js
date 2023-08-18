"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const stream_1 = require("stream");
class OrderPublisherSubscriber {
    constructor(emitter) {
        this.emitter = emitter;
    }
    subscribe(event, executor) {
        this.emitter.on(event, executor);
        return this;
    }
    publish(event, payload) {
        this.emitter.emit(event, payload);
        return this;
    }
}
class Order {
    constructor(_id, _name, _product) {
        this._id = _id;
        this._name = _name;
        this._product = _product;
    }
    get id() {
        return this._id;
    }
    get name() {
        return this._name;
    }
    get product() {
        return this._product;
    }
}
// Adding executor
const executor = () => {
    const emitter = new stream_1.EventEmitter();
    const order = new OrderPublisherSubscriber(emitter);
    order.subscribe('order-created', (order) => {
        console.log('Order created', Object.assign(Object.assign({}, order), { status: 'created' }));
        order.publish('order-paid', order);
    })
        .subscribe('order-paid', (order) => {
        console.log('Order paid', Object.assign(Object.assign({}, order), { status: 'paid' }));
        order.publish('order-finished', order);
    })
        .subscribe('order-finished', (order) => {
        console.log('Order finished', Object.assign(Object.assign({}, order), { status: 'finished' }));
    });
};
