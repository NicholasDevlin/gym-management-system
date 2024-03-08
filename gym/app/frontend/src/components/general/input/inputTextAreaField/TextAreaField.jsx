import Styles from "../Input.module.css";

function TextAreaField({ id, onChange, label, value }) {
  return (
    <div className={Styles.formItem}>
      <label htmlFor={id}>{label}</label>
      <div className={Styles.inputWrapper}>
        <textarea
          rows="2"
          id={id}
          onChange={onChange}
          type="text"
          autoComplete="off"
          autoCorrect="off"
          autoCapitalize="off"
          spellCheck="false"
          data-lpignore="true"
          value={value}
        />
      </div>
    </div>
  );
}

export default TextAreaField;
