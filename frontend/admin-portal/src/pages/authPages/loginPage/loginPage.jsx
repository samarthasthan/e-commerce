import React from "react";
import styles from "./loginPage.module.css";
import AuthHeader from "../../../components/headers/authHeader/authHeader";
import SimpleButton from "../../../components/buttons/simpleButton";
import TextInput from "../../../components/inputs/textInput/textInput";

function LoginPage() {
  return (
    <div className={styles.loginPage}>
      <AuthHeader></AuthHeader>
      <div className={styles.loginPageContent}>
        <div className={styles.signUpButton}>
          <SimpleButton
            backgroundColor="var(--white-color)"
            color="var(--black-color)"
          >
            Sign Up
          </SimpleButton>
        </div>
        <form action="" className={styles.loginForm}>
          <h2>Login into your account</h2>
          <p className="greyText">
            Enter your Email and Password below to login into your account
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
            By clicking continue, you agree to our Terms of Service and Privacy
            Policy.
          </p>
        </form>
        <div className={styles.bottom}></div>
      </div>
    </div>
  );
}

export default LoginPage;
