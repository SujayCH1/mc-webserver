import React, { useState } from "react";
import API from "../api/api";

function RegisterForm() {
  const [username, setUsername] = useState("");
  const [discordId, setDiscordId] = useState("");
  const [discordUsername, setDiscordUsername] = useState("");
  const [password, setPassword] = useState("");

  const register = async () => {
    try {
      await API.post("/register", {
        username,
        discord_id: discordId,
        discord_username: discordUsername,
        password,
      });

      alert("Registered successfully");
    } catch (err) {
      console.error(err);
      alert("Register failed");
    }
  };

  return (
    <div>
      <h3>Register</h3>

      <input
        placeholder="Minecraft IGN"
        onChange={(e) => setUsername(e.target.value)}
      />

      <input
        placeholder="Discord ID"
        onChange={(e) => setDiscordId(e.target.value)}
      />

      <input
        placeholder="Discord Username"
        onChange={(e) => setDiscordUsername(e.target.value)}
      />

      <input
        type="password"
        placeholder="Password"
        onChange={(e) => setPassword(e.target.value)}
      />

      <button onClick={register}>Register</button>
    </div>
  );
}

export default RegisterForm;