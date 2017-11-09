if ASP == nil then
    error("ASP NOT SET")
end


bot = require("bot")
json = require("json")
inv = require("inv")
_window = require("_window")
dofile(ASP+"/polyfill/polyfill.lua")

bot.log("LOADED")

function main()
    os.sleep(5)

end

function collect()
    for _, item in pairs(bot.items_near()) do
        bot.nav(item.X, item.Y, item.Z)
    end
end

function t(it)
    bot.log(json.encode(table.load(table.save(it))))
end

function j(it)
    bot.log(json.encode(it))
end

function l(it)
    bot.log(it)
end

function dig_trees()

end

function near_trees()
    local trees = {}
    local posx, posy, posz = bot.pos()
    for x = -20, 20 do
        for z = -20, 20 do
            for y = 0, 2 do
                local b = World:get(round(posx + x), round(posy + y), round(posz + z))
                if b.t:find("log") then
                    bot.log("FOUND TREE AT", b.x, b.y, b.z)
                end
            end
        end
    end
end