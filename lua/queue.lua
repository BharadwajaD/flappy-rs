local Queue = {}
Queue.__index = Queue

function Queue:new(capacity)
    self.capacity = capacity
    self.head = 0
    self.tail = 0
    self.content = {}

    for _ = 1, capacity, 1 do
        table.insert(self.content, nil)
    end

    return self
end

function Queue:enqueue(item)
    local deq_item
    if (self.tail + 1) % self.capacity == self.head then
        deq_item =  self:dequeue()
    end
    self.content[self.tail + 1] = item
    self.tail = (self.tail + 1) % self.capacity
    return deq_item
end

function Queue:dequeue()
    assert(self.head ~= self.tail, "Queue empty !")
    local item = self.content[self.head + 1] -- +1 bcoz of lua 1 indexing
    self.content[self.head + 1] = nil -- +1 bcoz of lua 1 indexing
    self.head = (self.head + 1)  % self.capacity

    return item
end

return Queue
