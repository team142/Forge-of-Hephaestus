# Forge-of-Hephaestus

The goal of this project is to use Turtles from ComputerCraft Tweaked as Kubernetes worker nodes to carry out service
roles captured as declarative configuration.

Turtles must be able to recover from de-syncs, reconnects and server restarts.


# Golden path

## Startup, connect and locations
- Server - websocket server
- Startup - connect to websocket
- Startup - send ID, label, x,y,z,orientation
- Server - file for highway
- Server - endpoint for add highway
- Client - submit highway
- Server - file for garage
- Server - endpoint for garage
- Client - submit garage

## post-start
- Server - on connect turtle mark turtle as LOST
- Server - if not on garage tile, make drop-off, refuel and park (new Goal)
- Server - instantiate goal
- Server - allocate Goal to Turtle
- Server - goal, create tasks
- Server - goal, task, create actions
- Server - send file of actions
- Client - write file to disk
- Client - run file
- Client - send - file done

## Post-park
- Server - if parking, mark turtle as AVAILABLE_FOR_WORK

## Scheduler
- Server - load all Orchestrations
- Server - Give Orchs to Scheduler
- Server - for each Orch and AVAILABLE turtle, apply to turtle
- Server - generate job, tasks and action first task





## Server

To do

- Cron manager ***
- transfer lib.lua file to turtle
- transfer startup.lua file to turtle
- keep track of which turtle has which hash of lib.lua
- keep track of which turtle has which hash of startup.lua
- Action planner per Task (SEE BELOW)
- Websocket server (SEE BELOW)

## Meta

- Notify server of HW tile
- Server: save HW tile

## Server - ws

- ws message handler ***
- ws send message handler ***
- Server: is this turtle parked in Garage ***

## Server - Tasks

- Park ***
- Leave Garage ***
- Travel to X ***
- Collect wheat seeds ***
- Wheat cut and plant ***
- Kelp cut
- Kelp Collect
- Drop off ***
- Refuel ***

## Turtle Startup

To do

- Lib: ws client ***
- Lib: correlation ID and return on OK ***
- Comms: sign in (label, ID, coords, orientation) ***
- Comms: tell server task is Done ***
- Lib: patiently move forward ***
- Lib: aggressively move forward ***
- Lib: turn X ***
- Lib: suck ***
- Lib: drop ***
- Lib: up ***
- Lib: down ***
- Lib: refuel ***
- Lib: select ***
- Lib: read content of X ***
- Lib: get remote file ***
- Lib: reboot ***

