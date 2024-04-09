import React from "react";
import styles from "./LoginPage.module.css";
import AuthHeader from "../../../components/headers/authHeader/AuthHeader";
import SimpleButton from "../../../components/buttons/SimpleButton";
import TextInput from "../../../components/inputs/textInput/TextInput";
import { Link } from "react-router-dom";

function LoginPage() {
  return (
    <div className={styles.loginPage}>
      <AuthHeader></AuthHeader>
      <div className={styles.loginPageContent}>
        <div className={styles.signUpButton}>
          <Link to={`/signup`}>
            <SimpleButton
              backgroundColor="var(--white-color)"
              color="var(--black-color)"
            >
              Sign Up
            </SimpleButton>
          </Link>
        </div>
        <form action="" className={styles.loginForm}>
          <h2>Login into your account</h2>
          <p className="greyText">
            Enter your email and password below to login into your account
          </p>
          <TextInput type="email">name@example.com</TextInput>
          <TextInput type="password">password123</TextInput>
          <SimpleButton
            backgroundColor="var(--black-color)"
            color="var(--white-color)"
          >
            Login
          </SimpleButton>
          <SimpleButton
            backgroundColor="var(--white-color)"
            color="var(--black-color)"
          >
            Forgot password
          </SimpleButton>
          <p className="greyText">
            By clicking continue, you agree to our{" "}
            <span style={{ textDecoration: "underline", cursor: "pointer" }}>
              Terms of Service
            </span>{" "}
            and{" "}
            <span style={{ textDecoration: "underline", cursor: "pointer" }}>
              Privacy Policy.
            </span>
          </p>
        </form>
        <div className={styles.bottom}></div>
      </div>
    </div>
  );
}

export default LoginPage;
