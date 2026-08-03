import { useState, useEffect } from "react";
import "./metaData.css"
import metaDataApi from "../../hooks/metaData";

export default function MetaDataHeader() {
  const [metaData, setMetaData] = useState(null);
  const [metaDataLoading, setMetaDataLoading] = useState(false);

  const [showLongestMe, setShowLongestMe] = useState(false);
  const [showLongestOther, setShowLongestOther] = useState(false);

  useEffect(() => {
    async function getMetaData() {
      try {
        setMetaDataLoading(true);
        const data = await metaDataApi();
        setMetaData(data);
      }
      catch (err) {
        console.log(err)
      }
      finally {
        setMetaDataLoading(false);
      }
    }

    getMetaData()
  }, []);

  const totalMessages = metaData?.total_messages ?? 0;
  const messagesFromMe = metaData?.total_messages_me ?? 0;
  const messagesFromOthers = metaData?.total_messages_others ?? 0;

  const avgLengthMe = metaData?.avg_len_me ?? 0;
  const avgLengthOther = metaData?.avg_len_other ?? 0;
  const maxLengthMe = metaData?.max_len_me ?? 0;
  const maxLengthOther = metaData?.max_len_other ?? 0;

  const longestTextMe = metaData?.longestTextMe ?? "";
  const longestTextOther = metaData?.longestTextOther ?? "";

  // Who the longest message from you was sent to, and who sent you the longest message
  const sentTo = metaData?.sentTo ?? "";
  const sentBy = metaData?.sentBy ?? "";

  const myShare = totalMessages
    ? (messagesFromMe / totalMessages) * 100
    : 0;

  return (
    <div className="meta-page">
      <div className="meta-header">
        <div className="meta-title-row">
          <div className="header-text">
            <div className="header-eyebrow">Message Stats</div>
            <div className="header-title">Overview</div>
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
              A high-level summary of your entire message history.
            </span>
          </button>
        </div>

        {metaDataLoading ? (
          <div className="meta-skeleton">
            <div className="skeleton-line skeleton-totals" />
            <div className="skeleton-line skeleton-bar" />
            <div className="skeleton-line skeleton-comparison" />
          </div>
        ) : (
          <>
            <div className="meta-totals">
              <div className="total-stat total-stat-main">
                <div className="total-value">{totalMessages.toLocaleString()}</div>
                <div className="total-label">Total Messages</div>
              </div>
              <div className="total-stat">
                <div className="total-value">{messagesFromMe.toLocaleString()}</div>
                <div className="total-label">Messages Sent</div>
              </div>
              <div className="total-stat">
                <div className="total-value">{messagesFromOthers.toLocaleString()}</div>
                <div className="total-label">Messages Received</div>
              </div>
            </div>

            <div className="meta-share-bar">
              <div
                className="meta-share-fill"
                style={{ width: `${myShare}%` }}
              />
            </div>

            <div className="meta-comparison">
              <div className="comparison-column">
                <div className="comparison-column-title">You</div>
                <div className="comparison-row">
                  <span className="comparison-label">Avg Length (char)</span>
                  <span className="comparison-value">{avgLengthMe}</span>
                </div>

                <button
                  type="button"
                  className="comparison-row comparison-row-toggle"
                  onClick={() => setShowLongestMe((open) => !open)}
                  aria-expanded={showLongestMe}
                >
                  <span className="comparison-label">Longest Message (char)</span>
                  <span className="comparison-value-group">
                    <span className="comparison-value">{maxLengthMe.toLocaleString()}</span>
                    <svg
                      className={`toggle-chevron ${showLongestMe ? "toggle-chevron-open" : ""}`}
                      viewBox="0 0 20 20"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path d="M5 7.5L10 12.5L15 7.5" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" />
                    </svg>
                  </span>
                </button>

                {showLongestMe && (
                  <div className="longest-message-popup">
                    {sentTo && (
                      <div className="longest-message-meta">Sent to {sentTo}</div>
                    )}
                    {longestTextMe || "No message found."}
                  </div>
                )}
              </div>

              <div className="comparison-divider" />

              <div className="comparison-column">
                <div className="comparison-column-title">Sent to you</div>
                <div className="comparison-row">
                  <span className="comparison-label">Avg Length (char)</span>
                  <span className="comparison-value">{avgLengthOther}</span>
                </div>

                <button
                  type="button"
                  className="comparison-row comparison-row-toggle"
                  onClick={() => setShowLongestOther((open) => !open)}
                  aria-expanded={showLongestOther}
                >
                  <span className="comparison-label">Longest Message (char)</span>
                  <span className="comparison-value-group">
                    <span className="comparison-value">{maxLengthOther.toLocaleString()}</span>
                    <svg
                      className={`toggle-chevron ${showLongestOther ? "toggle-chevron-open" : ""}`}
                      viewBox="0 0 20 20"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path d="M5 7.5L10 12.5L15 7.5" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" />
                    </svg>
                  </span>
                </button>

                {showLongestOther && (
                  <div className="longest-message-popup">
                    {sentBy && (
                      <div className="longest-message-meta">From {sentBy}</div>
                    )}
                    {longestTextOther || "No message found."}
                  </div>
                )}
              </div>
            </div>
          </>
        )}
      </div>
    </div>
  );
}