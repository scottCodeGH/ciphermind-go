package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ANSI color codes for terminal output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

// Game configuration
const (
	CodeLength   = 4
	MaxAttempts  = 10
	AvailableSymbols = "ABCDEF" // Using 6 colors like classic Mastermind
)

// Game represents the state of the Mastermind game
type Game struct {
	secretCode   string
	attempts     int
	maxAttempts  int
	guessHistory []GuessResult
}

// GuessResult stores a guess and its feedback
type GuessResult struct {
	guess        string
	exactMatches int
	partialMatches int
}

// NewGame creates a new Mastermind game with a random secret code
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	game := &Game{
		maxAttempts:  MaxAttempts,
		guessHistory: make([]GuessResult, 0),
	}

	game.secretCode = game.generateSecretCode()
	return game
}

// generateSecretCode creates a random code using available symbols
func (g *Game) generateSecretCode() string {
	code := make([]byte, CodeLength)
	for i := 0; i < CodeLength; i++ {
		code[i] = AvailableSymbols[rand.Intn(len(AvailableSymbols))]
	}
	return string(code)
}

// evaluateGuess compares the guess with the secret code and returns feedback
func (g *Game) evaluateGuess(guess string) GuessResult {
	exactMatches := 0
	partialMatches := 0

	// Create maps to track which positions have been matched
	secretUsed := make([]bool, CodeLength)
	guessUsed := make([]bool, CodeLength)

	// First pass: find exact matches
	for i := 0; i < CodeLength; i++ {
		if guess[i] == g.secretCode[i] {
			exactMatches++
			secretUsed[i] = true
			guessUsed[i] = true
		}
	}

	// Second pass: find partial matches (correct symbol, wrong position)
	for i := 0; i < CodeLength; i++ {
		if !guessUsed[i] {
			for j := 0; j < CodeLength; j++ {
				if !secretUsed[j] && guess[i] == g.secretCode[j] {
					partialMatches++
					secretUsed[j] = true
					break
				}
			}
		}
	}

	return GuessResult{
		guess:          guess,
		exactMatches:   exactMatches,
		partialMatches: partialMatches,
	}
}

// validateGuess checks if the input is a valid guess
func (g *Game) validateGuess(input string) (string, error) {
	guess := strings.ToUpper(strings.TrimSpace(input))

	if len(guess) != CodeLength {
		return "", fmt.Errorf("guess must be exactly %d symbols", CodeLength)
	}

	for _, char := range guess {
		if !strings.ContainsRune(AvailableSymbols, char) {
			return "", fmt.Errorf("invalid symbol '%c'. Use only: %s", char, AvailableSymbols)
		}
	}

	return guess, nil
}

// printWelcome displays the game introduction
func printWelcome() {
	fmt.Println(ColorBold + ColorCyan + "\n╔════════════════════════════════════════════╗" + ColorReset)
	fmt.Println(ColorBold + ColorCyan + "║      🧠 CIPHERMIND - MASTERMIND PUZZLE 🧠  ║" + ColorReset)
	fmt.Println(ColorBold + ColorCyan + "╚════════════════════════════════════════════╝" + ColorReset)
	fmt.Println()
	fmt.Println(ColorYellow + "Welcome to CipherMind!" + ColorReset)
	fmt.Println("I've created a secret code using " + ColorBold + fmt.Sprint(CodeLength) + " symbols" + ColorReset + ".")
	fmt.Println("Your mission: crack the code in " + ColorBold + fmt.Sprint(MaxAttempts) + " attempts or less!" + ColorReset)
	fmt.Println()
	fmt.Println(ColorPurple + "Available symbols: " + ColorBold + AvailableSymbols + ColorReset)
	fmt.Println()
	fmt.Println("After each guess, I'll give you clues:")
	fmt.Println("  " + ColorGreen + "●" + ColorReset + " Green dots = symbols in the correct position")
	fmt.Println("  " + ColorYellow + "●" + ColorReset + " Yellow dots = correct symbols but wrong position")
	fmt.Println()
	fmt.Println(ColorCyan + "Let's begin! Enter your guess (e.g., ABCD):" + ColorReset)
	fmt.Println()
}

