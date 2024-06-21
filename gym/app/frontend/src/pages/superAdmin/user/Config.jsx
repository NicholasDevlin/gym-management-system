import { Popconfirm, Select, Tag } from "antd";
import { DangerButton } from '../../../components/general/button/Button.jsx'

const column = (handleDelete, roleOptions, updateUser) => {
  return ([
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      width: '20%'
    },
    {
      title: 'Phone Number',
      dataIndex: 'phoneNumber',
      key: 'phoneNumber',
      width: '20%',
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
      width: '20%',
    },
    {
      title: 'Gender',
      dataIndex: 'gender',
      key: 'gender',
      width: '10%',
      render: ((text) => {
        let color = 'blue';
        if (text === 'female') {
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
      title: 'Birtdate',
      dataIndex: 'birthDate',
      key: 'birthDate',
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
      title: 'Membership End date',
      dataIndex: 'subscriptionDueDate',
      key: 'subscriptionDueDate',
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
      title: 'Role',
      dataIndex: 'role',
      key: 'role',
      width: 150,
      render: ((_, value) => {
        const onSelect = (selected) => {
          value.role.role = roleOptions.find(x=> x.value === selected).label;
          updateUser(value);
        }
        return (
          <>
            <Select style={{ selectorBg: '#1f2124', color: "#000"}} onSelect={onSelect} className="w-100" defaultValue={value.role.id} options={roleOptions} />
          </>
        );
      })
    },
    {
      title: 'Action',
      dataIndex: 'action',
      key: 'action',
      width: '10%',
      render: (_, record) => {
        return (
          <div className="d-flex justify-content-between">
            <Popconfirm title="Sure to Delete?" onConfirm={() => handleDelete(record.uuid)}>
              <DangerButton text="Delete" />
            </Popconfirm>
          </div>
        );
      },
    },
  ])
}

export const activeOptions = [
  {
    name: "All",
    value: null
  },
  {
    name: "Active",
    value: "true"
  },
  {
    name: "Not Active",
    value: "false"
  }
];

export default column;