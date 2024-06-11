import { Tag } from "antd";
import { NumericFormat } from 'react-number-format';

export const column = [
  {
    title: 'Transaction No',
    dataIndex: 'transactionNo',
    key: 'transactionNo',
    width: '20%'
  },
  {
    title: 'Transaction Date',
    dataIndex: 'transactionDate',
    key: 'transactionDate',
    width: '20%',
    render: ((text) => {
      let options = { day: 'numeric', month: 'long', year: 'numeric' };
      let date = new Date(text);
      return (
        <>
          {date.toLocaleDateString("id-ID", options)}
        </>
      )
    })
  },
  {
    title: 'Status',
    dataIndex: 'status',
    key: 'status',
    width: '20%',
    render: ((text) => {
      let color = 'green';
      if (text === 'Waiting for payment') {
        color = 'volcano';
      }
      return (
        <Tag color={color} key={text}>
          {text.toUpperCase()}
        </Tag>
      );
    })
  },
  {
    title: 'Total',
    dataIndex: 'total',
    key: 'total',
    width: '20%',
    render: ((text) => {
      return (
        <NumericFormat value={text} displayType='text' thousandSeparator prefix="Rp. " />
      )
    })
  },
  {
    title: 'Action',
    dataIndex: 'action',
    key: 'action',
    width: '20%',
    render: (_, record) => (
      <div className="d-flex justify-content-between">
        <button>Update</button>
        <button>Delete</button>
      </div>
    ),
  },
]