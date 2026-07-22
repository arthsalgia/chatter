import { useState } from "react";
import "./overview.css"

import BestFriend from "../../bestFriend/bestFriend";
import BiggestFan from "../../biggestFan/biggestFan";
import Celebrity from "../../celebrity/celebrity";

export default function Overview() {

  return (
    <div className="overview-grid">
        
        <BestFriend />
        <BiggestFan />
        <Celebrity />

    </div>
  );
}