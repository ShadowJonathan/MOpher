Block = { x = 0, y = 0, z = 0, t = "minecraft:air" }

---@param x int
---@param y int
---@param z int
function Block:new(x, y, z, t)
    local o = { x = x, y = y, z = z, t = t }
    setmetatable(o, self)
    self.__index = self
    return o
end

function Block:refresh()
    self.t = bot._block(self.x, self.y, self.z)
    return self
end

return Block