import Styles from "../Input.module.css";
import { NumericFormat } from 'react-number-format';

function NumericField({ id, onChange, label, value, disabled, className }) {

  const onValueChange = (values, sourceInfo) => {
    onChange && onChange({
      target: {
        id: id,
        value: parseInt(values.value || value)
      }
    });
  }

  return (
    <div className={Styles.formItem}>
      {label ? <label htmlFor={id}>{label}</label> : <></>}
      <div className={Styles.inputWrapper}>
        <NumericFormat
          id={id}
          disabled={disabled}
          // onChange={onChange}
          onValueChange={onValueChange}
          className={className}
          displayType="input"
          autoComplete="off"
          value={value}
          thousandSeparator={true}
        />
      </div>
    </div>
  );
}

export default NumericField;
