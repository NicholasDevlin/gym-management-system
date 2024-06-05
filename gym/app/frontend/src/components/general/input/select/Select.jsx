import SelectSearch from 'react-select-search';
import 'react-select-search/style.css';
import Styles from '../Input.module.css';
import { useRef, useState } from 'react';

function Select({ options, name, placeholder, label, value, getOptions }) {
  const [selected, setSelected] = useState(value);
  const selectRef = useRef(null);
  const [placement, setPlacement] = useState('on-bottom');

  const onChange = (e) => {
    setSelected(e);
  }

  const onFocus = () => {
    const { bottom } = selectRef.current.getBoundingClientRect();
    const viewportHeight = window.innerHeight;
    const middleScreen = viewportHeight / 2;
    if (bottom > middleScreen) {
      setPlacement('on-top');
    } else {
      setPlacement('on-bottom');
    }
  }

  return (
    <div className={Styles.formItem}>
      <label htmlFor={name}>{label}</label>
      <div className={Styles.inputWrapper}>
        <SelectSearch ref={selectRef} onFocus={onFocus} className={`${placement} select-search`} options={options} name={name} onChange={onChange} getOptions={getOptions} value={selected} search="true" placeholder={placeholder} />
      </div>
    </div>
  );
}
// source: https://www.npmjs.com/package/react-select-search
export default Select;
