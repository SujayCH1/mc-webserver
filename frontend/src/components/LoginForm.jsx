import React, { useState } from "react";
import API from "../api/api";

function LoginForm({ setUser }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const login = async () => {
    try {
      const res = await API.post("/login", {
        username,
        password,
      });

      localStorage.setItem("token", res.data.token);
      setUser(res.data);
    } catch (err) {
      console.error(err);
      alert("Login failed");
    }
  };

  return (
    <div>
      <h3>Login</h3>

      <input
        placeholder="Minecraft IGN"
        onChange={(e) => setUsername(e.target.value)}
      />

      <input
        type="password"
        placeholder="Password"
        onChange={(e) => setPassword(e.target.value)}
      />

      <button onClick={login}>Login</button>
    </div>
  );
}

export default LoginForm;