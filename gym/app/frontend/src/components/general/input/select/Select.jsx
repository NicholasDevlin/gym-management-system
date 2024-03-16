import SelectSearch from 'react-select-search';
import 'react-select-search/style.css';
import Styles from '../Input.module.css';

function Select({ options, name, placeholder, label }) {
  return (
    <div className={Styles.formItem}>
      <label htmlFor={name}>{label}</label>
      <div className={Styles.inputWrapper}>
        <SelectSearch options={options} name={name} search="true" placeholder={placeholder} />
      </div>
    </div>
  );
}

export default Select;
