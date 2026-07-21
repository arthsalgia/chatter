import React, { useState } from "react";
import { NavLink, useLocation } from "react-router-dom";

import "./Sidebar.css";

const PAGES = [
  { id: "dashboard", label: "Dashboard", path: "/dashboard" },
  { id: "validation", label: "Validation", path: "/validation" },
  { id: "watchlist", label: "Watchlist", path: "/watchlist" },
];

const SETTINGS_PAGE = {
  id: "settings",
  label: "Settings",
  path: "/settings",
};

export default function Sidebar() {
  const [collapsed, setCollapsed] = useState(false);
  const location = useLocation();
  const isActive = (path) => location.pathname === path;

  return (
    <div className={`sidebar${collapsed ? " collapsed" : ""}`}>
      <div className="sidebar-brand">
        <NavLink to="/" aria-label="Home">
          <img
            src={
              collapsed
                ? "src/assets/chatterIconCollapsed.svg"
                : "src/assets/chatterIcon.png"
            }
            alt="Home"
            className={collapsed ? "brand-icon-collapsed" : "brand-icon-full"}
          />
        </NavLink>
      </div>

      <nav className="sidebar-nav">
        {PAGES.map((page) => (
          <NavLink
            key={page.id}
            to={page.path}
            end={page.path === "/"}
            title={collapsed ? page.label : undefined}
            className={`sidebar-item${isActive(page.path) ? " active" : ""}`}
          >
            {collapsed ? page.label.charAt(0) : page.label}
          </NavLink>
        ))}
      </nav>

      <div className="sidebar-spacer" />

        <div className="sidebar-bottom">
          <NavLink
            to="/settings"
            title={collapsed ? "Settings" : undefined}
            className={`sidebar-item${isActive("/settings") ? " active" : ""}`}
          >
            {collapsed ? "S" : "Settings"}
          </NavLink>

          <button
            className="sidebar-toggle-btn"
            onClick={() => setCollapsed((c) => !c)}
            aria-label={collapsed ? "Expand sidebar" : "Collapse sidebar"}
          >
            {collapsed ? "›" : "‹"}
          </button>
        </div>
    </div>
  );
}