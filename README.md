# CodeSync

A real-time collaborative interviewing tool written in Go. CodeSync lets interviewers and candidates spin up a shared coding session—complete with a live code editor, problem statement pane, and “Run” / “Submit” buttons—so you can focus on the conversation, not the setup.

---

## 🚀 Features

- **Real-time Collaboration**: Both interviewer and interviewee edit the same code buffer simultaneously.  
- **LeetCode-Style UI**: Problem description on the left, editable code pane on the right.  
- **Run & Submit**: Instantly compile and run code against built-in test cases, or submit for full test suite.  
- **Question Provisioning via JSON**: Define new problems by dropping in a JSON file containing:
  - Problem metadata (title, description, difficulty)  
  - Test cases (inputs & expected outputs)  
  - Optional images (render diagrams, charts, etc.)  
- **WebRTC-Powered Calls**: Low-latency audio/video channel integrated alongside the editor.  
- **Go Backend**: Lightweight, idiomatic Go for server, WebSocket coordination, and test orchestration.

---

## 📦 Table of Contents

1. [Getting Started](#getting-started)  
2. [Prerequisites](#prerequisites)  
3. [Installation](#installation)  
4. [Running Locally](#running-locally)  
5. [Usage](#usage)  
6. [Question JSON Format](#question-json-format)  
7. [Tech Stack](#tech-stack)  
8. [Contributing](#contributing)  
9. [License](#license)  

---

## 🔧 Getting Started

These instructions will get you a copy of CodeSync up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go 1.22+](https://golang.org/dl/)  
- [Docker](https://www.docker.com/) (optional, if you want to containerize)  

---

## ⚙️ Installation

1. **Clone the repo**  
   ```bash
   git clone https://github.com/your-org/codesync.git
   cd codesync
   ```
2. **Fetch Go dependencies**
```bash
go mod download
```
3. **Build the server**
```bash
make build
```

4. **Run the server**
```bash
make run
```

5. **Clean build**
```bash
make clean
```

## 📝 Usage

- Create a room: Click “New Session” and share the link with your interviewee.
- Upload a question: In the sidebar, click “Load Question” and select your question JSON file.
- Collaborate: As soon as both parties are connected, you’ll see each other’s cursors in the code editor.
- Run & Submit: Use the buttons at the top to compile & run against sample tests, or submit for full evaluation.
- Review results: Pass/fail statuses and stdout/stderr output appear inline.

## 📄 Question JSON Format

Drop your question files into /questions—the server will serve them automatically. Here’s a minimal example:
```json
{
  "id": "two-sum",
  "title": "Two Sum",
  "description": "Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.",
  "difficulty": "Easy",
  "images": [
    "https://example.com/diagram1.png"
  ],
  "template": {
    "language": "go",
    "code": "func twoSum(nums []int, target int) []int {\n    // your code here\n}\n"
  },
  "tests": [
    {
      "input": {
        "nums": [2, 7, 11, 15],
        "target": 9
      },
      "expected": [0, 1]
    },
    {
      "input": {
        "nums": [3, 2, 4],
        "target": 6
      },
      "expected": [1, 2]
    }
  ]
}
```
id: Unique slug for routing & storage.
template.code: Starter code loaded into the editor.
tests: Array of { input: {...}, expected: ... } pairs.

## 🛠️ Tech Stack

- Backend: Go, net/http, Gorilla WebSocket, Docker
- Frontend: TO BE DECIDED
- Data: JSON-based problem loader, in-memory session store (Redis support planned)
- CI/CD: GitHub Actions → Docker Hub


## 🤝 Contributing

Fork the repo
Create a feature branch: git checkout -b feature/awesome-thing
Commit your changes: git commit -m "Add awesome thing"
Push: git push origin feature/awesome-thing
Open a Pull Request
Please read CONTRIBUTING.md for details on our code of conduct and pull-request process.


📬 Contact

Have questions or just want to say hi?
— Mustafa “Mushie” Siddiqui
— Email: mustafa.a.r.siddiqui@outlook.com

Happy interviewing! 🎉