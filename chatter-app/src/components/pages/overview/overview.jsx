import { useState } from "react";
import "./overview.css"

import BestFriend from "../../bestFriend/bestFriend";

export default function Overview() {
  const [startDate, setStartDate] = useState("");

  return (
    <div>
        
    <BestFriend />
    </div>
  );
}