import Styles from '../Input.module.css'
import 'react-phone-number-input/style.css'
import PhoneInput from 'react-phone-number-input'

function PhoneNumberInput({ id, label, onChange, value }) {
  const handlePhoneInputChange = (phoneNumber) => {
    onChange({ target: { id, value: phoneNumber } });
  };
  return (
    <div className={Styles.formItem}>
      <label htmlFor={id}>{label}</label>
      <div className={Styles.inputWrapper}>
        <PhoneInput
          international
          countryCallingCodeEditable={false}
          defaultCountry="ID"
          value={value}
          onChange={handlePhoneInputChange} />
      </div>
    </div>
  )
}
// Doc: https://www.npmjs.com/package/react-phone-number-input
export default PhoneNumberInput;