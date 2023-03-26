import * as React from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import dayjs, {Dayjs} from 'dayjs';

export default function NewBugDialog() {
  const [open, setOpen] = React.useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };
  const handleSave = () => {
    setOpen(false);
  };
  return (
    <div style={{ display: 'inline' }}>
      <Button variant="outlined" onClick={handleClickOpen}>
        New Bug
      </Button>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Bug Creation</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Please input the new bug details.
          </DialogContentText>
          <TextField
            disabled
            margin="dense"
            id="date"
            label="date"
            type="surname"
            fullWidth
            defaultValue={dayjs(new Date()).toString()}
          />
          <TextField
            autoFocus
            margin="dense"
            id="name"
            label="Bug Title"
            type="bug_title"
            fullWidth
          /> 
          <TextField
            multiline
            margin="dense"
            rows={3}
            id="name"
            label="Bug Description"
            type="surname"
            fullWidth
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleSave}>Cancel</Button>
          <Button onClick={handleClose}>Submit</Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}
