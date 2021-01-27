
-- Get GPS
x,y,z = gps.locate()

-- Get orientation
ok = turtle.forward()
if ok then
    x2,y2,z2 = gps.locate()
    if x < x2 then
        print("facing west")
    else if x < x2 then
        print("facing east")
    end

    if z > z2 then
        print("facing north")
    else if x < x2 then
        print("facing south")
    end

end

-- Get capabilities

-- Get name

-- Sign-in
-- .. block for instructions
