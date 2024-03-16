import React from "react";
import styles from "./LoginHeader.module.css"

function LoginHeader(){
    return (
        <div className={styles.loginHeader}>
            <div className={styles.content}>
                <div className={styles.logo}>Amce Inc</div>
            <div>
                <p>Acme Inc
    “This library has saved me countless hours of work and helped me deliver stunning designs to my clients faster than ever before.”</p>
                <p className={styles.author}>Sofia Davis</p>
            </div>
            </div>
        </div>
    )
}

export default LoginHeader