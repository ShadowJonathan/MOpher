window = {}

---Returns all items in the window inventory
---@return table<number, Slot>
function window.inv()
    return json.decode(_window._inv())
end

---Transfer an item from an inventory slot to a window slot
---@param invSlot number
---@param windowSlot number
---@param swap boolean returns an error when set to false and there is already an item in the window slot
---@return string the Error
function window.transfer_to(invSlot, windowSlot, swap)
    return _window._TT(invSlot, windowSlot, swap)
end

---Transfer an item from a window slot to an inventory slot
---@param windowSlot number
---@param invSlot number
---@param swap boolean returns an error when set to false and there is already an item in the inventory slot
---@return string the Error
function window.transfer_from(windowSlot, invSlot, swap)
    return _window._TF(invSlot, windowSlot, swap)
end

---Transfer an item from a slot to another slot, no conversion
---@param from number
---@param to number
---@param swap boolean returns an error when set to false and there is already an item in the "to" slot
---@return string the Error
function window.transfer_complex(from, to, swap)
    return _window._TC(from, to, swap)
end
