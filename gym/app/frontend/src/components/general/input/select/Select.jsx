import SelectSearch from 'react-select-search';
import 'react-select-search/style.css';
import Styles from '../Input.module.css';
import { useState } from 'react';

function Select({ options, name, placeholder, label, value }) {
  const [selected, setSelected] = useState(value);
  const onChange = (e) => {
    setSelected(e);
  }

  return (
    <div className={Styles.formItem}>
      <label htmlFor={name}>{label}</label>
      <div className={Styles.inputWrapper}>
        <SelectSearch options={options} name={name} onChange={onChange} value={selected} search="true" placeholder={placeholder} />
      </div>
    </div>
  );
}
// source: https://www.npmjs.com/package/react-select-search
export default Select;
