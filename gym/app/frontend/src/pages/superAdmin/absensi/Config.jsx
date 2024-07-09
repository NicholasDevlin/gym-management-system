import { Popconfirm } from "antd";
import Button, { DangerButton } from '../../../components/general/button/Button.jsx'

const column = (handleDelete, setSelected, checkOutAbsensi) => {
  return ([
    {
      title: 'Member',
      dataIndex: 'name',
      key: 'name',
      width: '20%',
      render: ((_, record) => {
        return (
          <>
            {record.user.name} - ({record.user.phoneNumber})
          </>
        )
      })
    },
    {
      title: 'Check In',
      dataIndex: 'date',
      key: 'date',
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
      title: 'Check Out',
      dataIndex: 'checkOut',
      key: 'checkOut',
      width: '20%',
      render: ((text) => {
        let options = { day: 'numeric', month: 'long', year: 'numeric', hour: '2-digit', minute: '2-digit' };
        let date = new Date(text);
        return (
          <>
            {text ? date.toLocaleDateString("id-ID", options) : "-"}
          </>
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
          <div className="d-flex justify-content-end">
            <Button text="Check Out" className="me-3" onClick={() => checkOutAbsensi(record)} />
            <Button text="Update" className="me-3" onClick={() => setSelected(record)} />
            <Popconfirm title="Sure to Delete?" onConfirm={() => handleDelete(record.uuid)}>
              <DangerButton text="Delete" />
            </Popconfirm>
          </div>
        );
      },
    },
  ])
}

export default column;