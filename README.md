# Chatter

This is a full stack desktop application built by me (Arth salgia) with Go and React to analyze and give you stats on your iMessage history.

## Features

* **Smart Analytics:** Discover insights about your conversations - your best friend, biggest fan, most texted date, most common words, and more.
* **Lightning Fast Search:** Instantly query through your entire message history.
* **100% Private and Local:** Runs entirely on your machine using your local SQLite chat database. Nothing is ever sent anywhere.
* **Single Binary:** A completely self-contained Go server with the React frontend bundled inside — no dependencies to install.

---

## Installation and Usage

Install on macOS with the automated install script:

```bash
curl -sSL https://raw.githubusercontent.com/arthsalgia/chatter/main/install.sh | bash
```

This downloads the correct binary for your Mac (Intel or Apple Silicon), installs it to `/usr/local/bin`, and clears the macOS quarantine flag so it runs without extra prompts.

### Database Setup (IMPORTANT)

To run the analyzer, the application needs access to your iMessage database. On launch, it checks for one of two options:

1. **Local copy (recommended):** Place a copy of your `chat.db` file into the directory you're running the command from. You can find your system database at `~/Library/Messages/chat.db`.
2. **Full Disk Access:** Grant your Terminal application Full Disk Access via **System Settings → Privacy & Security → Full Disk Access** so it can read the system database directly.

### Running the Application

Once installed and your database is set up, just run:

```bash
chatter
```

Then open `http://localhost:8000` in your browser.

---

## Building from Source

If you'd rather build it yourself instead of using the install script (not sure why you would do this but if you insist on it):

```bash
git clone https://github.com/arthsalgia/chatter.git
cd chatter

# Build the frontend
cd chatter-app
npm install
npm run build
cd ..

# Build the Go binary (embeds the frontend build)
go build -o chatter .
```

Run it the same way as above: `./chatter`.

### Frontend development

For hot-reload during frontend development, run the Vite dev server alongside the Go API:

```bash
cd chatter-app
npm run dev
```

This serves the frontend on `http://localhost:5173`, which is already whitelisted in the backend's CORS config.

---

## API Reference

All endpoints are prefixed with `/api`.

| Method | Endpoint             | Description                               |
|--------|----------------------|-------------------------------------------|
| GET    | `/hello`             | Health check                              |
| GET    | `/get-all`           | Returns all raw messages                  |
| GET    | `/best-friend`       | Top message partner over a date range     |
| GET    | `/biggest-fan`       | Who's messaged you the most               |
| GET    | `/celebrity`         | Who you've messaged the most              |
| GET    | `/nth-common`        | Your N most frequently sent messages      |
| GET    | `/most-common-word`  | Most frequently used words                |
| GET    | `/most-texted-date`  | The day with the most message activity    |
| GET    | `/search`            | Search message history                    |
| GET    | `/meta-data`         | Overall stats summary                     |

---

## Privacy

Chatter reads your `chat.db` locally and serves everything from `localhost` - no message data is ever sent or stored by any external server or third party.

## Any questions?
Send an email to me on arthsalgia@gmail.com and I will do my best to answer any questions or concerns related to this project.