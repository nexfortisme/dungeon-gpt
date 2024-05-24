// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // 	"math/rand"
// // // // 	"time"
// // // // )

// // // // type Room struct {
// // // // 	ConnectingRooms []int
// // // // 	Items           []string
// // // // 	Monsters        []string
// // // // }

// // // // func generateDungeon(numRooms int) map[int]Room {
// // // // 	rand.Seed(time.Now().UnixNano())

// // // // 	items := []string{"sword", "armor", "key", "torch"}
// // // // 	monsters := []string{"monster1", "monster2", "monster3"}

// // // // 	// Create the basic dungeon structure
// // // // 	dungeon := make(map[int]Room)
// // // // 	for i := 1; i <= numRooms; i++ {
// // // // 		dungeon[i] = Room{
// // // // 			ConnectingRooms: []int{},
// // // // 			Items:           []string{},
// // // // 			Monsters:        []string{},
// // // // 		}
// // // // 	}

// // // // 	// Ensure each room has at least one connection
// // // // 	for i := 1; i <= numRooms; i++ {
// // // // 		for len(dungeon[i].ConnectingRooms) == 0 {
// // // // 			connection := rand.Intn(numRooms) + 1
// // // // 			if connection != i {
// // // // 				room := dungeon[i]
// // // // 				room.ConnectingRooms = append(room.ConnectingRooms, connection)
// // // // 				dungeon[i] = room

// // // // 				if rand.Intn(2) == 0 {
// // // // 					connectedRoom := dungeon[connection]
// // // // 					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
// // // // 					dungeon[connection] = connectedRoom
// // // // 				}
// // // // 			}
// // // // 		}
// // // // 	}

// // // // 	// Place items randomly
// // // // 	for _, item := range items {
// // // // 		roomIndex := rand.Intn(numRooms) + 1
// // // // 		room := dungeon[roomIndex]
// // // // 		room.Items = append(room.Items, item)
// // // // 		dungeon[roomIndex] = room
// // // // 	}

// // // // 	// Place monsters randomly
// // // // 	for _, monster := range monsters {
// // // // 		roomIndex := rand.Intn(numRooms) + 1
// // // // 		room := dungeon[roomIndex]
// // // // 		room.Monsters = append(room.Monsters, monster)
// // // // 		dungeon[roomIndex] = room
// // // // 	}

// // // // 	return dungeon
// // // // }

// // // //	func main() {
// // // //		numRooms := 20 // Define the number of rooms in the dungeon
// // // //		dungeon := generateDungeon(numRooms)
// // // //		for room, details := range dungeon {
// // // //			fmt.Printf("%d : { connectingRooms: %v, items: %v, monsters: %v }\n", room, details.ConnectingRooms, details.Items, details.Monsters)
// // // //		}
// // // //	}
// // // package main

// // // import (
// // // 	"fmt"
// // // 	"math/rand"
// // // 	"time"
// // // )

// // // type Room struct {
// // // 	ConnectingRooms []int
// // // 	Items           []string
// // // 	Monsters        []string
// // // }

// // // func generateDungeon(numRooms int) map[int]Room {
// // // 	rand.Seed(time.Now().UnixNano())

// // // 	items := []string{"sword", "armor", "key", "torch"}
// // // 	monsters := []string{"monster1", "monster2", "monster3"}

// // // 	// Create the basic dungeon structure
// // // 	dungeon := make(map[int]Room)
// // // 	for i := 1; i <= numRooms; i++ {
// // // 		dungeon[i] = Room{
// // // 			ConnectingRooms: []int{},
// // // 			Items:           []string{},
// // // 			Monsters:        []string{},
// // // 		}
// // // 	}

// // // 	// Ensure each room has at least one connection
// // // 	for i := 1; i <= numRooms; i++ {
// // // 		for len(dungeon[i].ConnectingRooms) == 0 {
// // // 			connection := rand.Intn(numRooms) + 1
// // // 			if connection != i {
// // // 				room := dungeon[i]
// // // 				room.ConnectingRooms = append(room.ConnectingRooms, connection)
// // // 				dungeon[i] = room

// // // 				if rand.Intn(2) == 0 {
// // // 					connectedRoom := dungeon[connection]
// // // 					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
// // // 					dungeon[connection] = connectedRoom
// // // 				}
// // // 			}
// // // 		}
// // // 	}

