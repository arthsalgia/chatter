import { useState } from "react";
import "./overview.css"

import BestFriend from "../../bestFriend/bestFriend";
import BiggestFan from "../../biggestFan/biggestFan";
import Celebrity from "../../celebrity/celebrity";
import TopMessages from "../../topMessages/topMessages";
import MostCommonWord from "../../mostCommonWord/mostCommonWord";
import MostTextedDate from "../../mostTextedDate/mostTextedDate";
import Search from "../../search/search";
import MetaDataHeader from "../../metaData/metaData";
import SentimentAnalysis from "../../sentiment/sentimentAnalysis"

export default function Overview() {

  return (
    <div className="overview-grid">

        <MetaDataHeader />
        <BestFriend />
        <BiggestFan />
        <Celebrity />
        <TopMessages />
        <MostCommonWord />
        <MostTextedDate />
        <Search />
        <SentimentAnalysis />


    </div>
  );
}