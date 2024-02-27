import "react-datetime/css/react-datetime.css";
import Styles from "./GenderPicker.module.css";

function GenderPicker({ id, onChange, label }) {
  return (
    <div className="row">
      <label>Gender</label>
      <div className="row">
        <input type="radio" name="sex" value="f" data-icon="" />
        <input type="radio" name="sex" value="m" data-icon="" />
      </div>
    </div>
  );
}
//source : https://codepen.io/morten-olsen/pen/QbvBYy
export default GenderPicker;
