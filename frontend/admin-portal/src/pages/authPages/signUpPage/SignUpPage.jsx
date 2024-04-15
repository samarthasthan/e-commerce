import React from "react";
import styles from "./SignUp.module.css";
import AuthHeader from "../../../components/headers/authHeader/AuthHeader";
import SimpleButton from "../../../components/buttons/SimpleButton";
import TextInput from "../../../components/inputs/textInput/TextInput";
import { Link } from "react-router-dom";

function SignUpPage() {
  return (
    <div className={styles.signUpPage}>
      <AuthHeader></AuthHeader>
      <div className={styles.signUpPageContent}>
        <div className={styles.loginButton}>
          <Link to={`/`}>
            <SimpleButton
              backgroundColor="var(--white-color)"
              color="var(--black-color)"
            >
              Login
            </SimpleButton>
          </Link>
        </div>
        <form action="" className={styles.signUpForm}>
          <h2>Create an account</h2>
          <p className="greyText">
            Enter your email below to create your account
          </p>
          <div className={`${styles.names}`}>
            <TextInput type="text">First Name</TextInput>
            <TextInput type="text">Last Name</TextInput>
          </div>
          <TextInput type="email">name@example.com</TextInput>
          <TextInput type="tel">+91 1234567890</TextInput>
          <TextInput type="password">password123</TextInput>
          <SimpleButton
            backgroundColor="var(--black-color)"
            color="var(--white-color)"
          >
            Sign up with Email
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

export default SignUpPage;
