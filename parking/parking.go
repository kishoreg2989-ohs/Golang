package parking

import (
    "fmt"
    "sort"
)

// ParkingSlot represents a parking slot in the garage.
type ParkingSlot struct {
    SlotID      int
    Floor       int
    VehicleType string
    IsOccupied  bool
    ParkedPlate string
}

var slots = make(map[int]*ParkingSlot)

func CreateSlot(slot ParkingSlot) error {
    if slot.SlotID <= 0 {
        return fmt.Errorf("slot ID must be greater than 0")
    }
    if slot.Floor <= 0 {
        return fmt.Errorf("floor must be greater than 0")
    }

    switch slot.VehicleType {
    case "Car", "Bike", "EV":
    default:
        return fmt.Errorf("invalid vehicle type %q: must be Car, Bike, or EV", slot.VehicleType)
    }

    if _, exists := slots[slot.SlotID]; exists {
        return fmt.Errorf("slot %d already exists", slot.SlotID)
    }

    slot.IsOccupied = false
    slot.ParkedPlate = ""
    slots[slot.SlotID] = &slot
    return nil
}

func FindSlot(slotID int) (*ParkingSlot, error) {
    slot, exists := slots[slotID]
    if !exists {
        return nil, fmt.Errorf("slot %d not found", slotID)
    }
    if slot.IsOccupied {
        return nil, fmt.Errorf("slot %d is occupied by vehicle %s", slotID, slot.ParkedPlate)
    }
    return slot, nil
}

func ParkVehicle(slotID int, plate string) error {
    slot, exists := slots[slotID]
    if !exists {
        return fmt.Errorf("slot %d not found", slotID)
    }
    if slot.IsOccupied {
        return fmt.Errorf("slot %d is already occupied by vehicle %s", slotID, slot.ParkedPlate)
    }
    if plate == "" {
        return fmt.Errorf("plate number cannot be empty")
    }

    slot.IsOccupied = true
    slot.ParkedPlate = plate
    return nil
}

func DecommissionSlot(slotID int) error {
    if _, exists := slots[slotID]; !exists {
        return fmt.Errorf("slot %d not found", slotID)
    }
    delete(slots, slotID)
    return nil
}

func GetAllSlots() []ParkingSlot {
    all := make([]ParkingSlot, 0, len(slots))
    for _, slot := range slots {
        all = append(all, *slot)
    }
    sort.Slice(all, func(i, j int) bool {
        return all[i].SlotID < all[j].SlotID
    })
    return all
}
