function marchForward()
    ok = turtle.forward()
    while (not ok) do
        turtle.dig()
        ok = turtle.forward()
    end
end

turtle.forward()
for i = 1, 7 do
    turtle.down()
end

turnLeft = false
for r = 1, 50 do
    if turnLeft then
        turtle.turnLeft()
    else
        turtle.turnRight()
    end
    turnLeft = not turnLeft
    for i = 1, 14 do
        marchForward()
    end
    if turnLeft then
        turtle.turnLeft()
    else
        turtle.turnRight()
    end
    marchForward()

end

