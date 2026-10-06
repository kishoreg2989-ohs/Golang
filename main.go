package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"parking-lot/parking"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\n=== PARKING LOT MENU ===")
        fmt.Println("1. Create slot")
        fmt.Println("2. Park vehicle")
        fmt.Println("3. Find slot")
        fmt.Println("4. Decommission slot")
        fmt.Println("5. Show all slots")
        fmt.Println("6. Exit")

        choice, err := readString(scanner, "Choose an option")
        if err != nil {
            fmt.Println("Invalid input.")
            continue
        }

        switch choice {
        case "1":
            slotID, err := readInt(scanner, "Slot ID")
            if err != nil {
                fmt.Println("Invalid slot ID.")
                continue
            }

            floor, err := readInt(scanner, "Floor")
            if err != nil {
                fmt.Println("Invalid floor number.")
                continue
            }

            vehicleType, err := readString(scanner, "Vehicle type (Car/Bike/EV)")
            if err != nil {
                fmt.Println("Vehicle type is required.")
                continue
            }

            normalizedType := normalizeVehicleType(vehicleType)
            if normalizedType == "" {
                fmt.Println("Invalid vehicle type. Use Car, Bike or EV.")
                continue
            }

            err = parking.CreateSlot(parking.ParkingSlot{
                SlotID:      slotID,
                Floor:       floor,
                VehicleType: normalizedType,
            })
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Printf("Slot %d created successfully.\n", slotID)
            }

        case "2":
            slotID, err := readInt(scanner, "Slot ID")
            if err != nil {
                fmt.Println("Invalid slot ID.")
                continue
            }

            plate, err := readString(scanner, "Vehicle plate")
            if err != nil {
                fmt.Println("Plate number is required.")
                continue
            }

            err = parking.ParkVehicle(slotID, plate)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Printf("Vehicle %s parked in slot %d.\n", plate, slotID)
            }

        case "3":
            slotID, err := readInt(scanner, "Slot ID")
            if err != nil {
                fmt.Println("Invalid slot ID.")
                continue
            }

            slot, err := parking.FindSlot(slotID)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Println("Slot found:")
                fmt.Printf("  Slot ID:      %d\n", slot.SlotID)
                fmt.Printf("  Floor:        %d\n", slot.Floor)
                fmt.Printf("  Vehicle Type: %s\n", slot.VehicleType)
                fmt.Printf("  Occupied:     %t\n", slot.IsOccupied)
                fmt.Printf("  Parked Plate: %s\n", slot.ParkedPlate)
            }

        case "4":
            slotID, err := readInt(scanner, "Slot ID")
            if err != nil {
                fmt.Println("Invalid slot ID.")
                continue
            }

            err = parking.DecommissionSlot(slotID)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Printf("Slot %d decommissioned successfully.\n", slotID)
            }

        case "5":
            allSlots := parking.GetAllSlots()
            if len(allSlots) == 0 {
                fmt.Println("No parking slots available.")
                continue
            }

            fmt.Println("Current slots:")
            for _, slot := range allSlots {
                status := "Free"
                if slot.IsOccupied {
                    status = "Occupied"
                }
                fmt.Printf("  Slot %d | Floor %d | %s | %s | Plate: %s\n",
                    slot.SlotID,
                    slot.Floor,
                    slot.VehicleType,
                    status,
                    slot.ParkedPlate,
                )
            }

        case "6":
            fmt.Println("Goodbye!")
            return

        default:
            fmt.Println("Invalid option. Please choose from 1 to 6.")
        }
    }
}

func readString(scanner *bufio.Scanner, promptText string) (string, error) {
    fmt.Print(promptText, ": ")

    if !scanner.Scan() {
        return "", fmt.Errorf("input error")
    }

    text := strings.TrimSpace(scanner.Text())
    if text == "" {
        return "", fmt.Errorf("empty input")
    }

    return text, nil
}

func readInt(scanner *bufio.Scanner, label string) (int, error) {
    value, err := readString(scanner, label)
    if err != nil {
        return 0, err
    }

    parsed, err := strconv.Atoi(value)
    if err != nil {
        return 0, err
    }

    return parsed, nil
}

func normalizeVehicleType(value string) string {
    switch strings.ToLower(strings.TrimSpace(value)) {
    case "car":
        return "Car"
    case "bike":
        return "Bike"
    case "ev":
        return "EV"
    default:
        return ""
    }
}
