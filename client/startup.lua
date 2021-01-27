print("   ")
local x
local y
local z
local orientation

-- Get ID
print("> Getting computer ID")
print(os.getComputerID())
print("   ")

-- Get GPS
print("> Getting x, y, z")
x,y,z = gps.locate()
print(x .. ", " ..  y .. ", " .. z)
print("   ")
sleep(1)

-- Get orientation
print("> Getting orientation")
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
    else if x < x2 then
        print("facing south")
        orientation = "SOUTH"
    end
    turtle.back()
else
    print("could not move forward")
end
print("   ")

-- Get capabilities

-- Sign-in
-- .. block for instructions
