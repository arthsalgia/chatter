import { useState, useEffect } from "react";
import "./MostTextedDate.css"
import mostTextedDateApi from "../../hooks/mostTextedDate";

export default function MostTextedDate() {
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");
  const [recordLimit, setRecordLimit] = useState(5);

  const [topMessages, setTopMessages] = useState([]);
  const [topMessagesLoading, setTopMessagesLoading] = useState(false);

  useEffect(() => {
    async function getTopMessages() {
      try {
        setTopMessagesLoading(true);
        const response = await mostTextedDateApi(startDate, endDate, recordLimit);
        const messageRankings = response?.[0] ?? [];
        setTopMessages(messageRankings);
      }
      catch (err) {
        console.log(err)
      }
      finally {
        setTopMessagesLoading(false);
      }
    }

    getTopMessages()
  }, [startDate, endDate, recordLimit]);

  const highestMessageCount = topMessages.reduce(
    (max, item) => Math.max(max, item.number_of_messages),
    0
  );

  return (
    <div className="bf-page">
      <div className="bf-card">
        <div className="bf-header">
          <div className="header-text">
            <div className="header-eyebrow">Message Stats</div>
            <div className="header-title">Most Texted Date</div>
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
              Your most texted on date between the given date fields. Format <br/>YYYY-MM-DD<br/>
              If dates are empty, checks for all dates.
            </span>
          </button>
        </div>

        <div className="date-section">
          <div className="date-group">
            <label className="date-text" htmlFor="tm-start">Start Date</label>
            <input
              id="tm-start"
              type="date"
              className="date-input"
              value={startDate}
              onChange={(e) => setStartDate(e.target.value)}
            />
          </div>

          <div className="date-group">
            <label className="date-text" htmlFor="tm-end">End Date</label>
            <input
              id="tm-end"
              type="date"
              className="date-input"
              value={endDate}
              onChange={(e) => setEndDate(e.target.value)}
            />
          </div>
        </div>

        <div className="record-limit-section">
          <div className="date-group">
            <label className="date-text" htmlFor="tm-record-limit">Number of Records</label>
            <input
              id="tm-record-limit"
              type="number"
              min="1"
              placeholder="5"
              className="date-input record-limit-input"
              value={recordLimit}
              onChange={(e) => setRecordLimit(Number(e.target.value) || 1)}
            />
          </div>
        </div>

        <div className="tm-result">
          <div className="result-title">Your Most Texted Dates</div>

          {topMessagesLoading ? (
            <div className="tm-skeleton">
              {[0, 1, 2, 3].map((i) => (
                <div key={i} className="skeleton-line skeleton-row" />
              ))}
            </div>
          ) : topMessages.length === 0 ? (
            <div className="tm-empty">No messages found for this range.</div>
          ) : (
            <ul className="tm-list">
              {topMessages.map((item, index) => {
                const messageText = item.data;
                const messageCount = item.number_of_messages;
                const barWidth = highestMessageCount
                  ? (messageCount / highestMessageCount) * 100
                  : 0;

                return (
                  <li key={`${messageText}-${index}`} className="tm-row">
                    <div className="tm-rank">{index + 1}</div>
                    <div className="tm-row-main">
                      <div className="tm-row-top">
                        <span className="tm-message-text">{messageText}</span>
                        <span className="tm-message-count">{messageCount.toLocaleString()}</span>
                      </div>
                      <div className="tm-bar-track">
                        <div
                          className="tm-bar-fill"
                          style={{ width: `${barWidth}%` }}
                        />
                      </div>
                    </div>
                  </li>
                );
              })}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}