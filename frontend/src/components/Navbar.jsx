import React from "react";

function Navbar({ user, logout }) {
  return (
    <div style={{ padding: 20, background: "#222", color: "white" }}>
      <h2>Minecraft Server</h2>

      {user ? (
        <>
          <span>Welcome {user.username}</span>
          <button onClick={logout} style={{ marginLeft: 10 }}>
            Logout
          </button>

          {user.is_admin && (
            <a href="/admin" style={{ marginLeft: 10 }}>
              Admin
            </a>
          )}
        </>
      ) : (
        <span>Please login</span>
      )}
    </div>
  );
}

export default Navbar;