// // // 	// Ensure more rooms have at least 2 connections
// // // 	for i := 1; i <= numRooms; i++ {
// // // 		for len(dungeon[i].ConnectingRooms) < 2 {
// // // 			connection := rand.Intn(numRooms) + 1
// // // 			if connection != i && !contains(dungeon[i].ConnectingRooms, connection) {
// // // 				room := dungeon[i]
// // // 				room.ConnectingRooms = append(room.ConnectingRooms, connection)
// // // 				dungeon[i] = room

// // // 				if rand.Intn(2) == 0 {
// // // 					connectedRoom := dungeon[connection]
// // // 					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
// // // 					dungeon[connection] = connectedRoom
// // // 				}
// // // 			}
// // // 		}
// // // 	}

// // // 	// Place items randomly
// // // 	for _, item := range items {
// // // 		roomIndex := rand.Intn(numRooms) + 1
// // // 		room := dungeon[roomIndex]
// // // 		room.Items = append(room.Items, item)
// // // 		dungeon[roomIndex] = room
// // // 	}

// // // 	// Place monsters randomly
// // // 	for _, monster := range monsters {
// // // 		roomIndex := rand.Intn(numRooms) + 1
// // // 		room := dungeon[roomIndex]
// // // 		room.Monsters = append(room.Monsters, monster)
// // // 		dungeon[roomIndex] = room
// // // 	}

// // // 	return dungeon
// // // }

// // // // Utility function to check if a slice contains an element
// // // func contains(slice []int, element int) bool {
// // // 	for _, e := range slice {
// // // 		if e == element {
// // // 			return true
// // // 		}
// // // 	}
// // // 	return false
// // // }

// // // func main() {
// // // 	numRooms := 10 // Define the number of rooms in the dungeon
// // // 	dungeon := generateDungeon(numRooms)
// // // 	for room, details := range dungeon {
// // // 		fmt.Printf("%d : { connectingRooms: %v, items: %v, monsters: %v }\n", room, details.ConnectingRooms, details.Items, details.Monsters)
// // // 	}
// // // }

// // package main

// // import (
// // 	"fmt"
// // 	"math/rand"
// // 	"time"
// // )

// // type Room struct {
// // 	ConnectingRooms []int
// // 	Items           []string
// // 	Monsters        []string
// // }

// // func generateDungeon(numRooms int) map[int]Room {
// // 	rand.Seed(time.Now().UnixNano())

// // 	items := []string{"sword", "armor", "key", "torch"}
// // 	monsters := []string{"monster1", "monster2", "monster3"}

// // 	// Create the basic dungeon structure
// // 	dungeon := make(map[int]Room)
// // 	for i := 1; i <= numRooms; i++ {
// // 		dungeon[i] = Room{
// // 			ConnectingRooms: []int{},
// // 			Items:           []string{},
// // 			Monsters:        []string{},
// // 		}
// // 	}

// // 	// Ensure each room has at least one connection
// // 	for i := 1; i <= numRooms; i++ {
// // 		for len(dungeon[i].ConnectingRooms) == 0 {
// // 			connection := rand.Intn(numRooms) + 1
// // 			if connection != i {
// // 				room := dungeon[i]
// // 				room.ConnectingRooms = append(room.ConnectingRooms, connection)
// // 				dungeon[i] = room

// // 				if rand.Intn(2) == 0 {
// // 					connectedRoom := dungeon[connection]
// // 					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
// // 					dungeon[connection] = connectedRoom
// // 				}
// // 			}
// // 		}
// // 	}

// // 	// Increase connectivity for some rooms
// // 	for i := 1; i <= numRooms; i++ {
// // 		if len(dungeon[i].ConnectingRooms) < 2 && rand.Intn(2) == 0 {
// // 			connection := rand.Intn(numRooms) + 1
// // 			if connection != i && !contains(dungeon[i].ConnectingRooms, connection) {
// // 				room := dungeon[i]
// // 				room.ConnectingRooms = append(room.ConnectingRooms, connection)
// // 				dungeon[i] = room

// // 				if rand.Intn(2) == 0 {
// // 					connectedRoom := dungeon[connection]
// // 					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
// // 					dungeon[connection] = connectedRoom
// // 				}
// // 			}
// // 		}
// // 	}

