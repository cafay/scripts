//==== This script is created by RockClubKASHMIR ====

//--- WARNING!!! This script can work ONLY if you are Discoverer! ---
fromSystem = 100 // Your can change this value as you want
toSystem = 200 // Your can change this value as you want
Pnbr = 1  // When Rnbr = 1, the script will search only debris for minimum 2 Pathfinders. You can change this value as you want
times = 1 // if times = 1, the script will full scan 2 times the galaxy, from system, to system you want. Change this value as you wish
//----
cycle = 0
curSystem = fromSystem
origin = nil
flts = 0
nbr = 0
err = nil
totalSlots = GetSlots().Total - GetFleetSlotsReserved()
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Pathfinder > flts {
        flts = ships.Pathfinder
        origin = celestial // Your Planet(or Moon), with more Pathfinders
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    for system = curSystem; system <= toSystem; system++ {
        systemInfo, _ = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        Dtarget, _ = ParseCoord(origin.GetCoordinate().Galaxy+":"+system+":"+16)
        Sleep(Random(500, 1500)) // for avoid ban
        slots = GetSlots().InUse
        if err != nil {slots = totalSlots}
        if slots < totalSlots {
            if Dtarget != nil {
                Print("Checking "+Dtarget)
                if systemInfo.ExpeditionDebris.PathfindersNeeded > Pnbr { 
                    ships, _ = origin.GetShips()
                    if systemInfo.ExpeditionDebris.Metal == 0 && systemInfo.ExpeditionDebris.Crystal > 0 {Print("Found Crystal: "+systemInfo.ExpeditionDebris.Crystal)}
                    if systemInfo.ExpeditionDebris.Metal > 0 && systemInfo.ExpeditionDebris.Crystal == 0 {Print("Found Metal: "+systemInfo.ExpeditionDebris.Metal)}
                    if systemInfo.ExpeditionDebris.Metal > 0 && systemInfo.ExpeditionDebris.Crystal > 0 {Print("Found Metal: "+systemInfo.ExpeditionDebris.Metal+" and Crystal: "+systemInfo.ExpeditionDebris.Crystal)}
                    f = NewFleet()
                    f.SetOrigin(origin)
                    f.SetDestination(Dtarget)
                    f.SetSpeed(HUNDRED_PERCENT)
                    f.SetMission(RECYCLEDEBRISFIELD)
                    nbr = systemInfo.ExpeditionDebris.PathfindersNeeded
                    if systemInfo.ExpeditionDebris.PathfindersNeeded > ships.Pathfinder {nbr = ships.Pathfinder}
                    f.AddShips(PATHFINDER, nbr)
                    a, err = f.SendNow()
                    if err == nil {
                        if nbr < systemInfo.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                        if nbr > 1 {
                            Print(nbr+" Pathfinders are sended successfully!")
                        } else {Print(nbr+" Pathfinder is sended successfully!")}
                    } else {
                        if nbr > 1 {
                            Print("The Pathfinders are NOT sended! "+err)
                            SendTelegram(TELEGRAM_CHAT_ID, "The Pathfinders are NOT sended! "+err)
                        } else {
                            Print("The Pathfinder is NOT sended! "+err)
                            SendTelegram(TELEGRAM_CHAT_ID, "The Pathfinder is NOT sended! "+err)
                        }
                    }
                }
            }
        } else {
            for slots == totalSlots {
                if err != 0 {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(2*60))
                    Sleep(120000)
                    ships, _ = origin.GetShips()
                    if ships.Pathfinder > 0 {slots = GetSlots().InUse}
                    err = nil
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(2*60))
                    Sleep(120000)
                    slots = GetSlots().InUse
                }
                curSystem = system-1
            }
        }
        if system >= toSystem {
            if times > 0 {
                if cycle < times {
                    cycle++
                    if nbr == 0 {Print("Not found any debris! Start searching again...")}
                    curSystem = fromSystem-1
                    system = curSystem
                    Sleep(4000)
                } else {
                    Print("You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")
                    SendTelegram(TELEGRAM_CHAT_ID, "You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")
                    break
                }
            } else {
                Print("You made full scan all systems chosen by you! The script turns off")
                SendTelegram(TELEGRAM_CHAT_ID, "You made full scan all systems chosen by you! The script turns off")
                break
            }
        }
    }
} else {Print("You don't have Pathfinders on your Planets/Moons!")}
