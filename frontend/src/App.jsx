import React, { useState } from "react";
import Home from "./pages/Home";
import Admin from "./pages/Admin";
import Navbar from "./components/Navbar";

function App() {
  const [user, setUser] = useState(null);

  const logout = () => {
    localStorage.removeItem("token");
    setUser(null);
  };

  const path = window.location.pathname;

  return (
    <>
      <Navbar user={user} logout={logout} />

      {path === "/admin" ? (
        <Admin />
      ) : (
        <Home user={user} setUser={setUser} />
      )}
    </>
  );
}

export default App;