// printGuessHistory displays all previous guesses and their feedback
func (g *Game) printGuessHistory() {
	if len(g.guessHistory) == 0 {
		return
	}

	fmt.Println(ColorBold + "\n--- Guess History ---" + ColorReset)
	for i, result := range g.guessHistory {
		fmt.Printf("Attempt %d: %s%s%s → ",
			i+1,
			ColorBold,
			result.guess,
			ColorReset)

		// Print green dots for exact matches
		for j := 0; j < result.exactMatches; j++ {
			fmt.Print(ColorGreen + "●" + ColorReset)
		}

		// Print yellow dots for partial matches
		for j := 0; j < result.partialMatches; j++ {
			fmt.Print(ColorYellow + "●" + ColorReset)
		}

		// Print description
		if result.exactMatches == 0 && result.partialMatches == 0 {
			fmt.Print(ColorRed + " None correct" + ColorReset)
		} else {
			fmt.Printf(" (%d exact, %d misplaced)", result.exactMatches, result.partialMatches)
		}

		fmt.Println()
	}
	fmt.Println()
}

// printEncouragement gives friendly feedback based on progress
func printEncouragement(exactMatches int, attempts int) {
	messages := []string{}

	if exactMatches == CodeLength-1 {
		messages = []string{
			"So close! Just one more symbol!",
			"You're almost there! One more to go!",
			"Nearly cracked it! Keep going!",
		}
	} else if exactMatches >= CodeLength/2 {
		messages = []string{
			"Good progress! You're on the right track!",
			"Nice work! You're getting warmer!",
			"Excellent deduction! Keep it up!",
		}
	} else if attempts > MaxAttempts/2 {
		messages = []string{
			"Don't give up! Try a different approach!",
			"Hmm, time to think outside the box!",
			"Keep analyzing the clues!",
		}
	} else {
		messages = []string{
			"Interesting guess! Study the feedback carefully.",
			"Use the clues to narrow down the possibilities!",
			"Logic will lead you to victory!",
		}
	}

	if len(messages) > 0 {
		fmt.Println(ColorCyan + messages[rand.Intn(len(messages))] + ColorReset)
	}
}

// Run starts the game loop
func (g *Game) Run() {
	printWelcome()

	scanner := bufio.NewScanner(os.Stdin)

	for g.attempts < g.maxAttempts {
		// Show remaining attempts
		remaining := g.maxAttempts - g.attempts
		fmt.Printf(ColorBlue+"[Attempt %d/%d]"+ColorReset+" Enter your guess: ",
			g.attempts+1, g.maxAttempts)

		// Read input
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		// Validate input
		guess, err := g.validateGuess(input)
		if err != nil {
			fmt.Println(ColorRed + "❌ " + err.Error() + ColorReset)
			continue
		}

		// Evaluate guess
		g.attempts++
		result := g.evaluateGuess(guess)
		g.guessHistory = append(g.guessHistory, result)

		// Check for win
		if result.exactMatches == CodeLength {
			g.printGuessHistory()
			fmt.Println(ColorGreen + ColorBold + "\n🎉 CONGRATULATIONS! 🎉" + ColorReset)
			fmt.Printf(ColorGreen+"You cracked the code "+ColorBold+"%s"+ColorReset+ColorGreen+" in %d attempts!\n"+ColorReset,
				g.secretCode, g.attempts)
			fmt.Println(ColorYellow + "Your deduction skills are impressive!" + ColorReset)
			return
		}

		// Show feedback
		g.printGuessHistory()
		printEncouragement(result.exactMatches, g.attempts)

		// Show remaining attempts warning
		if remaining <= 3 {
			fmt.Printf(ColorRed+"⚠️  Only %d attempts remaining!"+ColorReset+"\n", remaining-1)
		}
		fmt.Println()
	}

	// Game over - player lost
	g.printGuessHistory()
	fmt.Println(ColorRed + ColorBold + "\n💀 GAME OVER 💀" + ColorReset)
	fmt.Printf(ColorRed+"You ran out of attempts! The secret code was: "+ColorBold+"%s\n"+ColorReset,
		g.secretCode)
	fmt.Println(ColorYellow + "Better luck next time, code breaker!" + ColorReset)
}

func main() {
	game := NewGame()
	game.Run()

	// Ask if player wants to play again
	fmt.Println()
	fmt.Print("Play again? (y/n): ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		response := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if response == "y" || response == "yes" {
			fmt.Println()
			main() // Recursive call to start new game
		}
	}

	fmt.Println(ColorCyan + "\nThanks for playing CipherMind! 🧠" + ColorReset)
}