// // 	// Place items randomly
// // 	for _, item := range items {
// // 		roomIndex := rand.Intn(numRooms) + 1
// // 		room := dungeon[roomIndex]
// // 		room.Items = append(room.Items, item)
// // 		dungeon[roomIndex] = room
// // 	}

// // 	// Place monsters randomly
// // 	for _, monster := range monsters {
// // 		roomIndex := rand.Intn(numRooms) + 1
// // 		room := dungeon[roomIndex]
// // 		room.Monsters = append(room.Monsters, monster)
// // 		dungeon[roomIndex] = room
// // 	}

// // 	return dungeon
// // }

// // // Utility function to check if a slice contains an element
// // func contains(slice []int, element int) bool {
// // 	for _, e := range slice {
// // 		if e == element {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // func main() {
// // 	numRooms := 10 // Define the number of rooms in the dungeon
// // 	dungeon := generateDungeon(numRooms)
// // 	for room, details := range dungeon {
// // 		fmt.Printf("%d : { connectingRooms: %v, items: %v, monsters: %v }\n", room, details.ConnectingRooms, details.Items, details.Monsters)
// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// type Room struct {
// 	ConnectingRooms []int
// 	Items           []string
// 	Monsters        []string
// }

// func generateDungeon(numRooms int) map[int]Room {
// 	rand.Seed(time.Now().UnixNano())

// 	items := []string{"sword", "armor", "key", "torch"}
// 	monsters := []string{"monster1", "monster2", "monster3"}

// 	// Create the basic dungeon structure
// 	dungeon := make(map[int]Room)
// 	for i := 1; i <= numRooms; i++ {
// 		dungeon[i] = Room{
// 			ConnectingRooms: []int{},
// 			Items:           []string{},
// 			Monsters:        []string{},
// 		}
// 	}

// 	// Ensure each room has at least one connection
// 	for i := 1; i <= numRooms; i++ {
// 		for len(dungeon[i].ConnectingRooms) == 0 {
// 			connection := rand.Intn(numRooms) + 1
// 			if connection != i {
// 				room := dungeon[i]
// 				room.ConnectingRooms = append(room.ConnectingRooms, connection)
// 				dungeon[i] = room

// 				if rand.Intn(2) == 0 {
// 					connectedRoom := dungeon[connection]
// 					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
// 					dungeon[connection] = connectedRoom
// 				}
// 			}
// 		}
// 	}

// 	// Increase connectivity for some rooms
// 	for i := 1; i <= numRooms; i++ {
// 		if len(dungeon[i].ConnectingRooms) < 2 && rand.Intn(2) == 0 {
// 			connection := rand.Intn(numRooms) + 1
// 			if connection != i && !contains(dungeon[i].ConnectingRooms, connection) {
// 				room := dungeon[i]
// 				room.ConnectingRooms = append(room.ConnectingRooms, connection)
// 				dungeon[i] = room

// 				if rand.Intn(2) == 0 {
// 					connectedRoom := dungeon[connection]
// 					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
// 					dungeon[connection] = connectedRoom
// 				}
// 			}
// 		}
// 	}

// 	// Ensure no rooms are isolated by performing a DFS
// 	visited := make(map[int]bool)
// 	stack := []int{1}
// 	for len(stack) > 0 {
// 		room := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		if !visited[room] {
// 			visited[room] = true
// 			for _, neighbor := range dungeon[room].ConnectingRooms {
// 				if !visited[neighbor] {
// 					stack = append(stack, neighbor)
// 				}
// 			}
// 		}
// 	}

// 	// Add bi-directional connections to any unvisited rooms
// 	for i := 1; i <= numRooms; i++ {
// 		if !visited[i] {
// 			// Connect this room to a random visited room
// 			for visitedRoom := range visited {
// 				room := dungeon[i]
// 				room.ConnectingRooms = append(room.ConnectingRooms, visitedRoom)
// 				dungeon[i] = room

// 				connectedRoom := dungeon[visitedRoom]
// 				connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
// 				dungeon[visitedRoom] = connectedRoom

// 				visited[i] = true
// 				break
// 			}
// 		}
// 	}

// 	// Place items randomly
// 	for _, item := range items {
// 		roomIndex := rand.Intn(numRooms) + 1
// 		room := dungeon[roomIndex]
// 		room.Items = append(room.Items, item)
// 		dungeon[roomIndex] = room
// 	}

