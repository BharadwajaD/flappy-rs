
export class Pipe{
    /**
     * @param {number} capacity
     * @param {BufferWindow} bwindow
     */
    constructor(capacity, bwindow){
        this.capacity = capacity + 1
        this.head = 0
        this.tail = 0
        this.bwindow = bwindow

        this.queue = Array(this.capacity)
    }

    is_full(){
        return (this.tail + 1) % this.capacity == this.head
    }

    enqueue(item){
        if ((this.tail + 1) % this.capacity == this.head) {
            //throw full error
        }

        this.queue[this.tail] = item
        this.tail = (this.tail + 1) % this.capacity
    }

    dequeue(){
        if (this.tail == this.head) {
            //throw empty error
        }

        const item = this.queue[this.head]
        this.head = (this.head + 1) % this.capacity
        return item
    }

    Place(x, height){
        if (this.is_full()) {
            const p = this.dequeue()
            console.log("Removed pipe: ", p)
            this.bwindow.RemoveVline(p.x, p.height)
        }
        this.enqueue({x, height})
        this.bwindow.PlaceVline(x, height, 'p')
    }
}
