import Styles from "../Input.module.css";

function TextField({ id, onChange, label, value }) {
  return (
    <div className={Styles.formItem}>
      <label htmlFor={id}>{label}</label>
      <div className={Styles.inputWrapper}>
        <input
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

export default TextField;
