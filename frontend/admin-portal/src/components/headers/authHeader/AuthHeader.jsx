import styles from "./AuthHeader.module.css";
function AuthHeader() {
  return (
    <header className={styles.authHeader}>
      <div className={styles.content}>
        <h2>Frubay Inc</h2>
        <div className={styles.message}>
          <p className={styles.messageContent}>
            “This library has saved me countless hours of work and helped me
            deliver stunning designs to my clients faster than ever before.”
          </p>
          <p className={styles.author}>Sofia Davis</p>
        </div>
      </div>
    </header>
  );
}

export default AuthHeader;
