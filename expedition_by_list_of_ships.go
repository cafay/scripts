/***** This script is created by RockClubKASHMIR *****\

   DESCRIPTION
  This script will find automatically your planet/moon with ships AND their amount, by the List of ships!
  
  If you want to remove the automatical method of finding your moon/planet,
   replace all rows between // START and // END with origin = GetCachedCelestial("M:1:2:3") where on "M:1:2:3" must type your coordinate - M for the moon, P for planet
*/

fromSystem = 1 // Your can change this value as you want
toSystem = 200 // Your can change this value as you want
shipsList = {LARGECARGO: 200, ESPIONAGEPROBE: 11, BOMBER: 1, DESTROYER: 1}// Your can change ENTYRE List, even to left only 1 type of ships!
DurationOfExpedition = 1 // 1 for one hour, 2 for two hours... set this value from 1, to the number you want

//-------
curSystem = fromSystem
origin = nil
master = 0
nbr = 0
err = nil
// START
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    slt = 0
    flts = 0
    for ShipID, nbr in shipsList {
        if ships.ByID(ShipID) != 0 {
            if ships.ByID(ShipID) >= nbr {
                flts = flts + ships.ByID(ShipID)
                slt = slt + 1
            }
        }
    }
    if slt == len(shipsList) {
        if flts > master {
            master = flts
            origin = celestial // Your Planet(or Moon) with highest amount of ships from the list of ships
        }
    }
}
// END
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    if toSystem > 499 || toSystem == 0 {toSystem = -1}
    if fromSystem > toSystem {Print("Please, type correctly fromSystem and/or toSystem!")}
    for system = curSystem; system <= toSystem; system++ {
        systemInfos, b = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        Destination, _ = ParseCoord(origin.GetCoordinate().Galaxy+":"+system+":"+16)
        totalSlots = GetSlots().Total - GetFleetSlotsReserved()
        slots = GetSlots().InUse
        if slots < totalSlots {
            slots = GetSlots().ExpInUse
            totalSlots = GetSlots().ExpTotal
        }
        if err != nil {slots = totalSlots}
        if slots < totalSlots {
            if b == nil {
                rtt = 0
                tt = 0
                explist = []
                ExpFleet = {}
                myShips, _ = origin.GetShips()
                Sleep(Random(8*1000, 12*1000)) // For avoiding ban
                f = NewFleet()
                f.SetOrigin(origin)
                f.SetDestination(Destination)
                f.SetSpeed(HUNDRED_PERCENT)
                f.SetMission(EXPEDITION)
                if len(shipsList) > 0 {
                for ShipID, num in shipsList {
                    rtt = rtt + 1
                    if myShips.ByID(ShipID) != 0 {
                        if ShipID != PATHFINDER {
                            if myShips.ByID(ShipID) >= num {
                                ExpFleet[ShipID] = num
                                tt = tt + 1
                            }
                        } else {
                            if myShips.ByID(ShipID) < num {num = myShips.ByID(ShipID)}
                            ExpFleet[ShipID] = num
                            tt = tt + 1
                        }
                    }
                }
            }
            if rtt == tt {
                for ShipID, nbr in ExpFleet {
                    f.AddShips(ShipID, nbr)
                    explist += ShipID+": "+nbr
                }
            }
                f.SetDuration(DurationOfExpedition)
                a, err = f.SendNow()
                if err == nil {
                    Print("Sending "+explist+" is successfully to "+Destination)
                    Print("The ships are sended successfully to "+Destination)
                } else {
                    Print("The fleet is NOT sended! "+err)
                    SendTelegram(TELEGRAM_CHAT_ID, "The fleet is NOT sended! "+err)
                }
            }
        } else {
            for slots == totalSlots {
                if err != nil {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(2*60))
                    Sleep(2*60*1000)
                    ships, _ = origin.GetShips()
                    for ShipID in shipsList {
                        if ships.ByID(ShipID) != 0 {slots = GetSlots().ExpInUse}
                        err = nil
                    }
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(4*60))
                    Sleep(4*60*1000)
                    slots = GetSlots().ExpInUse
                }
            }
            curSystem = system-1
        }
        if b == nil {
            if system >= toSystem {
                curSystem = fromSystem-1
                system = curSystem
            }
        } else {
            Print("Please, type correctly fromSystem and/or toSystem!")
            break
        }
    }
} else {Print("Not found any ships from the List of ships on your Planets/Moons!")}
