import { Popconfirm, Tag } from "antd";
import { NumericFormat } from 'react-number-format';
import Button, { DangerButton } from '../../../components/general/button/Button.jsx'
import { Link } from "react-router-dom";

const column = (handleDelete) => {
  return ([
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
      let options = { day: 'numeric', month: 'long', year: 'numeric', hour: '2-digit', minute: '2-digit' };
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
    render: (_, record) => {
      return (
        <div className="d-flex justify-content-between">
          <Link to={`/transaction/editor/${record.uuid}`}>
            <Button text="Update" />
          </Link>
          <Popconfirm title="Sure to Delete?" onConfirm={() => handleDelete(record.uuid)}>
            <DangerButton text="Delete" />
          </Popconfirm>
        </div>
      );
    },
  },
])}

export default column;