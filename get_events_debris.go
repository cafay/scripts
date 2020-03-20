//***** This script is created by RockClubKASHMIR *****\\

   
origin = "M:1:2:3" // Replace "M:1:2:3" with your coordinates - M for the moon, P for planet

fromSystem = 1 // Set from what system you want start to scan
toSystem = 499 // Set to what system you want to end to scan

Telegram = false // Do you want to have TELEGRAM messages?  YES = true / NO = false
Pnbr = 3  // Will ignore debris less than for PATHFINDER with quantity as this value. The maximum is not limited even if you left this value as it is! Change it if/as you want.
times = 1 // if times = 1, the script will full scan 6 times the entire galaxy, from system, to system you set. You can set this value from 1, to the number you want
useCycles = false // Do you want to use the limited repeats?  YES = true / NO = false


//----
cycle = 0
wrong = 0
flag = 0
homeworld = nil
for celestial in GetCachedCelestials() {
    if GetCachedCelestial(celestial) == GetCachedCelestial(origin) {
        homeworld = GetCachedCelestial(origin)
    } else {flag++}
    if flag == 1 {wrong = origin}
}
nbr = 0
err = nil
if (Pnbr < 1) {Pnbr = 1}
if (times < 0) {times = 0}
if useCycles != false && useCycles != true {useCycles = false}
totalSlots = GetSlots().Total - GetFleetSlotsReserved()
curSystem = fromSystem
if homeworld != nil {
    origin = homeworld
    Print("Your origin is "+origin.Coordinate)
    if toSystem > 499 || toSystem == 0 {toSystem = -1}
    if fromSystem > toSystem {Print("Please, type correctly fromSystem and/or toSystem!")}
    for system = curSystem; system <= toSystem; system++ {
        pp = 0
        dflag = 0
        abr = 0
        nbr = 0
        systemInfos, b = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        Dtarget, _ = ParseCoord(origin.GetCoordinate().Galaxy+":"+system+":"+17)
        Debris, _ = ParseCoord("D:"+origin.GetCoordinate().Galaxy+":"+system+":"+17)
        Sleep(Random(500, 1500)) // for avoid ban
        slots = GetSlots().InUse
        if err != nil {slots = totalSlots}
        if slots < totalSlots {
            if b == nil {
                Print("Checking "+Dtarget)
                if systemInfos.ExpeditionDebris.PathfindersNeeded >= Pnbr { 
                    ships, _ = origin.GetShips()
                    pp = systemInfos.ExpeditionDebris.PathfindersNeeded
                    if systemInfos.ExpeditionDebris.Metal == 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                    if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal == 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal)}
                    if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal+" and Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                    fleet, _ = GetFleets()
                    for f in fleet {
                        if f.Mission == RECYCLEDEBRISFIELD && f.ReturnFlight == false {
                            if Debris == f.Destination {
                                if f.Ships.Pathfinder < pp {
                                    abr = pp - f.Ships.Pathfinder
                                } else {dflag = 1}
                            }
                        }
                    }
                    if dflag == 0 {
                        f = NewFleet()
                        f.SetOrigin(origin)
                        f.SetDestination(Dtarget)
                        f.SetSpeed(HUNDRED_PERCENT)
                        f.SetMission(RECYCLEDEBRISFIELD)
                        if abr == 0 {
                            nbr = systemInfos.ExpeditionDebris.PathfindersNeeded
                        } else {nbr = abr}
                        if nbr > ships.Pathfinder {nbr = ships.Pathfinder}
                        if ships.Pathfinder != 0 {
                            f.AddShips(PATHFINDER, nbr)
                        }
                        a, err = f.SendNow()
                        if err == nil {
                            if nbr < systemInfos.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                            if nbr > 1 {
                                Print(nbr+" Pathfinders are sended successfully!")
                            } else {Print(nbr+" Pathfinder is sended successfully!")}
                        } else {
                            if nbr > 1 {
                                Print("The Pathfinders are NOT sended! "+err)
                            } else {Print("The Pathfinder is NOT sended! "+err)}
                        }
                    } else {Print("Needed ships already are sended!")}
                }
            }
        } else {
            for slots == totalSlots {
                delay = Random(4*60, 8*60)
                if err != nil {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(delay))
                    Sleep(delay*1000)
                    ships, _ = origin.GetShips()
                    if ships.Pathfinder > 0 {slots = GetSlots().InUse}
                    err = nil
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(delay))
                    Sleep(delay*1000)
                    slots = GetSlots().InUse
                }
                curSystem = system-1
            }
        }
        if b == nil {
            if system >= toSystem {
                if useCycles == true {
                    if times > 0 {
                        if cycle < times {
                            cycle++
                            if nbr == 0 {Print("Not found any debris! Start searching again...")}
                            curSystem = fromSystem-1
                            system = curSystem
                            delay = Random(30*60, 45*60)
                            Print("Will Start searching again after "+ShortDur(delay))
                            Sleep(delay*1000)
                        } else {
                            Print("You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")
                            if Telegram == true {SendTelegram(TELEGRAM_CHAT_ID, "You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")}
                            break
                        }
                    } else {
                        Print("You made full scan all systems chosen by you! The script turns off")
                        if Telegram == true {SendTelegram(TELEGRAM_CHAT_ID, "You made full scan all systems chosen by you! The script turns off")}
                        break
                    }
                } else {
                    if nbr == 0 {Print("Not found any debris! Start searching again...")}
                    curSystem = fromSystem-1
                    system = curSystem
                    delay = Random(30*60, 45*60)
                    Print("Will Start searching again after "+ShortDur(delay))
                    Sleep(delay*1000)
                }
            }
        } else {
            Print("Please, type correctly fromSystem and/or toSystem!")
            break
        }
    }
} else {Print("You input wrong coordinates! - "+wrong)}
