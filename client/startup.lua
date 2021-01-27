local function isempty(s)
  return s == nil or s == ''
end

print("-----------")
local x
local y
local z
local orientation = ""
local ID = turtle.getComputerID()
local label = turtle.getComputerLabel()

-- Get ID
print(">> Getting computer ID")
print(os.getComputerID())
print("   ")

-- Get GPS
print(">> Getting x, y, z")
x,y,z = gps.locate()
if isempty(x) then
    print(x .. " is nil. Shutting down")
    return
end
-- TODO: if anything here is null, the satellite is DOWN

print(x .. ", " ..  y .. ", " .. z)
print("   ")
sleep(1)

-- Get orientation
print(">> Getting orientation")
while orientation == "" do
    ok = turtle.forward()
    if ok then
        x2,y2,z2 = gps.locate()
        if x > x2 then
            print("facing west")
            orientation = "WEST"
        elseif x < x2 then
            print("facing east")
            orientation = "EAST"
        end

        if z > z2 then
            print("facing north")
            orientation = "NORTH"
        elseif z < z2 then
            print("facing south")
            orientation = "SOUTH"
        end
        turtle.back()
    else
        print("could not move forward")
    end
    if orientation == "" then
        ok = turtle.back()
        if ok then
            x2,y2,z2 = gps.locate()
            if x < x2 then
                print("facing west")
                orientation = "WEST"
            elseif x > x2 then
                print("facing east")
                orientation = "EAST"
            end

            if z < z2 then
                print("facing north")
                orientation = "NORTH"
            elseif z > z2 then
                print("facing south")
                orientation = "SOUTH"
            end
            turtle.forward()
        else
            print("could not move back")
        end
    end

    if orientation == "" then
        print("sleeping ... then trying again")
        sleep(2)
    end
end
print("   ")

-- Get capabilities

-- Sign-in
-- .. block for instructions
