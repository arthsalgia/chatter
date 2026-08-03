import { useState, useEffect } from "react";
import "./search.css";
import searchApi from "../../hooks/search";

export default function Search() {
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");
  const [word, setWord] = useState("");
  const [partial, setPartial] = useState(false);

  const [total, setTotal] = useState(0);
  const [fromMe, setFromMe] = useState(0);
  const [fromOthers, setFromOthers] = useState(0);

  const [searchLoading, setSearchLoading] = useState(false);

  useEffect(() => {
    async function getSearchResults() {
      try {
        setSearchLoading(true);

        const response = await searchApi(startDate, endDate, word, partial);

        setTotal(response.total ?? 0);
        setFromMe(response.from_me ?? 0);
        setFromOthers(response.from_others ?? 0);
      } catch (err) {
        console.log(err);
      } finally {
        setSearchLoading(false);
      }
    }

    getSearchResults();
  }, [startDate, endDate, word, partial]);

  return (
    <div className="search-page">
      <div className="search-card">
        <div className="search-header">
          <div className="header-text">
            <div className="header-eyebrow">Message Search</div>
            <div className="header-title">Search Messages</div>
          </div>

          <button
            type="button"
            className="info-button"
            aria-label="Search information"
          >
            <svg
              viewBox="0 0 20 20"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <circle
                cx="10"
                cy="10"
                r="8.5"
                stroke="currentColor"
                strokeWidth="1.5"
              />
              <path
                d="M10 9v5"
                stroke="currentColor"
                strokeWidth="1.5"
                strokeLinecap="round"
              />
              <circle cx="10" cy="6.3" r="1" fill="currentColor" />
            </svg>

            <span className="info-tooltip">
              Search for a word or phrase.
              <br />
              The results show how many times it appears in your messages.
              <br />
              Leave the dates empty to search your entire history.
            </span>
          </button>
        </div>

        <div className="date-section">
          <div className="date-group">
            <label className="date-text" htmlFor="search-start">
              Start Date
            </label>

            <input
              id="search-start"
              type="date"
              className="date-input"
              value={startDate}
              onChange={(e) => setStartDate(e.target.value)}
            />
          </div>

          <div className="date-group">
            <label className="date-text" htmlFor="search-end">
              End Date
            </label>

            <input
              id="search-end"
              type="date"
              className="date-input"
              value={endDate}
              onChange={(e) => setEndDate(e.target.value)}
            />
          </div>
        </div>

        <div className="search-input-section">
          <div className="date-group">
            <label className="date-text" htmlFor="search-word">
              Search Word
            </label>

            <div className="search-word-row">
              <input
                id="search-word"
                type="text"
                className="date-input search-word-input"
                placeholder=""
                value={word}
                onChange={(e) => setWord(e.target.value)}
              />

              <label className="partial-toggle-label" title="Enable partial matching">
                <input
                  type="checkbox"
                  className="partial-checkbox"
                  checked={partial}
                  onChange={(e) => setPartial(e.target.checked)}
                />
                <span className="partial-custom-tick">
                  <svg
                    viewBox="0 0 12 10"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      d="M1 5L4.5 8.5L11 1.5"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    />
                  </svg>
                </span>
                <span className="partial-text">Partial</span>
              </label>
            </div>
          </div>
        </div>

        <div className="search-result">
          {searchLoading ? (
            <div className="search-skeleton">
              <div className="skeleton-line skeleton-total" />
              <div className="skeleton-line skeleton-label" />

              <div className="search-breakdown">
                <div className="skeleton-box" />
                <div className="skeleton-box" />
              </div>
            </div>
          ) : (
            <>
              <div className="result-title">Occurrences</div>

              <div className="result-total">
                {total.toLocaleString()}
              </div>

              <div className="result-total-label">
                Total Matches
              </div>

              <div className="search-breakdown">
                <div className="breakdown-card">
                  <div className="breakdown-title">
                    From Me
                  </div>

                  <div className="breakdown-value">
                    {fromMe.toLocaleString()}
                  </div>
                </div>

                <div className="breakdown-card">
                  <div className="breakdown-title">
                    From Others
                  </div>

                  <div className="breakdown-value">
                    {fromOthers.toLocaleString()}
                  </div>
                </div>
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
}