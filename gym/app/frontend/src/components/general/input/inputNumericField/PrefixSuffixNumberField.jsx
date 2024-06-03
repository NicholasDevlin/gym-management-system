import { useEffect, useState } from "react";
import Styles from "../Input.module.css";

function PrefixSuffixNumericField({ id, onChange, label, value, prefixOnClick, suffixOnClick }) {
  const [count, setCount] = useState(value || 0);

  const handlePrefix = () => {
    if (count > 0) {
      onChange(count - 1);
      setCount(count - 1);
      prefixOnClick && prefixOnClick();
    }
  }

  useEffect(() => {
    setCount(value);
  }, [value])

  const handleSuffix = () => {
    onChange(count + 1);
    setCount(count + 1);
    suffixOnClick && suffixOnClick();
  }

  return (
    <div className={Styles.formItem}>
      <label htmlFor={id}>{label}</label>
      <div className={Styles.inputWrapper}>
        <div className={`${Styles.prefix} w-30`} onClick={handlePrefix}>-</div>
        <input
          id={id}
          className="w-40 text-end"
          disabled
          type="number"
          autoComplete="off"
          autoCorrect="off"
          autoCapitalize="off"
          spellCheck="false"
          data-lpignore="true"
          value={count}
        />
        <div className={`${Styles.suffix} w-30`} onClick={handleSuffix}>+</div>
      </div>
    </div>
  );
}

export default PrefixSuffixNumericField;