// 	// Place monsters randomly
// 	for _, monster := range monsters {
// 		roomIndex := rand.Intn(numRooms) + 1
// 		room := dungeon[roomIndex]
// 		room.Monsters = append(room.Monsters, monster)
// 		dungeon[roomIndex] = room
// 	}

// 	return dungeon
// }

// // Utility function to check if a slice contains an element
// func contains(slice []int, element int) bool {
// 	for _, e := range slice {
// 		if e == element {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	numRooms := 10 // Define the number of rooms in the dungeon
// 	dungeon := generateDungeon(numRooms)
// 	for room, details := range dungeon {
// 		fmt.Printf("%d : { connectingRooms: %v, items: %v, monsters: %v }\n", room, details.ConnectingRooms, details.Items, details.Monsters)
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Room struct {
	ID              int
	ConnectingRooms []int
	Items           []string
	Monsters        []string
	Exit            bool
}

var Dungeon map[int]Room = make(map[int]Room)

func GenerateDungeon(numRooms int) map[int]Room {
	rand.Seed(time.Now().UnixNano())

	items := []string{"sword", "armor", "key", "torch"}
	monsters := []string{"monster1", "monster2", "monster3"}

	// Create the basic dungeon structure
	// dungeon := make(map[int]Room)
	for i := 1; i <= numRooms; i++ {
		Dungeon[i] = Room{
			ID:              i,
			ConnectingRooms: []int{},
			Items:           []string{},
			Monsters:        []string{},
			Exit:            false,
		}
	}

	// Ensure each room has at least one connection
	for i := 1; i <= numRooms; i++ {
		for len(Dungeon[i].ConnectingRooms) == 0 {
			connection := rand.Intn(numRooms) + 1
			if connection != i {
				room := Dungeon[i]
				room.ConnectingRooms = append(room.ConnectingRooms, connection)
				Dungeon[i] = room

				if rand.Intn(2) == 0 {
					connectedRoom := Dungeon[connection]
					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
					Dungeon[connection] = connectedRoom
				}
			}
		}
	}

	// Increase connectivity for some rooms
	for i := 1; i <= numRooms; i++ {
		if len(Dungeon[i].ConnectingRooms) < 2 && rand.Intn(2) == 0 {
			connection := rand.Intn(numRooms) + 1
			if connection != i && !contains(Dungeon[i].ConnectingRooms, connection) {
				room := Dungeon[i]
				room.ConnectingRooms = append(room.ConnectingRooms, connection)
				Dungeon[i] = room

				if rand.Intn(2) == 0 {
					connectedRoom := Dungeon[connection]
					connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
					Dungeon[connection] = connectedRoom
				}
			}
		}
	}

	// Ensure no rooms are isolated by performing a DFS
	visited := make(map[int]bool)
	stack := []int{1}
	for len(stack) > 0 {
		room := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[room] {
			visited[room] = true
			for _, neighbor := range Dungeon[room].ConnectingRooms {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}

	// Add bi-directional connections to any unvisited rooms
	for i := 1; i <= numRooms; i++ {
		if !visited[i] {
			// Connect this room to a random visited room
			for visitedRoom := range visited {
				room := Dungeon[i]
				room.ConnectingRooms = append(room.ConnectingRooms, visitedRoom)
				Dungeon[i] = room

				connectedRoom := Dungeon[visitedRoom]
				connectedRoom.ConnectingRooms = append(connectedRoom.ConnectingRooms, i)
				Dungeon[visitedRoom] = connectedRoom

				visited[i] = true
				break
			}
		}
	}

	// Place items randomly
	for _, item := range items {
		roomIndex := rand.Intn(numRooms) + 1
		room := Dungeon[roomIndex]
		room.Items = append(room.Items, item)
		Dungeon[roomIndex] = room
	}

	// Place monsters randomly
	for _, monster := range monsters {
		roomIndex := rand.Intn(numRooms) + 1
		room := Dungeon[roomIndex]
		room.Monsters = append(room.Monsters, monster)
		Dungeon[roomIndex] = room
	}

	// Add an additional room for an exit
	exitRoomIndex := numRooms + 1
	exitRoom := Room{
		ID:              exitRoomIndex,
		ConnectingRooms: []int{},
		Items:           []string{},
		Monsters:        []string{},
		Exit:            true,
	}
	Dungeon[exitRoomIndex] = exitRoom

	// Connect the exit room to a random room in the dungeon
	randomRoom := rand.Intn(numRooms) + 1
	room := Dungeon[randomRoom]
	room.ConnectingRooms = append(room.ConnectingRooms, exitRoomIndex)
	Dungeon[randomRoom] = room

	exitRoom.ConnectingRooms = append(exitRoom.ConnectingRooms, randomRoom)
	Dungeon[exitRoomIndex] = exitRoom

	return Dungeon
}

// Utility function to check if a slice contains an element
func contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func main2() {
	numRooms := 10 // Define the number of rooms in the dungeon
	dungeon := GenerateDungeon(numRooms)
	for room, details := range dungeon {
		fmt.Printf("%d : { connectingRooms: %v, items: %v, monsters: %v, exit: %v }\n", room, details.ConnectingRooms, details.Items, details.Monsters, details.Exit)
	}
}

type Player struct {
	location  int
	inventory []string
}

var dungeon = make(map[int]Room)

var player = Player{location: 1, inventory: []string{}}

func main() {

	dungeon = GenerateDungeon(10)

	fmt.Println("Welcome to the Go Adventure Game!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		handleInput(input)
	}
}

func handleInput(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "look":
		lookAround()
	case "pickup":
		if len(parts) < 2 {
			fmt.Println("Pickup what?")
		} else {
			pickupItem(parts[1])
		}
	case "go":
		if len(parts) < 2 {
			fmt.Println("Go where?")
		} else {
			moveToRoom(parts[1])
		}
	case "fight":
		if len(parts) < 2 {
			fmt.Println("Fight whom?")
		} else {
			fightMonster(parts[1])
		}
	case "inventory":
		showInventory()
	default:
		fmt.Println("Unknown command.")
	}
}

func lookAround() {
	currentRoom := dungeon[player.location]
	fmt.Printf("You are in room %d.\n", currentRoom.ID)
	fmt.Printf("Connecting rooms: %v\n", currentRoom.ConnectingRooms)
	if len(currentRoom.Monsters) > 0 {
		fmt.Printf("Monsters here: %v\n", currentRoom.Monsters)
	}
	if len(currentRoom.Items) > 0 {
		fmt.Printf("Items here: %v\n", currentRoom.Items)
	}
	if currentRoom.Exit {
		fmt.Println("This room is an exit.")
	}
}

func pickupItem(item string) {
	currentRoom := dungeon[player.location]
	for i, roomItem := range currentRoom.Items {
		if roomItem == item {
			player.inventory = append(player.inventory, item)
			room := dungeon[player.location]
			room.Items = append(currentRoom.Items[:i], currentRoom.Items[i+1:]...)
			dungeon[player.location] = room
			fmt.Printf("Picked up %s.\n", item)
			return
		}
	}
	fmt.Println("Item not found.")
}

func moveToRoom(roomStr string) {
	roomID := -1
	fmt.Sscanf(roomStr, "%d", &roomID)
	currentRoom := dungeon[player.location]
	for _, connectingRoom := range currentRoom.ConnectingRooms {
		if connectingRoom == roomID {
			player.location = roomID
			fmt.Printf("Moved to room %d.\n", roomID)
			checkWinCondition()
			return
		}
	}
	fmt.Println("You can't go there.")
}

func fightMonster(monster string) {
	currentRoom := dungeon[player.location]
	for i, roomMonster := range currentRoom.Monsters {
		if roomMonster == monster {
			if hasItem("sword") {
				fmt.Printf("You fought and defeated the %s.\n", monster)
				room := dungeon[player.location]
				room.Monsters = append(currentRoom.Monsters[:i], currentRoom.Monsters[i+1:]...)
				dungeon[player.location] = room
			} else if hasItem("armor") {
				fmt.Printf("You escaped from the %s using your armor.\n", monster)
			} else {
				fmt.Printf("You were defeated by the %s.\n", monster)
				os.Exit(0)
			}
			return
		}
	}
	fmt.Println("Monster not found.")
}

func showInventory() {
	fmt.Println("Inventory:", player.inventory)
}

func hasItem(item string) bool {
	for _, invItem := range player.inventory {
		if invItem == item {
			return true
		}
	}
	return false
}

func checkWinCondition() {
	if dungeon[player.location].Exit && hasItem("key") {
		fmt.Println("You have the key and found the exit. You win!")
		os.Exit(0)
	}
}
