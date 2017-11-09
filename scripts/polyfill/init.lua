dofile(ASP+"/scripts/polyfill/table.lua")
dofile(ASP+"/scripts/polyfill/world.lua")
dofile(ASP+"/scripts/polyfill/window.lua")

os.sleep = function(sec)
    local timr = os.time()
    repeat until os.time() >= timr + sec
end

Slot = { Amount = 0, Type = "air" }

---Internal function, identical to inv.inv()
---@return table<number, Slot>
function bot.inv()
    return json.decode(bot._inv())
end

---Returns all items in the bot inventory
---@return table<number, Slot>
function inv.inv()
    return bot.inv()
end

---Returns the current item in the cursor
---@return Slot
function inv.cursor()
    return json.decode(bot._cursor())
end

Item = { Amount = 0, Type = "air", X = 0.0, Y = 0.0, Z = 0.0 }

---Returns all items on the ground nearby
---@return table<nil, Item>
function bot.items_near()
    return json.decode(bot._items_near())
end

function round(num)
    return math.floor(num + .5)
end