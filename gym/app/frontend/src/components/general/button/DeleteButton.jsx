import React, { useState } from 'react'
import Styles from './DeleteButton.module.css'

function DeleteButton() {
  const [dataDirection, setDataDirection] = useState('');

  const btnRef = React.useRef(null);
  let btn;
  const btnFrontClickHandler = (event) => {
    btn = document.getElementById('btn')
    btn.classList.toggle(Styles.btnIsOpen);

    const mx = event.clientX - btnRef.current.offsetLeft;
    const my = event.clientY - btnRef.current.offsetTop;

    const w = btnRef.current.offsetWidth;
    const h = btnRef.current.offsetHeight;

    var directions = [
      { id: 'top', x: w / 2, y: 0 },
      { id: 'right', x: w, y: h / 2 },
      { id: 'bottom', x: w / 2, y: h },
      { id: 'left', x: 0, y: h / 2 }
    ];

    directions.sort(function (a, b) {
      return distance(mx, my, a.x, a.y) - distance(mx, my, b.x, b.y);
    });

    setDataDirection(directions.shift().id);
  };

  const btnYesClickHandler = () => {
    btn = document.getElementById('btn')
    btn.classList.toggle(Styles.btnIsOpen);
  };

  const btnNoClickHandler = () => {
    btn = document.getElementById('btn')
    btn.classList.toggle(Styles.btnIsOpen);
  };

  const distance = (x1, y1, x2, y2) => {
    const dx = x1 - x2;
    const dy = y1 - y2;
    return Math.sqrt(dx * dx + dy * dy);
  };
  return (
    <div data-direction={dataDirection} className={Styles.btn} id="btn" ref={btnRef}>
      <div className={Styles.btnBack}>
        <p>Are you sure you want to do that?</p>
        <button className={Styles.yes} onClick={btnYesClickHandler}>Yes</button>
        <button className={Styles.no} onClick={btnNoClickHandler}>No</button>
      </div>
      <div className={Styles.btnFront} onClick={btnFrontClickHandler}>Delete</div>
    </div>
  );
}

// source https://codepen.io/hakimel/pen/ZYRgwB
export default DeleteButton;