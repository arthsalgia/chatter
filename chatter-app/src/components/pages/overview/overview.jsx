import { useState } from "react";
import "./overview.css"

export default function Overview() {
  const [startDate, setStartDate] = useState("");

  return (
    <input
      type="date"
      className="date-input"
      value={startDate}
      onChange={(e) => setStartDate(e.target.value)}
    />
  );
}