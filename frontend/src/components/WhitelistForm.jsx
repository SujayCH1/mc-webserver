import React from "react";
import API from "../api/api";

function WhitelistRequest({ user }) {

  const requestWhitelist = async () => {
    try {
      await API.post("/whitelist/request", {
        username: user.username,
        discord_id: user.discord_id,
        discord_username: user.discord_username,
      });

      alert("Whitelist request submitted!");
    } catch (err) {
      console.error(err);
      alert("Request failed");
    }
  };

  return (
    <div>
      <h3>Request Whitelist</h3>
      <button onClick={requestWhitelist}>
        Request Access
      </button>
    </div>
  );
}

export default WhitelistRequest;