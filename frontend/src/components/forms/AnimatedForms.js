import React, { useState } from "react";
import { useSpring, animated } from "react-spring";
import SignUpForm from './signup/SignUpForm'
import LoginForm from './login/LogInForm'
import './AnimatedForms.css'
import FormSuccess from "./signup/FormSuccess";

export const AnimatedForms = ({signup,login}) => {

  const [signupFormStatus, setSignupFormStatus] = useState(signup);
  const [loginFormStatus, setLoginFormStatus] = useState(login);
  const [isSubmitted, setIsSubmitted] = useState(false);

  const loginProps = useSpring({
    left: signupFormStatus ? -500 : 0, // Login form sliding positions
  });
  const signupProps = useSpring({
    left: loginFormStatus ? 500 : 0, // Signup form sliding positions
  });

  const loginBtnProps = useSpring({
    borderBottom: signupFormStatus
      ? "solid 0px transparent"
      : "solid 2px #ff0000",  //Animate bottom border of login button
  });
  const signupBtnProps = useSpring({
    borderBottom: signupFormStatus
      ? "solid 2px #ff0000"
      : "solid 0px transparent", //Animate bottom border of signup button
  });

  function signupClicked() {
    setSignupFormStatus(true);
    setLoginFormStatus(false)
  }

  function loginClicked() {
    setSignupFormStatus(false);
    setLoginFormStatus(true)
  }

  function submitForm() {
    setIsSubmitted(true);
  }

  return (
    <div className="form-content">
      <div className="login-signup-wrapper">
        <div className="form-buttons">
          <animated.button
            onClick={loginClicked}
            id="loginBtn"
            style={loginBtnProps}
          >
            Log in
          </animated.button>
          <animated.button
            onClick={signupClicked}
            id="signupBtn"
            style={signupBtnProps}
          >
            Sign up
          </animated.button>
        </div>
        <div className="form-group">
          <animated.form id="loginform" style={loginProps}>
            <LoginForm />
          </animated.form>
          <animated.form id="signupform" style={signupProps}>
            {!isSubmitted ? (
              <SignUpForm submitForm={submitForm} />
            ) : (
              <FormSuccess />
            )}
          </animated.form>
        </div>
      </div>
    </div>
  );
}

export default AnimatedForms;