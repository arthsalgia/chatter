import { useState } from "react";
import "./overview.css"

import BestFriend from "../../bestFriend/bestFriend";
import BiggestFan from "../../biggestFan/biggestFan"

export default function Overview() {

  return (
    <div className="overview-grid">
        
        <BestFriend />
        <BiggestFan />

    </div>
  );
}