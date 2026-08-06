import { useState, useEffect } from "react";
import "./sentimentAnalysis.css"
import chatsApi from "../../hooks/chats";
import sentimentAnalysisApi from "../../hooks/sentimentAnalysis";

export default function SentimentAnalysis() {
  const [allChats, setAllChats] = useState([]);
  const [chatsLoading, setChatsLoading] = useState(false);

  const [query, setQuery] = useState("");
  const [selectedChatRaw, setSelectedChatRaw] = useState("");
  const [selectedChatLabel, setSelectedChatLabel] = useState("");
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  const [sentiment, setSentiment] = useState(null);
  const [sentimentLoading, setSentimentLoading] = useState(false);

  useEffect(() => {
    async function getChats() {
      try {
        setChatsLoading(true);
        const data = await chatsApi();

        const rawChats = data?.chats || [];
        const formattedChats = rawChats.map((chat) => ({
          raw: chat,
          label: chat,
        }));

        setAllChats(formattedChats);
      }
      catch (err) {
        console.log(err);
      }
      finally {
        setChatsLoading(false);
      }
    }

    getChats();
  }, []);

  useEffect(() => {
    if (!selectedChatRaw) {
      setSentiment(null);
      return;
    }

    async function getSentiment() {
      try {
        setSentimentLoading(true);
        const data = await sentimentAnalysisApi(selectedChatRaw);
        setSentiment(data);
      }
      catch (err) {
        console.log(err)
      }
      finally {
        setSentimentLoading(false);
      }
    }

    getSentiment()
  }, [selectedChatRaw]);

  const matchingChats = query.trim() === ""
    ? []
    : allChats.filter((chat) =>
        chat.label.toLowerCase().includes(query.trim().toLowerCase())
      ).slice(0, 8);

  function handleInputChange(e) {
    const value = e.target.value;
    setQuery(value);
    setIsDropdownOpen(true);

    if (value !== selectedChatLabel) {
      setSelectedChatRaw("");
      setSelectedChatLabel("");
    }
  }

  function handleSelectChat(chat) {
    setQuery(chat.label);
    setSelectedChatRaw(chat.raw);
    setSelectedChatLabel(chat.label);
    setIsDropdownOpen(false);
  }

  function handleBlur() {
    setTimeout(() => setIsDropdownOpen(false), 120);

    if (query !== selectedChatLabel) {
      setQuery(selectedChatLabel);
    }
  }

  const totalSentiment = sentiment?.totalSentiment ?? "0.0";
  const sentimentMe = sentiment?.sentimentMe ?? "0.0";
  const sentimentOther = sentiment?.sentimentOther ?? "0.0";
  const posMe = sentiment?.posMe ?? 0;
  const negMe = sentiment?.negMe ?? 0;
  const neuMe = sentiment?.neuMe ?? 0;
  const posOther = sentiment?.posOther ?? 0;
  const negOther = sentiment?.negOther ?? 0;
  const neuOther = sentiment?.neuOther ?? 0;

  return (
    <div className="sentiment-page">
      <div className="sentiment-card">
        <div className="sentiment-header">
          <div className="header-text">
            <div className="header-eyebrow">Message Stats</div>
            <div className="header-title">Sentiment Analysis</div>
          </div>

          <button
            type="button"
            className="info-button"
            aria-label="What is this?"
          >
            <svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle cx="10" cy="10" r="8.5" stroke="currentColor" strokeWidth="1.5" />
              <path d="M10 9v5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
              <circle cx="10" cy="6.3" r="1" fill="currentColor" />
            </svg>
            <span className="info-tooltip">
              Pick a chat to see the sentiment breakdown of your messages versus theirs.
            </span>
          </button>
        </div>

        <div className="chat-search-section">
          <div className="sentiment-field-group chat-search-wrapper">
            <label className="sentiment-field-label" htmlFor="chat-search-input">
              Chat
            </label>
            <input
              id="chat-search-input"
              type="text"
              className="sentiment-input"
              placeholder={chatsLoading ? "Loading chats..." : "Search for a chat..."}
              value={query}
              disabled={chatsLoading}
              onChange={handleInputChange}
              onFocus={() => setIsDropdownOpen(true)}
              onBlur={handleBlur}
              autoComplete="off"
            />

            {isDropdownOpen && query.trim() !== "" && (
              <div className="chat-dropdown">
                {matchingChats.length > 0 ? (
                  matchingChats.map((chat) => (
                    <button
                      key={chat.raw}
                      type="button"
                      className="chat-dropdown-item"
                      onMouseDown={(e) => e.preventDefault()}
                      onClick={() => handleSelectChat(chat)}
                    >
                      {chat.label}
                    </button>
                  ))
                ) : (
                  <div className="chat-dropdown-empty">No matching chats</div>
                )}
              </div>
            )}
          </div>

          <div className="sentiment-scale-hint">
            A score of -1 is very negative while 1 is very positive, a score of 0 is neutral
          </div>
        </div>

        <div className="sentiment-result">
          {sentimentLoading ? (
            <div className="sentiment-skeleton">
              <div className="skeleton-line skeleton-total" />
              <div className="skeleton-line skeleton-label" />
              <div className="sentiment-breakdown">
                <div className="skeleton-box" />
                <div className="skeleton-box" />
              </div>
            </div>
          ) : (
            <>
              <div className="sentiment-result-title">Overall Sentiment</div>
              <div className="sentiment-result-total">{totalSentiment}</div>
              <div className="sentiment-result-total-label">
                {selectedChatLabel ? `for ${selectedChatLabel}` : "select a chat to see results"}
              </div>

              <div className="sentiment-breakdown">
                <div className="sentiment-breakdown-card">
                  <div className="sentiment-breakdown-title">Your sentiment</div>
                  <div className="sentiment-breakdown-value">{sentimentMe}</div>
                </div>
                <div className="sentiment-breakdown-card">
                  <div className="sentiment-breakdown-title">Their sentiment</div>
                  <div className="sentiment-breakdown-value">{sentimentOther}</div>
                </div>
              </div>

              <div className="sentiment-tone-breakdown">
                <div className="tone-column">
                  <div className="sentiment-breakdown-title">You</div>
                  <div className="tone-row">
                    <span className="tone-label tone-label-positive">Positive</span>
                    <span className="tone-value tone-value-positive">{posMe}</span>
                  </div>
                  <div className="tone-row">
                    <span className="tone-label tone-label-negative">Negative</span>
                    <span className="tone-value tone-value-negative">{negMe}</span>
                  </div>
                  <div className="tone-row">
                    <span className="tone-label tone-label-neutral">Neutral</span>
                    <span className="tone-value tone-value-neutral">{neuMe}</span>
                  </div>
                </div>

                <div className="tone-divider" />

                <div className="tone-column">
                  <div className="sentiment-breakdown-title">Them</div>
                  <div className="tone-row">
                    <span className="tone-label tone-label-positive">Positive</span>
                    <span className="tone-value tone-value-positive">{posOther}</span>
                  </div>
                  <div className="tone-row">
                    <span className="tone-label tone-label-negative">Negative</span>
                    <span className="tone-value tone-value-negative">{negOther}</span>
                  </div>
                  <div className="tone-row">
                    <span className="tone-label tone-label-neutral">Neutral</span>
                    <span className="tone-value tone-value-neutral">{neuOther}</span>
                  </div>
                </div>
              </div>
              <div className="sentiment-tone-hint">
                Counts how many of your messages were classified as positive, negative, or neutral.
              </div>
              
            </>
          )}
        </div>
      </div>
    </div>
  );
}