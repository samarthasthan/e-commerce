import React from 'react'
import LoginHeader from '../../components/Headers/LoginHeader'
import styles from "./Login.module.css"
function Login() {
  return (
    <div className={styles.loginPage}>
      <LoginHeader></LoginHeader>
      <div className={styles.loginForm}>
         <p className={styles.signUpBtn}>Sign Up</p>
         <div className={styles.form}>
          <h2>Login into account</h2>
          <p className='greyText'>Enter your email and password below to login into your account</p>
          <input type="text" placeholder='Enter you email'/>
          <input type="text" placeholder='Enter you password' />
          <button>Login</button>
          <p className='greyText'>By clicking continue, you agree to our Terms of Service and Privacy Policy.</p>
         </div>
         <div></div>
      </div>
    </div>
  )
}

export default Login