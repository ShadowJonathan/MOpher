function _is_poly(it)
    error(it + " is polyfilled and not properly initialised!")
end

if bot == nil then
    print("bot not found")

    ---Blocks, navigates to the given coordinates,
    ---and returns a string error when no path is found, or the block is not valid
    ---@param x number
    ---@param y number
    ---@param z number
    ---@return string|nil
    function bot.nav(x, y, z)
        _is_poly("bot.nav")
    end

    ---Use bot.block, internal function
    ---@param x number
    ---@param y number
    ---@param z number
    ---@return string
    function bot._block(x, y, z)
        _is_poly("bot._block")
    end

    ---Logs to web-interface output, can only accept numbers, booleans or strings
    ---(as `tostring` does not support everything)
    ---@param ... string[]
    function bot.log(...)
        _is_poly("bot.log")
    end

    ---Returns X,Y,Z of current position of the bot
    ---@return number,number,number
    function bot.pos()
        _is_poly("bot.pos")
    end
end

if inv == nil then
    print("inv not found")

    ---Pick up item in inventory slot `at` and place it in the cursor, 
    ---returns false if there is an item in the cursor already
    ---@param at number
    ---@return boolean
    function inv.pickup(at)
        _is_poly("inv.pickup")
    end

    ---Drop item in inventory slot `at` from the cursor, 
    ---returns false if there is already an item in the slot
    ---@param at number
    ---@return boolean
    function inv.drop(at)
        _is_poly("inv.drop")
    end

    ---Swaps the items from slot `at` and cursor
    ---@param at number
    function inv.swap(at)
        _is_poly("inv.swap")
    end

    ---Simple true/false for if the cursor has an item in it.
    ---@return boolean
    function inv.cursorIsHolding()
        _is_poly("inv.cursorIsHolding")
    end
end

if json == nil then
    print("json not found")

    ---Encode everything but functions to string
    ---@param any any
    function json.encode(any)
        _is_poly("json.encode")
    end

    ---Decode json string to value
    ---@param value string
    function json.decode(value)
        _is_poly("json.decode")
    end
end

if _window == nil then
    print("_window not found")

    ---Returns the type of the current window:
    --- -1: Player Inventory
    --- 0: Container window
    --- 1: Chest Window
    --- 2: Crafting Table Window
    --- 3: Furnace Window
    --- 4: Dispenser Window
    --- 5: Enchanting Table Window
    --- 6: Brewing Stand Window
    --- 7: Villager Window
    --- 8: Beacon Window
    --- 9: Anvil Window
    --- 10: Hopper Window
    --- 11: Dropper Window
    --- 12: Shulker Box Window
    --- 13: Horse Window
    ---@return number
    function _window._current_type()
        _is_poly("_window._current_type")
    end

    ---Returns all items in the current window inventory
    ---@return string json encoded string
    function _window._inv()

    end

    ---[INTERNAL] Transfer an item from an inventory slot to a window slot
    ---@param invSlot number
    ---@param windowSlot number
    ---@param swap boolean returns an error when set to false and there is already an item in the window slot
    ---@return string the Error
    function _window._TT(invSlot, windowSlot, swap)
        _is_poly("_window._TT")
    end

    ---[INTERNAL] Transfer an item from a window slot to an inventory slot
    ---@param windowSlot number
    ---@param invSlot number
    ---@param swap boolean returns an error when set to false and there is already an item in the inventory slot
    ---@return string the Error
    function _window._TF(windowSlot, invSlot, swap)
        _is_poly("_window._TF")
    end

    ---[INTERNAL] Transfer an item from a slot to another slot, no conversion
    ---@param from number
    ---@param to number
    ---@param swap boolean returns an error when set to false and there is already an item in the "to" slot
    ---@return string the Error
    function _window._TC(from, to, swap)
        _is_poly("_window._TC")
    end
end

dofile(ASP .. "/polyfill/init.lua")