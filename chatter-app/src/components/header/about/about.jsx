import "./about.css";

export default function AboutModal({ onClose }) {
  return (
    <div className="modal-backdrop">
      <div className="modal-content">
        <button className="modal-close-btn" onClick={onClose}>
          &times;
        </button>

        <div className="modal-body">
          <h2>Hello there!</h2>

          <p className="intro-text">I'm Arth Salgia.</p>

          <div className="modal-footer"></div>

          <p>
            Thank you for checking out my project! 
            This application is a full stack analytics tool designed to process and parse local iMessage SQLite database
            to give you some interesting details about how you communicate with other people.
          </p>


          <div className="modal-footer">
            <span>Find me on:</span>
            <a href="https://github.com/arthsalgia" target="_blank" rel="noreferrer">
              Github
            </a>
            <a href="https://linkedin.com/in/arthsalgia" target="_blank" rel="noreferrer">
              LinkedIn
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}