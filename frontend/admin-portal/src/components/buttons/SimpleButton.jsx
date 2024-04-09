import styles from "./SimpleButton.module.css";
function SimpleButton({
  children,
  backgroundColor = "var(--black-color)",
  color = "var(--white-color)",
}) {
  return (
    <div
      className={styles.simpleButton}
      style={{ backgroundColor: backgroundColor, color: color }}
    >
      {children}
    </div>
  );
}

export default SimpleButton;
