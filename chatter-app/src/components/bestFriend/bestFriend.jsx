import { useState, useEffect, useRef } from "react";
import "./bestFriend.css"
import bestFriendApi from "../../hooks/bestFriend";
import formatNumber from "../../services/formatNumber";

export default function BestFriend() {
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

  const [bf, setBf] = useState("");
  const [noMessages, setNoMessages] = useState(0);
  const [displayCount, setDisplayCount] = useState(0);

  const [bfLoading, setBfLoading] = useState(false);
  const rafRef = useRef(null);

  useEffect(() => {
    async function getBestFriend() {
      try {
        setBfLoading(true);
        const data = await bestFriendApi(startDate, endDate);
        setBf(formatNumber(data.best_friend));
        setNoMessages(data.number_of_messages);
      }
      catch (err) {
        console.log(err)
      }
      finally {
        setBfLoading(false);
      }
    }

    getBestFriend()
  }, [startDate, endDate]);

  useEffect(() => {
    if (rafRef.current) cancelAnimationFrame(rafRef.current);

    const target = Number(noMessages) || 0;
    const duration = 600;
    const start = performance.now();
    const from = displayCount;

    function tick(now) {
      const progress = Math.min((now - start) / duration, 1);
      const eased = 1 - Math.pow(1 - progress, 3); 
      setDisplayCount(Math.round(from + (target - from) * eased));
      if (progress < 1) {
        rafRef.current = requestAnimationFrame(tick);
      }
    }

    rafRef.current = requestAnimationFrame(tick);
    return () => cancelAnimationFrame(rafRef.current);
  }, [noMessages]);

  return (
    <div className="bf-page">
      <div className="bf-card">
        <div className="bf-header">
          <div className="header-text">
            <div className="header-eyebrow">Message Stats</div>
            <div className="header-title">Best Friend</div>
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
              Your most chatted with person between the given date fields.<br/>
              If dates are empty, checks for all dates.
            </span>
          </button>
        </div>

        <div className="date-section">
          <div className="date-group">
            <label className="date-text" htmlFor="bf-start">Start Date</label>
            <input
              id="bf-start"
              type="date"
              className="date-input"
              value={startDate}
              onChange={(e) => setStartDate(e.target.value)}
            />
          </div>

          <div className="date-group">
            <label className="date-text" htmlFor="bf-end">End Date</label>
            <input
              id="bf-end"
              type="date"
              className="date-input"
              value={endDate}
              onChange={(e) => setEndDate(e.target.value)}
            />
          </div>
        </div>

        <div className="bf-result">
          {bfLoading ? (
            <div className="bf-skeleton">
              <div className="skeleton-line skeleton-label" />
              <div className="skeleton-line skeleton-name" />
              <div className="skeleton-line skeleton-count" />
            </div>
          ) : (
            <>
              <div className="result-title">Your Best Friend</div>
              <div className="result-name">
                {bf || "—"}
                <span className="result-underline" aria-hidden="true" />
              </div>
              <div className="message-count">
                <span className="message-count-number">{displayCount.toLocaleString()}</span> messages
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
}