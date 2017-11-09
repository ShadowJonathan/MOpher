Block = dofile(ASP .. "/polyfill/block.lua")

local Coords = { x = 0, y = 0, z = 0 }

function Coords:new(x, y, z)
    local o = { x, y, z }
    setmetatable(o, self)
    self.__index = self
    return o
end

---@type table<Coords, Block>
World = {}

---Gets a block from the world and returns it
---@param x int
---@param y int
---@param z int
---@return Block
function World:get(x, y, z)
    for key, value in pairs(World) do
        if key.x == x and key.y == y and key.z == z then
            return value:refresh()
        end
    end
    local b = World:_get(x, y, z)
    World[Coords:new(x, y, z)] = b
    return b
end

---Internal function to getting blocks
---@param x number
---@param y number
---@param z number
---@return Block
function World:_get(x, y, z)
    return Block:new(x, y, z, bot._block(x, y, z))
end