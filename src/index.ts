import { EventEmitter } from "stream"

interface ISubscribe {
    subscribe(
        event: string,
        executor: (payload: any) => any | Promise <any>,
    ): this
}

type OrderEvent = 'order-created' | 'order-paid'

class OrderPublisherSubscriber implements ISubscribe {
    constructor(private readonly emitter: EventEmitter) {}

    subscribe(event: string, executor: (payload: any) => any): this {
        this.emitter.on(event, executor)
        return this
    }

    publish(event: OrderEvent | (string & {}), payload: any): this {
        this.emitter.emit(event, payload)
        return this
    }

}

class Order {
    constructor(
    private readonly _id: number,
    private readonly _name: string,
    private readonly _product: string,    
    ) {}

    get id() {
        return this._id
    }

    get name() {
        return this._name
    }

    get product() {
        return this._product
    }
}

const executor = () => {
    const emitter = new EventEmitter()
    const order = new OrderPublisherSubscriber(emitter)

    order.subscribe('order-created', (order: any) => {
        console.log('Order created', {...order, status: 'created' })
    })

}