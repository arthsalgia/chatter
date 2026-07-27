import { useState } from "react";
import AboutModal from "./about/about";
import "./Header.css";

export default function Header() {
  const [isAboutOpen, setIsAboutOpen] = useState(false);

  return (
    <>
      <div className="header">
        <div className="brand">
          <img
            src={"src/assets/chatterIcon.png"}
            alt="Home"
            className={"brand-icon-full"}
          />
        </div>

        <div className="header-title">
          <h1>iMessage Analyzer</h1>
        </div>

        <div className="header-actions">
          <button 
            className="about-btn" 
            onClick={() => setIsAboutOpen(true)}
          >
            About
          </button>
        </div>
      </div>

      {isAboutOpen && <AboutModal onClose={() => setIsAboutOpen(false)} />}
    </>
  );
}