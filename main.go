package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	currentPlayer := CreatePlayerWithPosition(50.0, 50.0)

	rl.SetTargetFPS(60)

	// establishing connection to server
	connection, connected := establishConnection()
	setUserName(os.Args[1])

	for !connected && !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Connecting to server...", 190, 200, 20, rl.LightGray)
		rl.EndDrawing()
	}

	for !rl.WindowShouldClose() {
		// connection stuff
		encodeAndSend(connection, currentPlayer.position.X, currentPlayer.position.Y)
		playerPositions, status := receiveDataAndDecode(connection)

		// Handle Events
		if rl.IsKeyDown(rl.KeyUp) {
			applyForce(currentPlayer, rl.Vector2{X: 0.0, Y: -1.0})
		} else if rl.IsKeyDown(rl.KeyDown) {
			applyForce(currentPlayer, rl.Vector2{X: 0.0, Y: 1.0})
		} else if rl.IsKeyDown(rl.KeyLeft) {
			applyForce(currentPlayer, rl.Vector2{X: -1.0, Y: 0.0})
		} else if rl.IsKeyDown(rl.KeyRight) {
			applyForce(currentPlayer, rl.Vector2{X: 1.0, Y: 0.0})
		} else {
			applyForce(currentPlayer, rl.Vector2Scale(currentPlayer.velocity, -0.2))
		}

		// rendering stuff
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		showPlayer(*currentPlayer)
		updatePlayer(currentPlayer)

		if status {
			for UserId, playerPos := range playerPositions {
				if UserId != userName {
					rl.DrawRectangle(int32(playerPos.X), int32(playerPos.Y), 20, 20, rl.Green)
					showPlayerName(rl.Vector2(playerPos), UserId, 0, -10)
				}
			}
		}

		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
	connection.Close()
	rl.CloseWindow()
}
