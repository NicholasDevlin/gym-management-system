import Styles from "../Input.module.css";

function NumericField({ id, onChange, label, value, disabled }) {
  return (
    <div className={Styles.formItem}>
      <label htmlFor={id}>{label}</label>
      <div className={Styles.inputWrapper}>
        <input
          id={id}
          className="text-end"
          disabled={disabled}
          onChange={onChange}
          type="number"
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

export default NumericField;
