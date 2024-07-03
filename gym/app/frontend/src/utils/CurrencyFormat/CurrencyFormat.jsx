import React from 'react';

function CurrencyFormat({ value }) {
  const formattedValue = new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
  }).format(value);

  return <span>{formattedValue}</span>;
}

export function formatNumberWithCommas(number) {
  const formatter = new Intl.NumberFormat('en-US', {
    style: 'decimal',
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  });

  return formatter.format(number);
}

export function formatDateTime(input) {
  const date = new Date(input);
  const day = String(date.getDate()).padStart(2, '0'); // Ensure two-digit day
  const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-indexed
  const year = date.getFullYear();
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');

  return `${day}/${month}/${year} ${hours}:${minutes}`;
}

export default CurrencyFormat;
