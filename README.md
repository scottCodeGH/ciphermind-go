# 🧠 CipherMind - Terminal Mastermind Puzzle Game

A smart and engaging terminal-based code-breaking puzzle game written in Go, inspired by the classic Mastermind board game.

## 🎮 Game Description

CipherMind challenges you to crack a secret code using logic and deduction. The computer generates a random 4-symbol code from 6 possible symbols (A-F), and you have 10 attempts to guess it correctly. After each guess, you receive clues that help you narrow down the possibilities:

- **Green dots (●)**: Symbols that are correct and in the right position
- **Yellow dots (●)**: Symbols that are correct but in the wrong position

## ✨ Features

- **Classic Mastermind Gameplay**: Traditional code-breaking logic puzzle
- **Colorful Terminal Interface**: ANSI color-coded output for better readability
- **Smart Feedback System**: Precise clues after each guess
- **Input Validation**: Robust error handling for invalid inputs
- **Encouraging Messages**: Dynamic feedback based on your progress
- **Guess History**: Track all your previous attempts and their feedback
- **Replay Option**: Play multiple rounds without restarting

## 🚀 Quick Start

### Prerequisites

- Go 1.16 or higher

### Build and Run

```bash
# Build the game
go build -o ciphermind main.go

# Run the game
./ciphermind
```

Or run directly without building:

```bash
go run main.go
```

## 🎯 How to Play

1. The game generates a secret 4-symbol code using symbols A, B, C, D, E, and F
2. You have 10 attempts to guess the code
3. Enter your guess (e.g., `ABCD`) and press Enter
4. Review the feedback:
   - Green dots show how many symbols are in the correct position
   - Yellow dots show how many symbols are correct but in the wrong position
5. Use the clues to refine your next guess
6. Win by guessing the exact code before running out of attempts!

### Example Gameplay

```
╔════════════════════════════════════════════╗
║      🧠 CIPHERMIND - MASTERMIND PUZZLE 🧠  ║
╚════════════════════════════════════════════╝

[Attempt 1/10] Enter your guess: ABCD
Attempt 1: ABCD → ●● (1 exact, 1 misplaced)

[Attempt 2/10] Enter your guess: AECF
Attempt 2: AECF → ●●● (2 exact, 1 misplaced)

[Attempt 3/10] Enter your guess: AEFC
🎉 CONGRATULATIONS! 🎉
You cracked the code AEFC in 3 attempts!
```

## 🧩 Strategy Tips

1. **Start Broad**: Try different symbols in your first guesses to identify which ones are in the code
2. **Track Patterns**: Keep notes on which positions yielded exact matches
3. **Process of Elimination**: Use yellow dots to determine where symbols DON'T belong
4. **Logical Deduction**: Combine clues from multiple guesses to narrow possibilities
5. **Stay Systematic**: Don't randomly guess - use the feedback strategically

## 🛠️ Technical Highlights

- **Pure Go Implementation**: No external dependencies required
- **Efficient Algorithm**: Accurate feedback calculation with two-pass matching
- **Cross-Platform**: Works on Linux, macOS, and Windows terminals
- **Safe & Robust**: Comprehensive error handling and input validation
- **Clean Code**: Well-structured with clear separation of concerns

## 📝 Code Structure

```
main.go
├── Game struct: Core game state management
├── generateSecretCode(): Random code generation
├── evaluateGuess(): Feedback calculation logic
├── validateGuess(): Input validation
├── printWelcome(): Game introduction
├── printGuessHistory(): Display past attempts
└── Run(): Main game loop
```

## 🎨 Customization

You can easily customize the game by modifying these constants in `main.go`:

```go
const (
    CodeLength       = 4      // Length of secret code
    MaxAttempts      = 10     // Number of allowed guesses
    AvailableSymbols = "ABCDEF" // Symbols to use
)
```

## 🏆 Why CipherMind?

This game showcases Go's strengths:
- **Performance**: Instant response times even with complex logic
- **Safety**: No crashes or panics with robust error handling
- **Simplicity**: Clean, readable code that's easy to understand
- **Portability**: Runs anywhere Go runs, no graphics dependencies

## 📜 License

This project is open source and available for educational and entertainment purposes.

## 🙏 Credits

Inspired by the classic Mastermind board game invented by Mordecai Meirowitz in 1970.

---

**Enjoy cracking codes! 🧠🔐**
