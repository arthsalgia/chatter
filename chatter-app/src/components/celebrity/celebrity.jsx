import { useState, useEffect, useRef } from "react";
import "./celebrity.css";
import celebrityApi from "../../hooks/celebrity";
import formatNumber from "../../services/formatNumber";

export default function Celebrity() {
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

  const [celebrity, setCelebrity] = useState("");
  const [noMessages, setNoMessages] = useState(0);
  const [displayCount, setDisplayCount] = useState(0);

  const [celebrityLoading, setCelebrityLoading] = useState(false);
  const rafRef = useRef(null);

  useEffect(() => {
    async function getCelebrity() {
      try {
        setCelebrityLoading(true);

        const data = await celebrityApi(startDate, endDate);

        setCelebrity(formatNumber(data.celebrity));
        setNoMessages(data.number_of_messages);
      } catch (err) {
        console.log(err);
      } finally {
        setCelebrityLoading(false);
      }
    }

    getCelebrity();
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
    <div className="celebrity-page">
      <div className="celebrity-card">
        <div className="celebrity-header">
          <div className="header-text">
            <div className="header-eyebrow">Message Stats</div>
            <div className="header-title">Celebrity</div>
          </div>

          <button
            type="button"
            className="info-button"
            aria-label="What is this?"
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
              The celebrity you messaged the most between the selected dates.
              <br />
              If the dates are empty, all messages are included.
            </span>
          </button>
        </div>

        <div className="date-section">
          <div className="date-group">
            <label className="date-text" htmlFor="celebrity-start">
              Start Date
            </label>

            <input
              id="celebrity-start"
              type="date"
              className="date-input"
              value={startDate}
              onChange={(e) => setStartDate(e.target.value)}
            />
          </div>

          <div className="date-group">
            <label className="date-text" htmlFor="celebrity-end">
              End Date
            </label>

            <input
              id="celebrity-end"
              type="date"
              className="date-input"
              value={endDate}
              onChange={(e) => setEndDate(e.target.value)}
            />
          </div>
        </div>

        <div className="celebrity-result">
          {celebrityLoading ? (
            <div className="celebrity-skeleton">
              <div className="skeleton-line skeleton-label" />
              <div className="skeleton-line skeleton-name" />
              <div className="skeleton-line skeleton-count" />
            </div>
          ) : (
            <>
              <div className="result-title">Your Celebrity</div>

              <div className="result-name">
                {celebrity || "—"}
                <span
                  className="result-underline"
                  aria-hidden="true"
                />
              </div>

              <div className="message-count">
                <span className="message-count-number">
                  {displayCount.toLocaleString()}
                </span>{" "}
                messages
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
