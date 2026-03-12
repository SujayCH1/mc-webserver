import React from "react";

import WhitelistRequest from "../components/WhitelistForm";
import RegisterForm from "../components/RegisterForm";
import LoginForm from "../components/LoginForm";

function Home({ user, setUser }) {

  if (!user) {
    return (
      <>
        <LoginForm setUser={setUser} />
        <RegisterForm />
      </>
    );
  }

  if (!user.whitelisted) {
    return <WhitelistRequest user={user} />;
  }

  return <h2>You are whitelisted 🎉</h2>;
}

export default Home;