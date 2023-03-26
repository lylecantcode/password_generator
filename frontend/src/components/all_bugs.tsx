import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import dayjs, {Dayjs} from 'dayjs';
import { DataGrid, GridColDef, GridValueGetterParams } from '@mui/x-data-grid';


const columns: GridColDef[] = [
  { field: 'id', headerName: 'ID', width: 70 },
  { field: 'bugTitle', headerName: 'Bug Title', width: 130 },
  { field: 'description', headerName: 'Description', width: 130 },
  {
    field: 'date',
    headerName: 'Date',
    // type: 'number',
    width: 90,
  },
  {
    field: 'assigned_to',
    headerName: 'Assigned To',
    description: 'This column has a value getter and is not sortable.',
    sortable: false,
    width: 160,
    valueGetter: (params: GridValueGetterParams) =>
      `${params.row.bugTitle || ''} ${params.row.description || ''}`,
  },
];

const rows = [
  { id: 1, description: 'Snow', bugTitle: 'Jon', date: 35 },
  { id: 2, description: 'Lannister', bugTitle: 'Cersei', date: 42 },
  { id: 3, description: 'Lannister', bugTitle: 'Jaime', date: 45 },
  { id: 4, description: 'Stark', bugTitle: 'Arya', date: 16 },
  { id: 5, description: 'Targaryen', bugTitle: 'Daenerys', date: null },
  { id: 6, description: 'Melisandre', bugTitle: null, date: 150 },
  { id: 7, description: 'Clifford', bugTitle: 'Ferrara', date: 44 },
  { id: 8, description: 'Frances', bugTitle: 'Rossini', date: 36 },
  { id: 9, description: 'Roxie', bugTitle: 'Harvey', date: 65 },
];


export default function DataTable() {
  const [tableRows, setTableRows] = React.useState(rows)

  React.useEffect(() => {
    const getAllOpenBugs = async() => {
      const data = await fetch('http://127.0.0.1:8000/bugs/all/')
      await data.json().then(res => console.log(res))
    }
    getAllOpenBugs()
  }, [tableRows])

  return (
    <div style={{ height: 400, width: '90%', margin:'5%' }}>
      <DataGrid
        rows={tableRows}
        columns={columns}
        pageSize={5}
        rowsPerPageOptions={[5]}
        checkboxSelection
      />
    </div>
  );
}
