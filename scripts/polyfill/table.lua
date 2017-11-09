--[[
   Save Table to File
   Load Table from File
   v 1.0

   Lua 5.2 compatible

   Only Saves Tables, Numbers and Strings
   Insides Table References are saved
   Does not save Userdata, Metatables, Functions and indices of these
   ----------------------------------------------------
   table.save( table , filename )

   on failure: returns an error msg

   ----------------------------------------------------
   table.load( filename or stringtable )

   Loads a table that has been saved via the table.save function

   on success: returns a previously saved table
   on failure: returns as second argument an error msg
   ----------------------------------------------------

   Licensed under the same terms as Lua itself.
]]--
do
    -- declare local variables
    --// exportstring( string )
    --// returns a "Lua" portable version of the string
    local function exportstring( s )
        return string.format("%q", s)
    end

    --// The Save Function
    function table.save(tbl)
        local charS, charE = "   ", " "
        local total = ""

        -- initiate variables for save procedure
        local tables, lookup = { tbl }, { [tbl] = 1 }
        total = total .. ( "return {" .. charE )

        for _, t in ipairs( tables ) do
            total = total .. ( "{" .. charE )
            local thandled = {}

            for i, v in ipairs( t ) do
                thandled[i] = true
                local stype = type( v )
                -- only handle value
                if stype == "table" then
                    if not lookup[v] then
                        table.insert( tables, v )
                        lookup[v] = #tables
                    end
                    total = total .. ( charS .. "{" .. lookup[v] .. "}," .. charE )
                elseif stype == "string" then
                    total = total .. (  charS .. exportstring( v ) .. "," .. charE )
                elseif stype == "number" then
                    total = total .. (  charS .. tostring( v ) .. "," .. charE )
                end
            end

            for i, v in pairs( t ) do
                -- escape handled values
                if (not thandled[i]) then

                    local str = ""
                    local stype = type( i )
                    -- handle index
                    if stype == "table" then
                        if not lookup[i] then
                            table.insert( tables, i )
                            lookup[i] = #tables
                        end
                        str = charS .. "[{" .. lookup[i] .. "}]="
                    elseif stype == "string" then
                        str = charS .. "[" .. exportstring( i ) .. "]="
                    elseif stype == "number" then
                        str = charS .. "[" .. tostring( i ) .. "]="
                    end

                    if str ~= "" then
                        stype = type( v )
                        -- handle value
                        if stype == "table" then
                            if not lookup[v] then
                                table.insert( tables, v )
                                lookup[v] = #tables
                            end
                            total = total .. ( str .. "{" .. lookup[v] .. "}," .. charE )
                        elseif stype == "string" then
                            total = total .. ( str .. exportstring( v ) .. "," .. charE )
                        elseif stype == "number" then
                            total = total .. ( str .. tostring( v ) .. "," .. charE )
                        end
                    end
                end
            end
            total = total .. ( "}," .. charE )
        end
        total = total .. ( "}" )
        return total
    end

    --// The Load Function
    function table.load(tstring)
        return loadstring(tstring)()
    end
    -- close do
end
