import React from 'react';

function CurrencyFormat({ value }) {
  const formattedValue = new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
  }).format(value);

  return <span>{formattedValue}</span>;
}

export default CurrencyFormat;
