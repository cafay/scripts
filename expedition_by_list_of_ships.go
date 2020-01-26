//==== This script is created by RockClubKASHMIR ====

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
// Start to Search highest amount of ships from your list to all your Planets and Moons(if you have some)
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    flts = 0
    for ShipID in shipsList {
        if ships.ByID(ShipID) != 0 {
            flts = flts + ships.ByID(ShipID)
        }
    }
    if flts > master {
        master = flts
        origin = celestial // Your Planet(or Moon) with highest amount of ships from the list of ships
    }
}
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
                ships, _ = origin.GetShips()
                Sleep(Random(8*1000, 12*1000)) // For avoiding ban
                f = NewFleet()
                f.SetOrigin(origin)
                f.SetDestination(Destination)
                f.SetSpeed(HUNDRED_PERCENT)
                f.SetMission(EXPEDITION)
                for id, nbr in shipsList {
                    if ships
