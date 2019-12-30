//==== This script is created by RockClubKASHMIR ====

fromSystem = 1 // Your can change this value as you want
toSystem = 499 // Your can change this value as you want
times = 1 // if times = 1, the script will full scan 2 times the galaxy, from system, to system you want. You can set this value from 0, to the number you want

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
    if ships.Recycler > flts {
        flts = ships.Recycler
        origin = celestial // Your Planet(or Moon), with more Recyclers
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    for system = curSystem; system <= toSystem; system++ {
        systemInfos, b = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        Dtarget = NewCoordinate(origin.GetCoordinate().Galaxy, system, 17, PLANET_TYPE)
        Sleep(Random(500, 1500)) // for avoid ban
        slots = GetSlots().InUse
        if err != nil {slots = totalSlots}
        if slots < totalSlots {
            if b == nil {
                Print("Checking "+Dtarget)
                if systemInfos.Events.Darkmatter > 0 {
                    ships, _ = origin.GetShips()
                    Print("Found DM Debris: "+systemInfos.Events.Darkmatter)
                    f = NewFleet()
                    f.SetOrigin(origin)
                    f.SetDestination(Dtarget)
                    f.SetSpeed(HUNDRED_PERCENT)
                    f.SetMission(RECYCLEDEBRISFIELD)
                    nbr = systemInfos.Events.Darkmatter
                    f.AddShips(RECYCLER, 1)
                    a, err = f.SendNow()
                    if err == nil {
                        Print("The Recycler is sended successfully!")
                    } else {
                        Print("The Recycler is NOT sended! "+err)
                        SendTelegram(TELEGRAM_CHAT_ID, "The Recycler is NOT sended! "+err)
                    }
                }
            }
        } else {
            for slots == totalSlots {
                if err != 0 {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(2*60))
                    Sleep(120000)
                    ships, _ = origin.GetShips()
                    if ships.Recycler > 0 {slots = GetSlots().InUse}
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
} else {Print("You don't have Recyclers on your Planets/Moons!")}
