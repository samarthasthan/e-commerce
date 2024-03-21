import styles from "./textInput.module.css";
function TextInput({ children, type }) {
  return (
    <input
      type={type}
      name=""
      id=""
      className={styles.textInput}
      placeholder={children}
    />
  );
}

export default TextInput;
