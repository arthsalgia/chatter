import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Header from "./components/header/header";
import Overview from "./components/pages/overview/overview";
import "./App.css";

export default function App() {
  return (
    <BrowserRouter>
      <div className="app-layout">

        <div className="app-main">
          <Header />

          <main className="app-content">
            <Routes>
              <Route path="/dashboard" element={<Overview />} />
            </Routes>
          </main>
        </div>
      </div>
    </BrowserRouter>
  );
}